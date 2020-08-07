#!/bin/sh

cd kadai1/eimg
go test -coverprofile=profile ./...
go tool cover -html=profile
