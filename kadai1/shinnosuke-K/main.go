package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dirPath := flag.String("path", pwd, "Directory path to convert image file's extension")
	flag.Parse()
	fmt.Println(*dirPath)
}
