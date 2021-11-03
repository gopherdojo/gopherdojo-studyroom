#!/bin/sh
rm -rf testdata
mkdir testdata
mkdir testdata/fuga

for i in {1..9}
do
    cp sample/morpeco-sample.jpeg testdata/hoge"$i".jpg
done

for i in {1..9}
do
    cp sample/morpeco-sample.jpeg testdata/fuga/hoge"$i".jpeg
done
