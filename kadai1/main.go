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
		help()
		os.Exit(0)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	c := newConvert(arguments.args)
	fmt.Println(c.args)

	if err := c.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
