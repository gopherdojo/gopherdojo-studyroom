package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"github.com/edm20627/gopherdojo-studyroom/kadai1/edm20627/imageconvert"
)

var supportedFormat = []string{"png", "jpg", "jpeg", "gif"}

var ci = imageconvert.ConvertImage{}

func init() {
	flag.StringVar(&ci.From, "from", "jpg", "変更元")
	flag.StringVar(&ci.From, "f", "jpg", "変更元(short)")
	flag.StringVar(&ci.To, "to", "png", "変更先")
	flag.StringVar(&ci.To, "t", "png", "変更先(short)")
}

func main() {
	flag.Parse()
	dirs := flag.Args()

	if !valid(ci.From) || !valid(ci.To) {
		fmt.Fprintln(os.Stderr, "supported formt is "+strings.Join(supportedFormat, ", "))
		os.Exit(1)
	}

	err := ci.Get(dirs)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	err = ci.Convert()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func valid(extension string) bool {
	for _, v := range supportedFormat {
		if v == extension {
			return true
		}
	}
	return false
}
