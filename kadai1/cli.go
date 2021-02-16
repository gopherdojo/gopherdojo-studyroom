package main

import (
	"flag"
	"fmt"
	"io"

	"./imgconv"
)

const (
	ExitCodeOK    = 0
	ExitCodeError = 1
)

type CLI struct {
	outStream, errStream io.Writer
}

func (c *CLI) Run(args []string) int {
	var (
		dir  string
		from string
		to   string
	)

	flags := flag.NewFlagSet("imgconv", flag.ContinueOnError)
	flags.SetOutput(c.errStream)

	flags.StringVar(&dir, "dir", "", "")
	flags.StringVar(&from, "from", "jpg", "")
	flags.StringVar(&to, "to", "png", "")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	images, err := imgconv.FileWalk(dir, from, to)
	if err != nil {
		fmt.Fprintln(c.errStream, err)
		return ExitCodeError
	}
	for _, img := range images {
		err := img.Convert()
		if err != nil {
			fmt.Fprintln(c.errStream, err)
			return ExitCodeError
		}
	}

	return ExitCodeOK
}
