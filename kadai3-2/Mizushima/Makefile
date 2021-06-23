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
	$(BINARY_NAME) http://i.imgur.com/z4d4kWk.jpg -o ./test_download
	$(BINARY_NAME) http://socia-enterprise.co.jp -o ./test_download
	$(BINARY_NAME) http://i.imgur.com/z4d4kWk.jpg http://socia-enterprise.co.jp -o ./test_download

clean:
	$(GOCLEAN)
	rm -rf $(BINARY_NAME) test_download/*

.PHONY: test clean
