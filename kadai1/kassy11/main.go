package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kassy11/gopherdojo-studyroom/kadai1/kassy11/convert"
)

type ConvertImage struct {
	inputFormat     string
	outputFormat    string
	inputDirectory  string
	outputDirectory string
}

var format *convert.FormatType

func init() {
	format = convert.LoadConfig()
}

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: imgconv[options..] <directory>")
		fmt.Fprintln(os.Stderr, "Options")
		flag.PrintDefaults()
	}

	var inputFormat string
	var outputFormat string
	var outputDirectory string
	flag.StringVar(&inputFormat, "i", format.Jpg, "-i <format(jpg or jpeg or png))>")
	flag.StringVar(&outputFormat, "f", format.Png, "-f <format(jpg or jpeg or png)>")
	flag.StringVar(&outputDirectory, "o", "output", "-o <directory name>")
	flag.Parse()

	if len(flag.Args()) <= 0 {
		fmt.Fprintln(os.Stderr, "imgconv: no Directory specified")
		fmt.Fprintln(os.Stderr, "imgconv: try 'imgconv--help' for more information")
		os.Exit(1)
	}

	dir := flag.Arg(0)

	if (inputFormat != format.Jpg && inputFormat != format.Jpeg && inputFormat != format.Png) || (outputFormat != format.Jpg && outputFormat != format.Jpeg && outputFormat != format.Png) {
		fmt.Fprintln(os.Stderr, "imgconv: Invalid image format")
		fmt.Fprintln(os.Stderr, "imgconv: try 'imgconv--help' for more information")
		os.Exit(1)
	} else if inputFormat == outputFormat {
		fmt.Fprintln(os.Stderr, "imgconv: Incorrect parameter combination")
		fmt.Fprintln(os.Stderr, "imgconv: ry 'imgconv--help' for more information")
		os.Exit(1)
	}
	convert.Do(dir, outputDirectory, inputFormat, outputFormat)
}
