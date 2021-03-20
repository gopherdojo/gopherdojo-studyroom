package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	var path string

	flag.StringVar(&path, "path", pwd, "Directory path to convert image file")
	flag.Parse()
}
