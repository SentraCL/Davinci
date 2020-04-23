
@echo off
if exist davinci.exe (
    del davinci.exe
)
go build ..\DaVinci.BackeEnd\davinci.go
start davinci.exe -p 1123 -c "..\DaVinci.BackeEnd\davinci.cfg"