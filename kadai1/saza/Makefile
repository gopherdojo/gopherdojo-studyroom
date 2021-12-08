.PHONY: build
build:
	@go build -o bin/imgconv .

.PHONY: run
run: ready build
	@./bin/imgconv testdata

.PHONY: ready
ready:
	@./init.sh
