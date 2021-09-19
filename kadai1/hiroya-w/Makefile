SRC := main.go imgconv/imgconv.go

.PHONY: build
build: bin/imgconv

bin/imgconv: $(SRC)
	go build -o bin/imgconv .

.PHONY: run
run:
	./bin/imgconv -input-type=png -output-type=jpg testdata

.PHONY: test
test: build testdata-gen run

.PHONY: testdata-gen
testdata-gen:
	cd testdata && ./init.sh

.PHONY: check
check:
	errcheck ./...

.PHONY: clean
clean:
	-rm bin/imgconv
	-rm testdata/*.jpg testdata/*jpeg testdata/*.png testdata/*.gif
