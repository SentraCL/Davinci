
@echo off
if exist davinci.exe (
    del davinci.exe
)
go build ..\DaVinci.BackEnd\davinci.go
start davinci.exe -p 1123 -c "..\DaVinci.BackEnd\davinci.cfg"