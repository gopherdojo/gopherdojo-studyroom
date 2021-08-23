package main

import (
	"fmt"
	"gopherdojo-studyroom/kadai1/mypkg"
	"os"
)
func main() {
	arguments, err := mypkg.ParseArguments()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if arguments.IsHelp {
		mypkg.Help()
		os.Exit(0)
	}


	if err := arguments.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
