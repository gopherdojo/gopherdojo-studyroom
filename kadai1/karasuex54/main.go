package main

import (
	"flag"
	"fmt"

	"github.com/karasuex54/gopherdojo-studyroom/kadai1/karasuex54/converter"
)

var (
	targetDir string
	fromExt   string
	toExt     string
)

func main() {
	flag.StringVar(&fromExt, "from", "jpg", "target image file extention")
	flag.StringVar(&toExt, "to", "png", "convert image file extention")
	flag.Parse()

	args := flag.Args()
	switch len(args) {
	case 0:
		fmt.Println("few arguments")
		return
	case 1:
		targetDir = args[0]
	default:
		fmt.Println("too arguments")
		return
	}

	c := converter.NewConverter(targetDir, converter.FromExt(fromExt), converter.ToExt(toExt))
	c.Run()
}
