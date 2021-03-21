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
	flag.StringVar(&oldExt, "o", ".jpg", "")
	flag.StringVar(&newExt, "n", ".png", "")
	flag.Parse()

	if err := lib.ImgConv(path, oldExt, newExt); err != nil {
		log.Fatal(err)
	}
}
