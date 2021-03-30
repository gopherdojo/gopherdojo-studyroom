package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kassy11/gopherdojo-studyroom/kadai1/kassy11/convert"
)

//type ConvertImage struct {
//	inputFormat     string
//	outputFormat    string
//	inputDirectory  string
//	outputDirectory string
//}

//type FormatName struct {
//	Jpg string
//	Png string
//}
//
//var format FormatName
//
//func init() {
//	file, err := os.Open("config.json")
//	if err != nil {
//		fmt.Fprintln(os.Stderr, "Cannot open config file")
//		os.Exit(1)
//	}
//	decoder := json.NewDecoder(file)
//	format = FormatName{}
//	err = decoder.Decode(&format)
//	if err != nil {
//		fmt.Fprintln(os.Stderr, "Cannot get configuration from file")
//		os.Exit(1)
//	}
//}

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
		fmt.Fprintln(os.Stderr, "imgconv: no Directory specified")
		fmt.Fprintln(os.Stderr, "imgconv: try 'imgconv--help' for more information")
		os.Exit(1)
	}

	dir := flag.Arg(0)

	if (inputFormat != "jpg" && inputFormat != "png") || (outputFormat != "jpg" && outputFormat != "png") {
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
