#! /bin/sh

ROOT_PKG=$(cd $(dirname $0); pwd)
TEST_PROFILE_DIR=${ROOT_PKG}/testprofile

# mainのテスト
go test -coverprofile=${TEST_PROFILE_DIR}/imgconv ${ROOT_PKG}/imgconv
go tool cover -html=${TEST_PROFILE_DIR}/imgconv

# convertorのテスト
go test -coverprofile=${TEST_PROFILE_DIR}/converter ${ROOT_PKG}/converter
go tool cover -html=${TEST_PROFILE_DIR}/converter