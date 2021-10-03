#!/bin/sh
curl https://avatars.githubusercontent.com/hiroya-w -o hiroya-w.png

for i in {1..2}
do
    cp hiroya-w.png ../testdata/image$i.png
done

rm hiroya-w.png
