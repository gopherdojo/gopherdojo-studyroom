package main

import (
	"flag"
	"fmt"
	"log"
	"testing"

	picconvert "github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai2/Mizushima/picconvert"
)

var preFormat string
var afterFormat string
var srcPath string

func init() {
	testing.Init()
	flag.StringVar(&preFormat, "pre", "jpeg", "the file format before conversion")
	flag.StringVar(&afterFormat, "post", "png", "the file format after conversion")
	flag.StringVar(&srcPath, "path", ".", "the source directory for conversion")
	flag.Parse()
}

// isSupportedFormat returns true if "format" is supported, othrewise returns false.
func IsSupportedFormat(format string) bool {
	ext := []string{"jpg", "jpeg", "png", "gif"}

	for _, e := range ext {
		if e == format {
			return true
		}
	}
	return false
}

// validate returns error if there is something wrong with
// the entered parameters.
func Validate() error {
	// fmt.Printf("%v, %v, %v\n", preFormat, afterFormat, srcPath)
	if preFormat == afterFormat {
		return fmt.Errorf("the parameter of -pre is same as that of -post")
	}

	if (preFormat == "jpeg" && afterFormat == "jpg") || (preFormat == "jpg" && afterFormat == "jpeg") {
		return fmt.Errorf("the parameter of -pre is same as that of -post. 'jpg' is considered same as 'jpeg'")
	}

	if !IsSupportedFormat(preFormat) {
		return fmt.Errorf("-pre %s is not supported", preFormat)
	} else if !IsSupportedFormat(afterFormat) {
		return fmt.Errorf("-post %s is not supported", afterFormat)
	}

	return nil
}

func main() {
	// fmt.Println("converting" ,preFormat, "to", afterFormat)
	err := Validate()
	if err != nil {
		log.Fatal(err)
	}

	c := picconvert.NewPicConverter(srcPath, preFormat, afterFormat)

	err = c.Conv()
	if err != nil {
		log.Fatal(err)
	}
}
