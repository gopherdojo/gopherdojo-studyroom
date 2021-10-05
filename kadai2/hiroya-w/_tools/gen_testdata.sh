#!/bin/sh
DIR=${1:-.}

curl https://avatars.githubusercontent.com/hiroya-w -o $DIR/image_png.png
curl http://icb-lab.naist.jp/members/yoshi/ouec_lecture/image_recognition/image_files/lena.jpg -o $DIR/image_jpg.jpg
curl https://upload.wikimedia.org/wikipedia/commons/2/2c/Rotating_earth_%28large%29.gif -o $DIR/image_gif.gif
