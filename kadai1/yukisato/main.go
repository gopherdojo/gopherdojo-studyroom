package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yukisato/gopherdojo-studyroom/kadai1/yukisato/conv"
)

func main() {
	extFrom := flag.String("from", "jpg", "from ext")
	extTo := flag.String("to", "png", "to ext")
	flag.Parse()
	args := flag.Args()
	destDir := "."

	if len(args) > 1 {
		fmt.Fprintln(os.Stderr, "Too many arguments. Only one destination directory is needed")
		os.Exit(1)
	}

	if len(args) == 1 {
		destDir = args[0]
	}

	// For checking the extensions precisely, add periods.
	conv.ConvertImages(destDir, "."+*extFrom, "."+*extTo)
}
