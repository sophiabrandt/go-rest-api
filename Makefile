SHELL := /bin/bash

export PROJECT = go-rest-api

# ==============================================================================
# Development

run: dev

dev:
	go run ./cmd/server

# ==============================================================================
# Modules support

tidy:
	go mod tidy
	go mod vendor

deps-upgrade:
	go get -u -t -d -v ./...
	go mod tidy
	go mod vendor

deps-cleancache:
	go clean -modcache
