NAME := imgconv
VERSION := $(gobump show -r)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := "-X main.revision=$(REVISION)"

## Install dependencies
.PHONY: deps
deps:
	go get -v -d

## Setup
.PHONY: devel-deps
devel-deps: devel-deps
	go install golang.org/x/lint/golint@latest
	go install github.com/x-motemen/gobump/cmd/gobump@latest
	go install github.com/Songmu/make2help/cmd/make2help@latest

## Run tests
.PHONY: test
test: deps
	go test -v -race -cover -coverprofile=coverage.out ./...

.PHONY: test-deps
test-deps:
	cd _tools && sh gen_testdata.sh ../testdata

.PHONY: cover
cover:
	go tool cover -html=coverage.out

## Lint
.PHONY: lint
lint: devel-deps
	go vet ./...
	golint -set_exit_status ./...

## build binaries
bin/%: cmd/%/main.go deps
	go build -ldflags $(LDFLAGS) -o $@ $<

## build binary
.PHONY: build
build: bin/imgconv

## Show help
.PHONY: help
help:
	@make2help $(MAKEFILE_LIST)
