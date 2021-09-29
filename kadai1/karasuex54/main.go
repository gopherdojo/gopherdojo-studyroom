package main

import (
	"flag"

	"github.com/karasuex54/gopherdojo-studyroom/converter"
)

var (
	targetDir string
	fromExt   string
	toExt     string
)

func main() {
	flag.Parse()
	flag.StringVar(&fromExt, "from", "jpg", "target image file extention")
	flag.StringVar(&toExt, "to", "png", "convert image file extention")

	args := flag.Args()
	switch len(args) {
	case 0:
		return
	case 1:
		targetDir = args[0]
	default:
		return
	}

	c := converter.NewConverter(targetDir, converter.FromExt(fromExt), converter.ToExt(toExt))
	c.Run()
}
