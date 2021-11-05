package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/saza-ku/gopherdojo-studyroom/kadai1/saza/imgconv"
)

var (
	root string
	src string
	dest string

	types = []string{"jpg", "png", "gif"}
)


func main() {
	s := flag.String("s", "jpg", "source image type")
	d := flag.String("d", "png", "destination image type")
	flag.Parse()

	root = flag.Args()[0]
	src = *s
	dest = *d

	if valid, message :=validateInput(); !valid {
		fmt.Println("Error: " + message)
		return
	}

	c := imgconv.NewConverter(root, src, dest)

	c.Run()
}

func validateInput() (bool, string) {
	// check whether root exists
	if f, err := os.Stat(root); err != nil || !f.IsDir() {
		return false, fmt.Sprintf("directory %s doesn't exists", root)
	}

	// check image types
	valid := false
	for _, t := range types {
		if src == t {
			valid = true
		}
	}
	if !valid {
		return false, fmt.Sprintf("image type %s is invalid", src)
	}

	valid = false
	for _, t := range types {
		if dest == t {
			valid = true
		}
	}
	if !valid {
		return false, fmt.Sprintf("image type %s is invalid", src)
	}

	return true, ""
}
