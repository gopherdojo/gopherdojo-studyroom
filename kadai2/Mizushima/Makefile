NAME := bin/converter

all:
	go build -o $(NAME)

test: clean all 
	cd ./picconvert; go test -v -cover

clean:
	rm -rf ./testdata/*_converted.*
	rm -rf ./testdata/testdata/*_converted.*
	rm -rf ./bin/*

.PHONY: test clean
