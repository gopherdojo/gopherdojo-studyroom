package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	from *string = flag.String("f", ".jpg", "Enter the image ext you want convert")
	to   *string = flag.String("t", ".png", "Enter the dest for ext you want to convert")
)

func init(){
	flag.Parse()
}

func main() {
	dir := flag.Arg(0)
	paths, err := dirwalk(dir, from)
	if err != nil {
		os.Exit(1)
	}
	for _, path := range paths {
		err := readFile(path, from, to)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
	}
	handleFile(paths)
}
