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

	var targetDir string

	flag.StringVar(&targetDir, "d", pwd, "path to the directory to save the downloaded file, filename will be taken from url")
	flag.Parse()

	// if err := download("header.jpg", targetDir, flag.Arg(0)); err != nil {
	// 	log.Fatal(err)
	// }
}
