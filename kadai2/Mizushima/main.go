package main

import (
	"flag"
	"log"

	picconvert "github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai2/Mizushima/picconvert"
)

var preFormat string
var afterFormat string

func init() {
	flag.StringVar(&preFormat, "pre", "jpeg", "the file format before conversion")
	flag.StringVar(&afterFormat, "post", "png", "the file format after conversion")
	flag.Parse()
}

// isSupportedFormat returns true if "format" is supported, othrewise returns false.
func isSupportedFormat(format string) bool {
	ext := []string{"jpg", "jpeg", "png", "gif"}

	for _, e := range ext {
		if e == format {
			return true
		}
	}
	return false
}

// validate occurs a error if there is something wrong with 
// the entered parameters.
func validate() {
	if (preFormat == afterFormat) || 
		(preFormat == "jpeg" && afterFormat == "jpg") ||
		(preFormat == "jpg" && afterFormat == "jpeg") {
		log.Fatal("the parameter of -pre is same as that of -post.")
	}

	if len(flag.Args()) == 0 {
		log.Fatal("please input the path.")
	}

	if !isSupportedFormat(preFormat) {
		log.Fatal(preFormat, " is not supported.")
	} else if !isSupportedFormat(afterFormat) {
		log.Fatal(afterFormat, " is not supported.")
	}
}

func main() {
	// fmt.Println("converting" ,preFormat, "to", afterFormat)
	validate()
	c := picconvert.NewPicConverter(flag.Args()[0], preFormat, afterFormat)
	err := c.Conv()

	if err != nil {
		log.Fatal(err)
	}
}
