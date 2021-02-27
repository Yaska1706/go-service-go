#!/bin/sh
echo "running go.sh..."

echo "running go vet for whole project"
go vet ./...

echo "running go fmt for whole project"
go fmt ./...

echo "sourcing env variables"
source ./.env

echo "go run go"
go run cmd/server/main.go
