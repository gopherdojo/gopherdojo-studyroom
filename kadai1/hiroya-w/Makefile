.PHONY: build
build:
	go build -o bin/imgconf main.go

.PHONY: run
run:
	./bin/imgconf -input-type=png -output-type=jpg testdata

.PHONY: clean
clean:
	rm -rf bin/main
