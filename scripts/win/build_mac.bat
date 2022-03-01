@echo off
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -o app ../../