package main

import (
	"flag"
	"gopher-dojo/kadai1/wada/convertor"
	"log"
	"os"
)

var beforeExt, afterExt string

func main() {
	flag.StringVar(&beforeExt, "before", "jpg", "Extension before conversion")
	flag.StringVar(&afterExt, "after", "png", "Extension after conversion")
	flag.Parse()

	validator(beforeExt, afterExt)

	c := convertor.NewConvertor(flag.Args()[0], beforeExt, afterExt)
	err := c.ConvertImage()
	if err != nil {
		log.Fatal(err)
	}
}

// Returns an error if there is a problem with a program argument
func validator(beforeExt, afterExt string) {
	if len(flag.Args()) == 0 {
		log.Fatal("Please enter an argument")
	}

	if dir, err := os.Stat(flag.Args()[0]); os.IsNotExist(err) || !dir.IsDir() {
		log.Fatal("The argument must be a directory")
	}

	if !checkExt(beforeExt) && !checkExt(afterExt) {
		log.Fatal("The extension must be jpeg, jpg, png, or gif")
	}

	if beforeExt == afterExt {
		log.Fatal("Please specify a different extension before and after conversion")
	}
}

// Check if the entered extension is a supported format
func checkExt(flagExt string) bool {
	exts := []string{"jpeg", "jpg", "png", "gif"}
	for _, ext := range exts {
		if ext == flagExt {
			return true
		}
	}
	return false
}
