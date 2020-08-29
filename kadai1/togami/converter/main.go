package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	ExitCodeOK  = 0
	ExitCodeErr = 1
)

var (
	from *string = flag.String("f", ".jpg", "Enter the image ext you want convert")
	to   *string = flag.String("t", ".png", "Enter the dest for ext you want to convert")
)

func init() {
	flag.Parse()
}

func main() {
	if *from != ".jpg" && *from != ".jpeg" && *from != ".png" && *from != ".gif" {
		fmt.Println("Unsupported extension")
		os.Exit(ExitCodeErr)
	}
	dir := flag.Arg(0)
	paths, err := dirwalk(dir, from)
	if err != nil {
		os.Exit(ExitCodeErr)
	}
	for _, path := range paths {
		err := readFile(path, from, to)
		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(ExitCodeErr)
		}
	}
	handleFile(paths)
}
