@echo off
cd D:\test-go-tutti\gestionlogistica-go
:loop
go run main.go config.go models.go handlers.go
if %ERRORLEVEL% neq 0 (
    timeout /t 5
    goto loop
)
