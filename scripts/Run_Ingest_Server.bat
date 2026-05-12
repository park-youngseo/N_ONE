@echo off
title NONE-OS Ingest Server (Persistent)
set SCRIPT_PATH=projects\NONE_AI\scripts\ingest_server.py
set LOG_FILE=projects\NONE_AI\scripts\ingest_server.log

echo [*] Starting Ingest Server...
echo [*] Logs will be saved to %LOG_FILE%

:loop
echo [%date% %time%] Starting %SCRIPT_PATH% >> %LOG_FILE%
python "%SCRIPT_PATH%" >> %LOG_FILE% 2>&1
echo [%date% %time%] Server crashed or stopped. Restarting in 5 seconds... >> %LOG_FILE%
echo [!] Server crashed or stopped. Restarting in 5 seconds...
timeout /t 5
goto loop
