#!/bin/bash

echo "Building for Windows..."
GOOS=windows GOARCH=amd64 go build -o govWin.exe
echo "Windows build done."

echo "Building for Linux..."
GOOS=linux GOARCH=amd64 go build -o govLin
echo "Linux build done."

echo "Building for MacOS..."
GOOS=darwin GOARCH=amd64 go build -o govMac
echo "MacOS build done."

echo "All builds complete."
