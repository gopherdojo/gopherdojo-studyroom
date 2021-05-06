package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/edm20627/gopherdojo-studyroom/kadai1/edm20627/imageconvert"
)

var ci = imageconvert.ConvertImage{}

func init() {
	flag.StringVar(&ci.From, "-from", "jpg", "Specify the image format before conversion")
	flag.StringVar(&ci.From, "f", "jpg", "Specify the image format before conversion (short)")
	flag.StringVar(&ci.To, "-to", "png", "Specify the converted image format")
	flag.StringVar(&ci.To, "t", "png", "Specify the converted image format (short)")
	flag.BoolVar(&ci.DeleteOption, "-delete", false, "Delete the image before conversion")
	flag.BoolVar(&ci.DeleteOption, "d", false, "Delete the image before conversion (short)")
}

func main() {
	flag.Parse()
	dirs := flag.Args()

	if !ci.Valid() {
		fmt.Fprintln(os.Stderr, "supported formt is "+strings.Join(imageconvert.SupportedFormat, ", "))
		os.Exit(1)
	}

	err := ci.Get(dirs)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = ci.Convert()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
