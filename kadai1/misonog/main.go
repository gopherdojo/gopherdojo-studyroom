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

	flag.StringVar(&path, "path", pwd, "Directory path to convert image file")
	flag.Parse()

	oldExt := ".png"
	newExt := ".jpg"
	if err := lib.ImgConv(path, oldExt, newExt); err != nil {
		log.Fatal(err)
	}
}
