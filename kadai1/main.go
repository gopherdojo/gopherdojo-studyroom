package main

import (
	"fmt"
	"gopherdojo-studyroom/kadai1/pkg"
	"os"
)
func main() {
	arguments, err := pkg.ParseArguments()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if arguments.IsHelp {
		pkg.Help()
		os.Exit(0)
	}


	if err := arguments.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
