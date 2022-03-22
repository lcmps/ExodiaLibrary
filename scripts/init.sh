#!/bin/bash

CGO_ENABLED=0 go build -ldflags="-s -w" -o exodialib-core -v .
docker build -t exodialib-core:latest -f Dockerfile .
docker-compose down
docker compose -f docker-compose.yaml up --remove-orphans