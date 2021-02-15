package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const (
	ExitCodeOK             = 0
	ExitCodeError          = 1
	ExitCodeParseFlagError = 1
)

type CLI struct {
	outStream, errStream io.Writer
}

func (c *CLI) Run(args []string) int {
	var dir string
	var from string
	var target []string

	flags := flag.NewFlagSet("imgconv", flag.ContinueOnError)
	flags.SetOutput(c.errStream)
	flags.StringVar(&dir, "dir", "", "")
	flags.StringVar(&from, "from", "jpg", "")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}

	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == "."+from {
				target = append(target, path)
			}
			return nil
		})
	if err != nil {
		return ExitCodeError
	}

	fmt.Println(target)

	return ExitCodeOK
}
