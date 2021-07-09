BINARY_NAME := bin/typing
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

all: build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

test: clean all
	$(GOTEST) -v -cover
	cd ./typing; $(GOTEST) -v -cover

clean:
	$(GOCLEAN)
	rm -rf $(BINARY_NAME)

.PHONY: test clean
