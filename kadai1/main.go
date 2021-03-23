package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"./imageConverter"
)

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		fmt.Fprintln(os.Stderr, "error: invalid argument")
		return
	}
	convert(flag.Args()[0])
}

func convert(dirpath string) {
	err := filepath.Walk(flag.Args()[0], func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			if path != dirpath {
				convert(path)
			}
		} else {
			fmt.Println(path)
			imageConverter.Jpg2png(path)
		}
		return nil
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
