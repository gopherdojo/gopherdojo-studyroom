package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kassy11/gopherdojo-studyroom/kadai1/kassy11/convert"
)

var format *convert.FormatType

func init() {
	format = convert.LoadConfig()
}

func main() {
	var img convert.ConvertImage

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: imgconv[options..] <directory>")
		fmt.Fprintln(os.Stderr, "Options")
		flag.PrintDefaults()
	}

	var input string
	var output string
	var outputDir string
	flag.StringVar(&input, "i", format.Jpg, "-i <format(jpg or jpeg or png))>")
	flag.StringVar(&output, "f", format.Png, "-f <format(jpg or jpeg or png)>")
	flag.StringVar(&outputDir, "o", "output", "-o <directory name>")
	flag.Parse()

	if len(flag.Args()) <= 0 {
		fmt.Fprintln(os.Stderr, "imgconv: no Directory specified")
		fmt.Fprintln(os.Stderr, "imgconv: try 'imgconv--help' for more information")
		os.Exit(1)
	}

	dir := flag.Arg(0)

	if (input != format.Jpg && input != format.Jpeg && input != format.Png) || (output != format.Jpg && output != format.Jpeg && output != format.Png) {
		fmt.Fprintln(os.Stderr, "imgconv: Invalid image format")
		fmt.Fprintln(os.Stderr, "imgconv: try 'imgconv--help' for more information")
		os.Exit(1)
	} else if input == output {
		fmt.Fprintln(os.Stderr, "imgconv: Incorrect parameter combination")
		fmt.Fprintln(os.Stderr, "imgconv: ry 'imgconv--help' for more information")
		os.Exit(1)
	}

	img = convert.ConvertImage{InputFormat: input, OutputFormat: output, InputDirectory: dir, OutputDirectory: outputDir}
	img.Do()
}
