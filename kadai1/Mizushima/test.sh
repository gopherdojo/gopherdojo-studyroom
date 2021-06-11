#!/bin/bash

#  "jpeg" "jpg" "png" "gif"
echo "testing..."

exe=$1
path=$2
list=("jpeg"
      "png"
      "gif"); #  "jpg"

$exe $path

for ext1 in ${list[@]}; do
  for ext2 in ${list[@]}; do
    if [ ! $ext1 = $ext2 ]; then
      $exe -pre $ext1 -post $ext2 $path
      converted="${path}/test01_converted.${ext2}"
      if [ -e $converted ]; then
        echo "${exe} -pre ${ext1} -post ${ext2} ${path} ...passed"
      else 
        echo "${exe} -pre ${ext1} -post ${ext2} ${path} ...failed"
      fi
    fi
  done
done
