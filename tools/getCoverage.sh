#!/bin/sh

cd nn
go test -coverprofile=profile ./...
go tool cover -html=profile
