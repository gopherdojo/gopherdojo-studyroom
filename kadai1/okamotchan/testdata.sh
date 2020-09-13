#!/bin/zsh

rm -r testdata

mkdir testdata
mkdir testdata/img

curl http://flat-icon-design.com/f/f_object_174/s512_f_object_174_0bg.jpg > ./testdata/azarashi.jpg
curl http://flat-icon-design.com/f/f_object_174/s512_f_object_174_0bg.jpg > ./testdata/img/azarashi.jpg

curl http://flat-icon-design.com/f/f_object_149/s512_f_object_149_0bg.jpg > ./testdata/tanuki.jpg
curl http://flat-icon-design.com/f/f_object_149/s512_f_object_149_0bg.jpg > ./testdata/img/tanuki.jpg

curl http://flat-icon-design.com/f/f_object_157/s512_f_object_157_0bg.png > ./testdata/osaru.png