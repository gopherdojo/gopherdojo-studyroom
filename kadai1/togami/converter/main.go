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

func main() {
	flag.Parse()
	dir := flag.Arg(0)
	paths := dirwalk(dir, from)
	fmt.Println(paths)
	for _, path := range paths {
		err := readFile(path, from, to)
		if err != nil {
			fmt.Fprintln(os.Stderr, "%s\n", err)
		}
	}

	answer := handleFile()
	switch answer {
	case "y":
		for _, path := range paths {
			deleteFile(path)
		}
	case "n":
		return
	default:
		fmt.Println("Please enter again")
		handleFile()
	}
}
