NAME := converter

all:
	go build -o $(NAME)

test: clean all
	./test.sh ./converter ./test

clean:
	rm -rf converter ./test/*_converted.*
	rm -rf ./test/test/*_converted.*

.PHONY: test clean
