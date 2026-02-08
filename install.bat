@echo off
TITLE Switchblade Installer

:: ========================================================
:: SWITCHBLADE v1.4 - AUTO INSTALLER
:: ========================================================


net session >nul 2>&1
if %errorLevel% == 0 (
    goto :RunInstall
) else (
    echo [INFO] Requesting Administrator rights...
    powershell -Command "Start-Process '%0' -Verb RunAs"
    exit
)

:RunInstall
cls
echo ========================================================
echo      INSTALLING SWITCHBLADE v1.4
echo ========================================================
echo.


if not exist "C:\Program Files\Switchblade" (
    mkdir "C:\Program Files\Switchblade"
    echo [OK] Created folder: C:\Program Files\Switchblade
)


copy "%~dp0switchblade.exe" "C:\Program Files\Switchblade\switchblade.exe" /Y >nul

if %errorlevel% neq 0 (
    echo.
    echo [ERROR] Could not find switchblade.exe! 
    echo Please make sure the .exe and this .bat are in the same folder.
    echo.
    pause
    exit
)
echo [OK] Copied switchblade.exe to Program Files


echo [INFO] Updating System PATH...
setx PATH "%PATH%;C:\Program Files\Switchblade" /M >nul
echo [OK] System PATH updated successfully.

echo.
echo ========================================================
echo      INSTALLATION COMPLETE! 🚀
echo ========================================================
echo.
echo You can now open a NEW terminal and type: switchblade help
echo.
pause