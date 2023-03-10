.PHONY: test build run install
SHELL := bash
.ONESHELL:
.EXPORT_ALL_VARIABLES:
.SHELLFLAGS := -eu -o pipefail -c

DEBUG=true

pre:
	go mod tidy

install:
	go install golang.org/x/vuln/cmd/govulncheck@latest

test:
	go test -race -v --coverprofile=coverage.out ./...
	go tool cover -func coverage.out | grep total

format:
	gofumpt -l -w .

bench:
	go test -bench=. -benchmem ./...

vuln:
	govulncheck ./...

lint:
	golangci-lint run

cov: test
	go tool cover -html=coverage.out
