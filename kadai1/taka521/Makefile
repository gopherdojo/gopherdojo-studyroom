NAME := conv

.PHONY: build
build: clean ## build go binary
	go build -o "$(NAME).exe" ./cmd/main.go

.PHONY: clean
clean: ## remove go binary
	@rm -f *.exe

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
