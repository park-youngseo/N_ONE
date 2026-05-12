@echo off
setlocal enabledelayedexpansion
:: ==========================================
:: NONE-OS Premiere Pro 2025 Uninstaller / Nuke
:: ==========================================
color 0C
echo Requesting Administrative Privileges...
net session >nul 2>&1
if %errorLevel% NEQ 0 (
    echo [ERROR] Administrator rights are required.
    echo Please right-click this file and select "Run as administrator".
    pause
    exit
)

echo.
echo ===================================================
echo Initiating TOTAL ERADICATION of Premiere Pro 2025
echo ===================================================
echo.

:: 1. Terminate any running Premiere processes
echo [PHASE 1] Terminating Premiere Pro processes...
taskkill /F /IM "Adobe Premiere Pro.exe" >nul 2>&1
taskkill /F /IM "LogTransport2.exe" >nul 2>&1
taskkill /F /IM "CEPHtmlEngine.exe" >nul 2>&1
taskkill /F /IM "dynamiclinkmanager.exe" >nul 2>&1

:: 2. Delete Main Program Files
echo [PHASE 2] Destroying Main Application Files...
set "PROG_PATH=C:\Program Files\Adobe\Adobe Premiere Pro 2025"
if exist "%PROG_PATH%" (
    rmdir /S /Q "%PROG_PATH%"
    echo [DELETED] %PROG_PATH%
) else (
    echo [SKIPPED] Main program folder not found.
)

:: 3. Delete AppData (Roaming)
echo [PHASE 3] Destroying AppData (Roaming) Settings...
set "ROAMING_PATH=%APPDATA%\Adobe\Premiere Pro\25.0"
if exist "%ROAMING_PATH%" (
    rmdir /S /Q "%ROAMING_PATH%"
    echo [DELETED] %ROAMING_PATH%
)

:: 4. Delete AppData (Local)
echo [PHASE 4] Destroying AppData (Local) Cache...
set "LOCAL_PATH=%LOCALAPPDATA%\Adobe\Premiere Pro\25.0"
if exist "%LOCAL_PATH%" (
    rmdir /S /Q "%LOCAL_PATH%"
    echo [DELETED] %LOCAL_PATH%
)

:: 5. Delete Documents (Preferences & Auto-Saves)
echo [PHASE 5] Destroying Documents Settings...
set "DOCS_PATH=%USERPROFILE%\Documents\Adobe\Premiere Pro\25.0"
if exist "%DOCS_PATH%" (
    rmdir /S /Q "%DOCS_PATH%"
    echo [DELETED] %DOCS_PATH%
)

echo.
echo ===================================================
echo OPERATION COMPLETE: Premiere Pro 2025 has been completely wiped.
echo ===================================================
pause
