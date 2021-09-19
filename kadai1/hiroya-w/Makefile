SRC := main.go imgconv/imgconv.go

.PHONY: build
build: bin/imgconv

bin/imgconv: $(SRC)
	go build -o bin/imgconv .

.PHONY: run
run:
	./bin/imgconv -input-type=png -output-type=jpg testdata

.PHONY: test
test: build run

.PHONY: clean
clean:
	-rm bin/imgconv
	-rm testdata/*.jpg testdata/*jpeg testdata/*.png testdata/*.gif
