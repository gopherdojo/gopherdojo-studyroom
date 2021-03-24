package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/dai65527/gopherdojo-studyroom/kadai1/imageConverter"
)

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		fmt.Fprintln(os.Stderr, "error: invalid argument")
		os.Exit(1)
	} else if _, err := os.Stat(flag.Args()[0]); err != nil {
		fmt.Fprintln(os.Stderr, "error: "+flag.Args()[0]+": no such file or directory")
		os.Exit(1)
	}
	convert(flag.Args()[0])
}

func convert(dirpath string) {
	err := filepath.Walk(flag.Args()[0], func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			imageConverter.Jpg2png(path)
		}
		return nil
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
