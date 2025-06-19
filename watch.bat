@echo off
set TARGET=C:\path\to\watch
set COMMAND=echo Jalankan program
set POLLING_INTERVAL=1

set PID=

:start
if not "%PID%"=="" (
    taskkill /pid %PID% /f >nul 2>&1
)

start "" /b cmd /c %COMMAND%
for /f "tokens=2 delims==" %%a in ('wmic process where "name='cmd.exe'" get ProcessId /value ^| find "ProcessId"') do set PID=%%a

:loop
timeout /t %POLLING_INTERVAL% >nul

dir /s "%TARGET%" | findstr "Date" >nul
if errorlevel 1 goto loop

taskkill /pid %PID% /f >nul 2>&1
goto start