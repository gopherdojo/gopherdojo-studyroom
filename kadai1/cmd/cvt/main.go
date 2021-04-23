package main

import (
	"flag"
	"fmt"
	"github.com/gopherdojo/gopherdojo-studyroom/kadai1/internal/cvt"
	"log"
)

var inputDir, outputDir, beforeExt, afterExt string
var removeSrc bool

func init() {
	flag.StringVar(&inputDir, "i", "", "input dir")
	flag.StringVar(&outputDir, "o", "", "output dir")
	flag.StringVar(&beforeExt, "be", ".jpg", "before ext")
	flag.StringVar(&afterExt, "ae", ".png", "after ext")
	flag.BoolVar(&removeSrc, "rm", false, "remove src")
	flag.Parse()
}

func main() {
	c := cvt.NewImageCvtGlue(inputDir, outputDir, beforeExt, afterExt, removeSrc)
	if err := c.Exec(); err != nil {
		log.Fatalf("Failed to execute image convert", fmt.Sprintf("%+v", err))
	}
}
