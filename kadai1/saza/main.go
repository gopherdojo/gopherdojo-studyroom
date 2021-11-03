package main

import (
	"os"

	"github.com/saza-ku/gopherdojo-studyroom/kadai1/saza/imgconv"
)

func main() {
	root := os.Args[1]
	c := imgconv.Converter{
		Root: root,
	}
	c.Run()
}
