package main

import (
	"fmt"
	"os"

	"github.com/kimuson13/gopherdojo-studyroom/kimuson13/download"
)

func main() {
	cli := download.New()
	if err := cli.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error:\n %v\n", err)
	}
	os.Exit(1)
}
