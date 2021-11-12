package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/saza-ku/gopherdojo-studyroom/kadai1/saza/imgconv"
)

var (
	root string
	src  string
	dest string

	types = []string{"jpg", "png", "gif"}
)

func init() {
	flag.StringVar(&src, "s", "jpg", "source image type")
	flag.StringVar(&dest, "d", "png", "destination image type")
}

func main() {
	flag.Parse()
	args := flag.Args()

	// TODO: validateInput と一緒にしたい
	if len(args) != 1 {
		fmt.Println("root directory is not provided")
		return
	}

	root = args[0]

	if valid, message := validateInput(); !valid {
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
