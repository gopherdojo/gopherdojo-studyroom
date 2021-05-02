package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	var targetDir string
	var timeout int

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&targetDir, "d", pwd, "path to the directory to save the downloaded file, filename will be taken from url")
	flag.IntVar(&timeout, "t", TIMEOUT, "timeout of checking request in seconds")
	flag.Parse()

	cli := New()
	if err := cli.Run(flag.Args(), targetDir, timeout); err != nil {
		log.Fatal(err)
	}
}
