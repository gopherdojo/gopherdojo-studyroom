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
    WORKING_DIR="kadai1/task4233/eimg"
fi

cd ${WORKING_DIR}
DOCNAME="doc"
PKGNAME=$(go list ./...)

# ------------
# Functions
# ------------

init() {
    # make directory setting for gh-pages
    echo "make directory setting for gh-pages"
    rm -rf ./${DOCNAME}
    mkdir ./${DOCNAME}
    
    # install config file for layout
    echo "install config file for layout"
    cp ${GOPATH}/pkg/mod/golang.org/x/tools@*/godoc/static/jquery.js ./${DOCNAME}/
    cp ${GOPATH}/pkg/mod/golang.org/x/tools@*/godoc/static/godocs.js ./${DOCNAME}/
    cp ${GOPATH}/pkg/mod/golang.org/x/tools@*/godoc/static/style.css ./${DOCNAME}/
}

generate_godoc() {
    godoc -url 'http://localhost:8080/pkg/github.com/task4233/gopherdojo-studyroom/kadai1/task4233/eimg/' | sed -e '1d' | sed 's/\/lib\/godoc/\./g' > ./${DOCNAME}/index.html
}

# -------------
# Main
# ------------
init
generate_godoc


exit ${SUCCESS}

