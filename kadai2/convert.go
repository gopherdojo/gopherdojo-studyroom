package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/dai65527/gopherdojo-studyroom/kadai1/imgcnv"
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

	// check input flag is valid
	switch *flgInput {
	case "jpg", "jpeg", "png":
		// nop
	default: // invalid extension
		fmt.Fprintln(os.Stderr, "error: "+*flgInput+": invalid input flag (should be jpg or png)")
		os.Exit(1)
	}

	// check output flag is valid
	switch *flgOutput {
	case "jpg", "jpeg", "png":
		// nop
	default: // invalid extension
		fmt.Fprintln(os.Stderr, "error: "+*flgOutput+": invalid output flag (should be jpg or png)")
		os.Exit(1)
	}

	// walk in directory and convert image files
	err := filepath.Walk(flag.Args()[0], func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || filepath.Ext(path) != "."+*flgInput {
			return nil
		}
		erri := imgcnv.Convert(path, imgcnv.Extension(*flgInput), imgcnv.Extension(*flgOutput))
		if erri != nil {
			fmt.Fprintln(os.Stdout, "error: "+path+":", erri)
		}
		return nil
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
