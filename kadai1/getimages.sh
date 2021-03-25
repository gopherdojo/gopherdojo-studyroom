#!/bin/sh
# reset and create nested images file

# delete old directory
if [ -e ./images ]; then
    echo delete old images folder...
    rm -rf images
fi

# make directory
echo make new images folder...
mkdir images

# download images
echo download images...
cd images && \
curl -O https://golang.org/doc/gopher/appenginegopher.jpg \
     -O https://golang.org/doc/gopher/bumper.png \
     -O https://golang.org/doc/gopher/bumper640x360.png \
     -O https://golang.org/doc/gopher/bumper640x360.png \
     -O https://golang.org/doc/gopher/fiveyears.jpg \
     -O https://golang.org/doc/gopher/modelsheet.jpg

mkdir dir
cp *.png *.jpg dir
