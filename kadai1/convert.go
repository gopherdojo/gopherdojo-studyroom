package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/dai65527/gopherdojo-studyroom/kadai1/imageConverter"
)

func main() {
	// flags
	var flgInput = flag.String("i", "jpg", "input type (jpg or png, default: jpg)")
	var flgOutput = flag.String("o", "png", "jpg or png (jpg or png, default: png)")
	flag.Parse()

	// check input
	if len(flag.Args()) != 1 {
		fmt.Fprintln(os.Stderr, "error: invalid argument")
		os.Exit(1)
	} else if _, err := os.Stat(flag.Args()[0]); err != nil {
		fmt.Fprintln(os.Stderr, "error: "+flag.Args()[0]+": no such file or directory")
		os.Exit(1)
	}

	// check flag is valid
	if *flgInput != "jpg" && *flgInput != "jpeg" && *flgInput != "png" {
		fmt.Fprintln(os.Stderr, "error: "+*flgInput+": invalid input flag (should be jpg or png)")
		os.Exit(1)
	} else if *flgOutput != "jpg" && *flgOutput != "jpeg" && *flgOutput != "png" {
		fmt.Fprintln(os.Stderr, "error: "+*flgOutput+": invalid output flag (should be jpg or png)")
		os.Exit(1)
	}

	// convert all
	convert(flag.Args()[0], *flgInput, *flgOutput)
}

func convert(dirpath string, inputFileExt string, outputFileExt string) {
	err := filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && filepath.Ext(path) == "."+inputFileExt {
			erri := imageConverter.Convert(path, imageConverter.Extension(inputFileExt), imageConverter.Extension(outputFileExt))
			if erri != nil {
				fmt.Fprintln(os.Stdout, "error: "+path+": ", err)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
