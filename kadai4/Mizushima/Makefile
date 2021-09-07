BINARY_NAME := bin/omikuji
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFMT=$(GOCMD) fmt

all: build test

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

test: build
	cd fortune; $(GOTEST) -v -cover
	$(GOTEST) -v -cover

clean:
	$(GOCLEAN)

format:
	$(GOFMT) ./...

coverage:
	$(GOTEST) -coverprofile=profile . ./fortune
	$(GOCMD) tool cover -html=profile

.PHONY: test clean format