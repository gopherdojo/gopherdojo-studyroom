package main

import (
	"fmt"
	"os"
)

func main() {
	cli := &CLI{inStream: os.Stdin, outStream: os.Stdout, errStream: os.Stderr}
	ExitCode, err := cli.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(ExitCode)
}
