package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kimuson13/gopherdojo-studyroom/kimuson13/conversion"
)

var before, after string

func init() {
	flag.StringVar(&before, "b", "jpg", "select extension of conversion source")
	flag.StringVar(&after, "a", "png", "select extension you want to convert")
}

func main() {
	flag.Parse()
	dirs := flag.Args()

	err := conversion.ExtensionCheck(before)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	for _, dir := range dirs {
		err := conversion.WalkDir(dir, after)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	}
}
