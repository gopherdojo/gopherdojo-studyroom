#! /bin/bash

ROOT_PKG=$(cd $(dirname $0); pwd)
TEST_PROFILE_DIR=${ROOT_PKG}/testprofile

# mainのテスト
go test -coverprofile=${TEST_PROFILE_DIR}/main ${ROOT_PKG}
go tool cover -html=${TEST_PROFILE_DIR}/main

# convertorのテスト
go test -coverprofile=${TEST_PROFILE_DIR}/convertor ${ROOT_PKG}/convertor
go tool cover -html=${TEST_PROFILE_DIR}/convertor