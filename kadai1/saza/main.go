package main

import (
	"flag"

	"github.com/saza-ku/gopherdojo-studyroom/kadai1/saza/imgconv"
)

var (
	root string
	src string
	dest string
)

func main() {
	flag.Parse()
	root = flag.Args()[0]
	src = *flag.String("s", "jpeg", "source image type")
	dest = *flag.String("d", "png", "destination image type")

	c := imgconv.NewConverter(root, src, dest)

	c.Run()
}
