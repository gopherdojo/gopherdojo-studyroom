init:
	go get -v -d

jpegdata:
	cp -R ./testdata/src/jpeg/* ./testdata/dst

pngdata:
	cp -R ./testdata/src/png/* ./testdata/dst

cleandata:
	rm -rf ./testdata/dst/*

test:
	cp -R ./testdata/src/* ./testdata/dst
	go test -v -race -cover -coverprofile=coverage.out ./...
	rm -rf ./testdata/dst/*

.PHONY: doc
doc:
	godoc -http=:8080

.PHONY: cover
cover:
	go tool cover -html=coverage.out
