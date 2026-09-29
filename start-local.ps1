$ErrorActionPreference = 'Stop'
Set-Location $PSScriptRoot
docker compose -p aeropath-local -f compose.local.yml up -d --wait
if ($LASTEXITCODE -ne 0) { throw 'PostgreSQL local indisponible. Vérifiez Docker Desktop.' }
New-Item -ItemType Directory -Force .local | Out-Null
if (-not (Get-NetTCPConnection -State Listen -LocalPort 8080 -ErrorAction SilentlyContinue)) {
    go build -o .local/api-gateway.exe ./cmd/api-gateway
    if ($LASTEXITCODE -ne 0) { throw 'Compilation Go échouée.' }
    Start-Process -FilePath "$PSScriptRoot/.local/api-gateway.exe" -WorkingDirectory $PSScriptRoot -WindowStyle Hidden -RedirectStandardOutput "$PSScriptRoot/.local/api.log" -RedirectStandardError "$PSScriptRoot/.local/api-error.log"
}
if (-not (Get-NetTCPConnection -State Listen -LocalPort 3000 -ErrorAction SilentlyContinue)) {
    Start-Process -FilePath (Get-Command node.exe).Source -ArgumentList 'node_modules/vite/bin/vite.js --host 127.0.0.1 --strictPort' -WorkingDirectory "$PSScriptRoot/web/frontend-react" -WindowStyle Hidden -RedirectStandardOutput "$PSScriptRoot/.local/frontend.log" -RedirectStandardError "$PSScriptRoot/.local/frontend-error.log"
}
Write-Host 'Aeropath : http://localhost:3000'
Write-Host 'Journaux : .local/'
