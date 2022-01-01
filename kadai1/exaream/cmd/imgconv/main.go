package main

import (
	"flag"
	"fmt"
	"os"

	ic "github.com/exaream/gopherdojo-studyroom/kadai1/exaream/imgconv"
)

func main() {
	// Get arguments and validate them
	flag.Parse()
	if err := ic.ValidateArgs(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// Convert image files
	conv := ic.NewConverter(*ic.SrcExt, *ic.DstExt, *ic.SrcDir, *ic.DstDir, *ic.FileDeleteFlag)
	err := conv.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
