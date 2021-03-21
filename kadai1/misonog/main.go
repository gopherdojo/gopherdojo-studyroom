package main

import (
	"flag"
	"log"
	"os"

	"github.com/misonog/gopherdojo-studyroom/kadai1/misonog/lib"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	var path string
	var oldExt string
	var newExt string

	flag.StringVar(&path, "path", pwd, "Directory path to convert image file")
	flag.StringVar(&oldExt, "o", ".jpg", "Image format before change")
	flag.StringVar(&newExt, "n", ".png", "Image format after change")
	flag.Parse()

	if err := lib.ImgConv(path, oldExt, newExt); err != nil {
		log.Fatal(err)
	}
}
