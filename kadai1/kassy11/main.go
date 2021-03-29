package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kassy11/gopherdojo-studyroom/kadai1/kassy11/convert"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: imgconv[options..] <directory>")
		fmt.Fprintln(os.Stderr, "Options")
		flag.PrintDefaults()
	}

	var inputFormat string
	var outputFormat string
	var outputDirectory string
	flag.StringVar(&inputFormat, "i", "jpg", "-i <format(jpg or png))>")
	flag.StringVar(&outputFormat, "f", "png", "-f <format(jpg or png)>")
	flag.StringVar(&outputDirectory, "o", "output", "-o <directory name>")
	flag.Parse()

	if len(flag.Args()) <= 0 {
		fmt.Fprintln(os.Stderr, "no Directory specified")
		fmt.Fprintln(os.Stderr, "try 'imgconv--help' for more information")
		os.Exit(1)
	}

	dir := flag.Arg(0)

	if (inputFormat != "jpg" && inputFormat != "png") || (outputFormat != "jpg" && outputFormat != "png") {
		fmt.Fprintln(os.Stderr, "Invalid image format")
		fmt.Fprintln(os.Stderr, "try 'imgconv--help' for more information")
		os.Exit(1)
	} else if inputFormat == outputFormat {
		fmt.Fprintln(os.Stderr, "Incorrect parameter combination")
		fmt.Fprintln(os.Stderr, "try 'imgconv--help' for more information")
		os.Exit(1)
	}
	convert.Do(dir, inputFormat, outputDirectory)
}
