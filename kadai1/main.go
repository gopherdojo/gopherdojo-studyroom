package main

import (
	"fmt"
	"os"
)
func main() {
	arguments, err := ParseArguments()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if arguments.isHelp {
		Help()
		os.Exit(0)
	}


	if err := arguments.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
