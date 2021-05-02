package main

import (
	"flag"
	"log"
	"os"
	"time"
)

const timeout = 10 * time.Second

func main() {
	var targetDir string
	var timeout time.Duration

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&targetDir, "d", pwd, "path to the directory to save the downloaded file, filename will be taken from url")
	flag.DurationVar(&timeout, "t", timeout, "timeout of checking request in seconds")
	flag.Parse()

	cli := New()
	if err := cli.Run(flag.Args(), targetDir, timeout); err != nil {
		log.Fatal(err)
	}
}
