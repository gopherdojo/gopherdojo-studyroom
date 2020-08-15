package main

import (
	"flag"
	"log"
	"os"

	"github.com/shinnosuke-K/gopherdojo-studyroom/kadai1/shinnosuke-K/conv"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	dirPath := flag.String("path", pwd, "Directory path to convert image file's extension")
	before := flag.String("b", "jpeg", "Image format before change")
	after := flag.String("a", "png", "Image format after change")
	flag.Parse()

	conv.Do(*dirPath, *before, *after)

}
