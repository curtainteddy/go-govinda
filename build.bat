@echo off
echo Building for Windows...
set GOOS=windows
set GOARCH=amd64
go build -o govWin.exe
echo Windows build done.

echo Building for Linux...
set GOOS=linux
set GOARCH=amd64
go build -o govLin
echo Linux build done.

echo Building for MacOS...
set GOOS=darwin
set GOARCH=amd64
go build -o govMac
echo MacOS build done.

echo All builds complete.
pause
