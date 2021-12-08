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
	flag.Parse()
}

func main() {
	if err := validateInput(); err != nil {
		fmt.Println(err)
		return
	}

	root = flag.Args()[0]
	c := imgconv.NewConverter(root, src, dest)

	c.Run()
}

func validateInput() error {
	// check argument
	if len(flag.Args()) == 0 {
		return fmt.Errorf("target directory is not provided")
	}
	if len(flag.Args()) > 1 {
		return fmt.Errorf("too many arguments")
	}

	// check whether root exists
	r := flag.Args()[0]
	if f, err := os.Stat(r); err != nil || !f.IsDir() {
		return fmt.Errorf("directory %s doesn't exists", root)
	}

	// check image types
	valid := false
	for _, t := range types {
		if src == t {
			valid = true
		}
	}
	if !valid {
		return fmt.Errorf("image type %s is invalid", src)
	}

	valid = false
	for _, t := range types {
		if dest == t {
			valid = true
		}
	}
	if !valid {
		return fmt.Errorf("image type %s is invalid", src)
	}

	return nil
}
