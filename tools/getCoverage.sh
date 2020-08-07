#!/bin/sh

cd kadai1/task4233/eimg
go test -coverprofile=profile ./...
go tool cover -html=profile
