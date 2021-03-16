package main

import (
	"flag"
	"fmt"
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

	if flg := lib.ExistDir(path); flg {
		fmt.Printf("Selected Dir is: %v\n", path)
	} else {
		fmt.Println("Error.")
	}
}
