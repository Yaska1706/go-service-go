#!/bin/sh

go vet ./...
go fmt ./...

source ./.env

go run cmd/server/main.go
