package main

import (
	picconvert "Mizushima/pic-conv"
	"flag"
	"log"
)

var preFormat string
var afterFormat string

func init() {
	flag.StringVar(&preFormat, "pre", "jpeg", "the file format before conversion")
	flag.StringVar(&afterFormat, "post", "png", "the file format after conversion")
	flag.Parse()
}

func isCorresponds(format string) bool {
	ext := []string{"jpg", "jpeg", "png", "gif"}

	var flag bool = false

	for _, e := range ext {
		if e == format {
			flag = true
		}
	}

	if !flag {
		return false
	}
	return true
}

// errHandler occur errors.
func errHandler() {
	if preFormat == afterFormat {
		log.Fatal("arguments of -pre and -post are same.")
	}

	if len(flag.Args()) == 0 {
		log.Fatal("please input the path.")
	}

	if !isCorresponds(preFormat) {
		log.Fatal(preFormat, " is not supported.")
	} else if !isCorresponds(afterFormat) {
		log.Fatal(afterFormat, " is not supported.")
	}
}

func main() {
	// fmt.Println("converting" ,preFormat, "to", afterFormat)
	errHandler()
	c := picconvert.NewPicConverter(flag.Args()[0], preFormat, afterFormat)
	c.Conv()
}
