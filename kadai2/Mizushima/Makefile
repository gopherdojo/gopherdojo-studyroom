# 
# 原典：https://frasco.io/golang-dont-afraid-of-makefiles-785f3ec7eb32
# 

BINARY_NAME := bin/converter
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
	cd ./picconvert; $(GOTEST) -v -cover

clean:
	$(GOCLEAN)
	rm -rf $(BINARY_NAME)

.PHONY: test clean
