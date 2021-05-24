NAME := bin/converter

all:
	go build -o $(NAME)

test: clean all
	./test.sh ./bin/converter ./testdata

clean:
	rm -rf converter ./testdata/*_converted.*
	rm -rf ./testdata/testdata/*_converted.*

.PHONY: test clean
