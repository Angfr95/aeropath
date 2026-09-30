#!/usr/bin/env bash
# Run as root on the existing AeroPath VM, with a clean release directory and immutable image.
set -Eeuo pipefail
umask 077
release=$(realpath "${1:?release directory required}")
image=${2:?immutable API image required}
case "$release" in /opt/aeropath/releases/*) ;; *) echo 'Unexpected release directory' >&2; exit 1;; esac
test -f "$release/nginx.conf"
test -f "$release/scripts/seed/ppl_content.sql"
exec 9>/run/lock/aeropath-deploy.lock
flock -n 9 || { echo 'Another deployment is running'; exit 1; }

cd /opt/aeropath
backup=/var/backups/aeropath-release-$(date -u +%Y%m%dT%H%M%SZ)
install -d -m 700 "$backup"
docker exec aeropath-postgres pg_dump -U aeropath -d aeropath -Fc > "$backup/database.dump"
test -s "$backup/database.dump"
docker exec -i aeropath-postgres pg_restore --list < "$backup/database.dump" > "$backup/database-contents.txt"
cp -p .env nginx.conf docker-compose.yml "$backup/"
old_image=$(docker inspect --format '{{.Image}}' aeropath-api)
if [ -f compose.release.yml ]; then
    cp -p compose.release.yml "$backup/compose.release.yml"
else
    printf 'services:\n  api-gateway:\n    image: %s\n' "$old_image" > "$backup/compose.release.yml"
fi
compose=(docker compose -f docker-compose.yml -f compose.release.yml)
armed=0
rollback() {
    result=$?
    trap - EXIT
    docker rm -f aeropath-api-candidate >/dev/null 2>&1 || true
    docker rm aeropath-release-export >/dev/null 2>&1 || true
    if [ "$result" -ne 0 ] && [ "$armed" = 1 ]; then
        cp "$backup/compose.release.yml" compose.release.yml
        "${compose[@]}" up -d --no-deps api-gateway nginx > "$backup/rollback.log" 2>&1 || true
        echo "Deployment failed; previous application configuration restored. Private diagnostics: $backup" >&2
    fi
    exit "$result"
}
trap rollback EXIT

# Export only compiled public files, never the repository or environment files.
docker create --name aeropath-release-export "$image" >/dev/null
install -d -m 755 "$release/public"
docker cp aeropath-release-export:/app/web/. "$release/public/"
docker rm aeropath-release-export >/dev/null
chmod -R a+rX "$release/public"
python3 - "$release/public" <<'PY'
import json, re, subprocess, sys
from pathlib import Path
web = Path(sys.argv[1])
env = json.loads(subprocess.check_output(['docker','inspect','--format','{{json .Config.Env}}','aeropath-api']))
secrets = [item.split('=',1)[1].encode() for item in env
           if re.search(r'PASSWORD|SECRET|TOKEN|DATABASE_URL', item.split('=',1)[0])
           and len(item.split('=',1)[1]) >= 8]
allowed = {'.html','.js','.css','.svg','.json','.ico','.png','.jpg','.jpeg','.webp','.woff','.woff2','.webmanifest','.txt'}
for path in web.rglob('*'):
    assert not path.is_symlink(), 'Public symlink rejected'
    if not path.is_file(): continue
    assert path.suffix in allowed and not any(p.startswith('.') for p in path.relative_to(web).parts), 'Unexpected public file'
    data = path.read_bytes()
    assert not any(secret in data for secret in secrets), 'Runtime secret found in public files'
    assert not re.search(rb'-----BEGIN (?:RSA |EC |OPENSSH )?PRIVATE KEY-----|\bAKIA[A-Z0-9]{16}\b|\bghp_[A-Za-z0-9]{30,}', data), 'Secret signature found'
print('Compiled public files checked; no runtime secrets found.')
PY

python3 "$release/deployments/aws/apply_content.py" "$release" "$backup/content-import-error.log"

cat > "$backup/new-compose.yml" <<EOF
services:
  api-gateway:
    image: $image
  nginx:
    volumes:
      - $release/nginx.conf:/etc/nginx/nginx.conf:ro
      - $release/public:/usr/share/nginx/html:ro
EOF
armed=1
cp "$backup/new-compose.yml" compose.release.yml
"${compose[@]}" run -d --no-deps --name aeropath-api-candidate \
    --publish 127.0.0.1:18080:8080 api-gateway > "$backup/candidate-container.txt"
ready=0
for attempt in $(seq 1 30); do
    if curl -fsS --max-time 2 http://127.0.0.1:18080/health >/dev/null; then ready=1; break; fi
    sleep 1
done
test "$ready" = 1
python3 "$release/deployments/aws/verify_release.py" "$release" http://127.0.0.1:18080 > "$backup/candidate-verification.json"
cat "$backup/candidate-verification.json"
docker rm -f aeropath-api-candidate >/dev/null

# Only the API and the web server are replaced. Stateful/analytics services are untouched.
"${compose[@]}" up -d --no-deps api-gateway nginx > "$backup/switch.log" 2>&1
ready=0
for attempt in $(seq 1 30); do
    if curl -fsS --max-time 3 https://aeropath.ultizion.com/health >/dev/null; then ready=1; break; fi
    sleep 1
done
test "$ready" = 1
docker exec aeropath-nginx nginx -t
python3 "$release/deployments/aws/verify_release.py" "$release" https://aeropath.ultizion.com --public > "$backup/public-verification.json"
cat "$backup/public-verification.json"
printf '%s\n' "$release" > /opt/aeropath/active-release
armed=0
echo "Deployment verified. Private backup: $backup"
