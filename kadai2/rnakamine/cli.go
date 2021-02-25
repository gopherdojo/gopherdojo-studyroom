package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"strings"

	"github.com/rnakamine/gopherdojo-studyroom/kadai1/imgconv"
)

var supportFormat = [...]string{"jpg", "jpeg", "png", "gif"}

type CLI struct {
	inStream             io.Reader
	outStream, errStream io.Writer
}

// Run invokes the CLI with the given arguments.
func (c *CLI) Run(args []string) error {

	flags := flag.NewFlagSet("imgconv", flag.ContinueOnError)
	flags.SetOutput(c.errStream)

	var dir, from, to string
	flags.StringVar(&dir, "dir", "", "Specify the directory to be converted")
	flags.StringVar(&from, "from", "jpg", "Extension before conversion")
	flags.StringVar(&to, "to", "png", "Extensions after conversion")

	var del bool
	flags.BoolVar(&del, "del", false, "Delete the original image.")

	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	if dir == "" {
		return errors.New("Directory is not specified.")
	}

	if !checkFormat(from) || !checkFormat(to) {
		return errors.New("Unsupported format. Supported formats are jpg, jpeg, png and gif.")
	}

	var deleteOption bool
	if del {
		_, err := fmt.Fprintln(c.outStream, "Do you really want to delete the original image? (Y/N)")
		if err != nil {
			return err
		}
		in := bufio.NewScanner(c.inStream)
		in.Scan()
		answer := in.Text()
		if answer == "y" || answer == "Y" {
			deleteOption = true
		} else {
			return errors.New("It suspends processing.")
		}
	} else {
		deleteOption = false
	}

	images, err := imgconv.GetConvertImages(dir, from, to)
	if err != nil {
		return err
	}
	for _, img := range images {
		_, err := fmt.Fprintf(c.outStream, "Converting.. %s -> %s\n", img.FromPath, img.ToPath)
		if err != nil {
			return err
		}
		if err := img.Convert(deleteOption); err != nil {
			return err
		}
	}

	return nil
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
