BINARY_NAME=imgconv

all: create test build

build:
	go build -o $(BINARY_NAME)

test:
	go test -v ./...

create:
	@sh init.sh

clean:
	@rm -r testdata
