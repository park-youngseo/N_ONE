@echo off
setlocal enabledelayedexpansion
:: ==========================================
:: NONE-OS Adobe Hosts Blocker
:: ==========================================
color 0B
echo Requesting Administrative Privileges...
net session >nul 2>&1
if %errorLevel% NEQ 0 (
    echo [ERROR] Administrator rights are required to modify the hosts file.
    echo Please right-click this file and select "Run as administrator".
    pause
    exit
)

set HOSTS_FILE=%WINDIR%\System32\drivers\etc\hosts
set BACKUP_FILE=%WINDIR%\System32\drivers\etc\hosts.bak

echo.
echo ===================================================
echo Initiating Adobe Servers Block via Hosts File
echo ===================================================
echo.

echo [PHASE 1] Backing up current hosts file...
copy /Y "%HOSTS_FILE%" "%BACKUP_FILE%" >nul
echo [SUCCESS] Backup created at: %BACKUP_FILE%
echo.

echo [PHASE 2] Appending Adobe Block List...

:: Adding an empty line first to ensure it doesn't merge with the last line
echo. >> "%HOSTS_FILE%"
echo # ======================================= >> "%HOSTS_FILE%"
echo # Adobe Premiere Pro Block List >> "%HOSTS_FILE%"
echo # ======================================= >> "%HOSTS_FILE%"

:: Common Adobe block domains
set "ADOBE_DOMAINS=lmlicenses.wip4.adobe.com lm.licenses.adobe.com na1r.services.adobe.com hlrcv.stage.dps.local practivate.adobe.com activate.adobe.com 3dms.adobe.com crl.adobe.com hl2rcv.adobe.com genuine.adobe.com cc-api-data.adobe.com ic.adobe.io p13n.adobe.io 192.150.14.69 192.150.18.101"

for %%A in (%ADOBE_DOMAINS%) do (
    echo 127.0.0.1 %%A >> "%HOSTS_FILE%"
)

echo [SUCCESS] Hosts file has been updated!
echo.

echo [PHASE 3] Flushing DNS Cache...
ipconfig /flushdns >nul
echo [SUCCESS] DNS Cache Flushed.
echo.

echo ===================================================
echo OPERATION COMPLETE: Adobe servers are now blocked.
echo ===================================================
pause
