@echo off
REM Build script for Todo List CLI on Windows

echo Building Todo List CLI...

go build -o todo-list-cli.exe main.go

if %ERRORLEVEL% EQU 0 (
    echo Build successful!
    echo Run with: todo-list-cli.exe help
) else (
    echo Build failed!
    exit /b 1
)
