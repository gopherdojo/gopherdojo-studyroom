package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dirPath := flag.String("path", pwd, "Directory path to convert image file's extension")
	flag.Parse()

	files, err := ioutil.ReadDir(*dirPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name(), f.IsDir(), filepath.Ext(f.Name()))
	}
}
