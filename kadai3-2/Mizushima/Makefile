BINARY_NAME := bin/paraDW
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

all: build test

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

test:
	$(BINARY_NAME) https://golang.org/dl/go1.16.5.linux-amd64.tar.gz -o ./test_download

clean:
	$(GOCLEAN)
	rm -rf $(BINARY_NAME) test_download/*

.PHONY: test clean
