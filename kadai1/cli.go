package main

import (
	"fmt"
	"io"
)

const ExitCodeOK = 1

type CLI struct {
	outStream, errStream io.Writer
}

func (c *CLI) Run(args []string) int {
	// flags := flag.NewFlagSet("imgconv", flag.ContinueOnError)

	fmt.Fprint(c.outStream, "Do awesome work\n")

	return ExitCodeOK
}
