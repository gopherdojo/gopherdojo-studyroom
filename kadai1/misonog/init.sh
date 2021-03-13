#!/bin/sh

# Create testdata dir
mkdir -p testdata/jpg
mkdir -p testdata/jpeg
mkdir -p testdata/png

# Create sample testdata
createSample() {
    for i in $(seq 1 3)
    do
        cp sample.png "$1"/"$i"."$2"
    done
}

createSample "testdata/jpg" "jpg"
createSample "testdata/jpeg" "jpeg"
createSample "testdata/png" "png"
createSample "testdata" "gif"
