package main

import (
	"fmt"
	"os"

	"example.com/ex01/imgconv"
)

func main() {
	dirs, from, to, err := imgconv.ValidateArgs(os.Args)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
	err = imgconv.Convert(dirs, from, to)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}
