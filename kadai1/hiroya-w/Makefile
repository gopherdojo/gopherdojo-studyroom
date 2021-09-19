.PHONY: build
build:
	go build -o bin/imgconf main.go

.PHONY: run
run:
	./bin/imgconf -input-type=png -output-type=jpg testdata

.PHONY: test
test: build run

.PHONY: clean
clean:
	-rm -rf bin/main
	-rm testdata/*.jpg testdata/*jpeg testdata/*.png testdata/*.gif
