package main

import (
	"flag"
	"fmt"
	"io"
	"strings"

	"github.com/rnakamine/gopherdojo-studyroom/kadai1/imgconv"
)

const (
	ExitCodeOK    = 0
	ExitCodeError = 1
)

var supportFormat = [...]string{"jpg", "jpeg", "png", "gif"}

type CLI struct {
	outStream, errStream io.Writer
}

// Run invokes the CLI with the given arguments.
func (c *CLI) Run(args []string) int {
	var dir, from, to string

	flags := flag.NewFlagSet("imgconv", flag.ContinueOnError)
	flags.SetOutput(c.errStream)

	flags.StringVar(&dir, "dir", "", "Specify the directory to be converted")
	flags.StringVar(&from, "from", "jpg", "Extension before conversion")
	flags.StringVar(&to, "to", "png", "Extensions after conversion")

	if err := flags.Parse(args[1:]); err != nil {
		fmt.Fprintln(c.errStream, err)
		return ExitCodeError
	}

	if dir == "" {
		fmt.Fprintln(c.errStream, "Directory is not specified.")
		return ExitCodeError
	}

	if !checkFormat(from) || !checkFormat(to) {
		fmt.Fprintln(c.errStream, "Unsupported format. Supported formats are jpg, jpeg, png and gif.")
		return ExitCodeError
	}

	images, err := imgconv.GetConvertImages(dir, from, to)
	if err != nil {
		fmt.Fprintln(c.errStream, err)
		return ExitCodeError
	}
	for _, img := range images {
		fmt.Fprintf(c.outStream, "Converting.. %s -> %s\n", img.FromPath, img.ToPath)
		if err := img.Convert(); err != nil {
			fmt.Fprintln(c.errStream, err)
			return ExitCodeError
		}
	}

	return ExitCodeOK
}

// CheckFormat is determine if the correct image is in the correct format.
func checkFormat(ext string) bool {
	for _, f := range supportFormat {
		if strings.ToLower(ext) == f {
			return true
		}
	}
	return false
}
