package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/r-uehara0219/gopherdojo-studyroom/imgconverter"
)

var option = imgconverter.Option{
	Input:  flag.String("i", ".jpeg", "input extension"),
	Output: flag.String("o", ".png", "output extension"),
}

func main() {
	flag.Parse()
	if !imgconverter.IsValidExtension(*option.Input, "input") ||
		!imgconverter.IsValidExtension(*option.Output, "output") {
		log.Fatalln(
			"Invalid extension has been specified.\n",
			"Please check README.md to see what extensions can be specified.",
		)
	}

	args := flag.Args()

	err := filepath.Walk(filepath.Dir(args[0]),
		func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == *option.Input {
				err = imgconverter.Do(path, &option)
				if err != nil {
					return err
				}
			}
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
}
