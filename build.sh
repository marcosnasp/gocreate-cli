#!/bin/bash
mkdir -p build
GOOS=linux GOARCH=amd64 go build -o build/gocreate-linux-amd64 ./cmd/gocreate
GOOS=windows GOARCH=amd64 go build -o build/gocreate-windows-amd64.exe ./cmd/gocreate
GOOS=darwin GOARCH=amd64 go build -o build/gocreate-darwin-amd64 ./cmd/gocreate
echo "Binários gerados em build/"