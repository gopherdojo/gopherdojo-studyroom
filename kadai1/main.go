package main

import (
	"flag"
	"fmt"
	"github.com/r-uehara0219/gopherdojo-studyroom/imgconverter"
	"os"
	"path/filepath"
)

var validInputExtension = []string{
	".jpeg",
	".jpg",
	".png",
	".gif",
}
var validOutputExtension = []string{
	".jpeg",
	".jpg",
	".png",
}

func isValidExtension(s string, arr []string) bool {
	for _, extension := range arr {
		if s == extension {
			return true
		}
	}
	return false
}

var option = imgconverter.Option{
	Input:  flag.String("i", ".jpeg", "input extension"),
	Output: flag.String("o", ".png", "output extension"),
}

func main() {
	flag.Parse()
	if !isValidExtension(*option.Input, validInputExtension) || !isValidExtension(*option.Output, validOutputExtension) {
		fmt.Fprintln(os.Stderr, "Invalid extension has been specified.")
		fmt.Fprintln(os.Stderr, "Please check the document to see what extensions can be specified.")
		os.Exit(1)
	}

	args := flag.Args()

	err := filepath.Walk(filepath.Dir(args[0]),
		func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == *option.Input {
				err = imgconverter.Do(path, option)
				if err != nil {
					return err
				}
			}
			return nil
		})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
