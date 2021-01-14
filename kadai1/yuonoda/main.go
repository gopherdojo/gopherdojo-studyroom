package main

import (
	"flag"
	"github.com/yuonoda/gopherdojo-studyroom/kadai1/yuonoda/lib"
)

var fromFmt = flag.String("from", "jpg", "Specify a format of your original image")
var toFmt = flag.String("to", "png", "Specify a format which you want to convert the image to")
var path = flag.String("path", ".", "Path to a directory in which images will be converted recursively")

func main() {
	convImages.Run(*fromFmt, *toFmt, *path)
	flag.Parse()
}
