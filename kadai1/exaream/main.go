package main

import (
	"flag"
	"fmt"
	"os"

	ic "kadai1/imgconv"
)

func main() {
	flag.Parse()
	if err := ic.ValidArgs(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	conv := ic.NewConverter(*ic.SrcExt, *ic.DstExt, *ic.SrcDir, *ic.DstDir, *ic.FileDeleteFlag)
	err := conv.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
