#!/bin/bash
DIR=${1:-.}

go install github.com/swaggo/swag/cmd/swag@latest
swag init -g ./cmd/web/main.go -o ./cmd/web/docs
