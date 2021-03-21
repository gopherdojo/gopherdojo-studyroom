package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kimuson13/gopherdojo-studyroom/kimuson13/conversion"
)

var cs conversion.ConvertStruct

func init() {
	flag.StringVar(&cs.Before, "b", "jpg", "select extension of conversion source")
	flag.StringVar(&cs.After, "a", "png", "select extension you want to convert")
}

func main() {
	flag.Parse()
	dirs := flag.Args()

	err := conversion.ExtensionCheck(cs.Before)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Extension Error")
		os.Exit(1)
	}

	err = cs.WalkDirs(dirs)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Directory walk Error")
		os.Exit(1)
	}

	err = cs.Convert()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Convert Error")
		os.Exit(1)
	}
}
