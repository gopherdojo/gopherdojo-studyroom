package main

import (
	"fmt"
	"os"
)

const (
	ExitCodeOK    = 0
	ExitCodeError = 1
)

func main() {
	cli := &CLI{inStream: os.Stdin, outStream: os.Stdout, errStream: os.Stderr}
	err := cli.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(ExitCodeError)
	}
	os.Exit(ExitCodeOK)
}
