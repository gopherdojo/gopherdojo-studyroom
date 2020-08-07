#!/bin/sh

set -e

# -------------
# Environments
# -------------

RUN=$1
WORKING_DIR=$2
GITHUB_TOKEN=$3
COMMENT=""
SUCCESS=0

# if not set, assign default value
if [ "$2" = "" ]; then
    WORKING_DIR="kadai1/eimg"
fi

cd ${WORKING_DIR}
PKGNAME=$(go list ./...)

# ------------
# Functions
# ------------

init() {
    # make directory setting for gh-pages
    mkdir -p /gh-pages
    rm -rf /gh-pages/dist
    mkdir -p /gh-pages/dist
    
    # install config file for layout
    cp /go/src/golang.org/x/tools/godoc/static/jquery.js /gh-pages/dist/
    cp /go/src/golang.org/x/tools/godoc/static/godocs.jp /gh-pages/dist/
    cp /go/src/golang.org/x/tools/godoc/static/style.css /gh-pages/dist/
}

generate_godoc() {
    godoc -url 'http://localhost:8080/pkg/github.com/task4233/gopherdojo-studyroom/kadai1/eimg/' > /gh-pages/dist/index.html
}


deploy() {
    npm install
    npm run deploy
}

# -------------
# Main
# ------------
init()
generate_godoc()
deploy()

