#!/bin/sh

# Create testdata dir
mkdir -p testdata/jpg
mkdir -p testdata/jpeg
mkdir -p testdata/png/recursive

# Create sample testdata
createSample() {
    for i in $(seq 1 3)
    do
        convert sample.png "$1"/"$i"."$2"
    done
}

createSample "testdata/jpg" "jpg"
createSample "testdata/jpeg" "jpeg"
createSample "testdata/png" "png"
createSample "testdata" "gif"
cp sample.png testdata/png/recursive/4.png
