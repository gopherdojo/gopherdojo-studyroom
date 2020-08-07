package main

import (
	"fmt"
	"os"

	// This package has not merged yet
	"github.com/task4233/gopherdojo-studyroom/kadai1/task4233/eimg"
)

func main() {
	cli := eimg.New()
	if err := cli.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error:\n%+v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
