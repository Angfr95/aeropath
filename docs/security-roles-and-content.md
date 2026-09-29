# Security: roles and lesson content

## Database migration

Apply `scripts/migrations/005_student_roles.sql` before starting the updated API.
Docker only runs initialization scripts automatically for a new database volume.
Existing local database:

```powershell
docker exec aeropath-local-postgres-1 psql -U aeropath -d aeropath -v ON_ERROR_STOP=1 -f /docker-entrypoint-initdb.d/005_student_roles.sql
```

All existing and newly registered accounts default to `student`. No account is
promoted automatically. An operator can grant or revoke `students.role` directly
in the database for a verified user ID. Registration and profile APIs cannot
change this field. Administrator permissions are checked in the database on each
restricted request; revocation does not require waiting for JWT expiry.

Students can read and delete their own profile and update their own language.
Reading/deleting another profile, listing/counting users, writing lessons or
questions, and reading global admin statistics require the admin role.

## Lesson rendering

Raw HTML in lesson content is escaped before Markdown formatting. HTML is shown
as text; headings, emphasis and lists remain supported. Existing stored content
is protected when rendered with the updated frontend.

Rebuild and serve the updated frontend as well as the API. In production the
nginx compose service mounts `web/`, so its static files must be updated from
`web/frontend-react/dist/`; rebuilding only the API image does not update that
nginx mount. No production deployment is performed by this local correction.

## Regression checks

```powershell
go test ./internal/transport/middleware ./internal/auth ./internal/learning
cd web/frontend-react
node --test tests/renderMarkdown.test.mjs
npm.cmd run build
```
