#!/bin/bash
DIR=${1:-.}

go install github.com/swaggo/swag/cmd/swag@latest
swag init -g ./apps/web/main.go -o ./docs