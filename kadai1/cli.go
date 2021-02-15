package main

import (
	"flag"
	"image"
	_ "image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
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

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == "."+from {
			err = convertImage(path, to)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return ExitCodeError
	}

	return ExitCodeOK
}

func convertImage(path string, ext string) error {
	sf, err := os.Open(path)
	if err != nil {
		return err
	}
	defer sf.Close()

	img, _, err := image.Decode(sf)
	if err != nil {
		return err
	}

	df, _ := os.Create(path[:len(path)-len(filepath.Ext(path))] + "." + ext)
	defer df.Close()

	err = png.Encode(df, img)
	if err != nil {
		return err
	}

	err = os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}
