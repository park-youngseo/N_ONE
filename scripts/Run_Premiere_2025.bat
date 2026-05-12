@echo off
set "PREM_PATH=C:\Program Files\Adobe\Adobe Premiere Pro 2025\Adobe Premiere Pro.exe"
if exist "%PREM_PATH%" (
    echo [INFO] Launching Premiere Pro 2025...
    start "" "%PREM_PATH%"
) else (
    echo [ERROR] Premiere Pro 2025 not found at:
    echo %PREM_PATH%
    echo.
    echo Please check if the folder was accidentally deleted or moved.
    pause
)
