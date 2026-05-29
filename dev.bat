@echo off
title AeroPath Dev

echo ========================================
echo    🛩️  AeroPath - Mode Développement
echo ========================================
echo.

:: Lancer le backend API (sert aussi les fichiers statiques)
echo Demarrage du serveur sur http://localhost:8080 ...
start "AeroPath" cmd /c "cd /d %~dp0 && go run ./cmd/api-gateway"

echo.
echo ========================================
echo    ✅ Serveur : http://localhost:8080
echo ========================================
echo.
echo Appuie sur Ctrl+C dans la fenetre pour arreter.
pause
