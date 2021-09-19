package main

import (
	"flag"
	"fmt"
	"imgconv/imgconv"
	"log"
)

type Options struct {
	inputType  string
	outputType string
	directory  string
}

var opt Options

func init() {
	flag.StringVar(&opt.inputType, "input-type", "jpg", "input type[jpg|jpeg|png|gif]")
	flag.StringVar(&opt.outputType, "output-type", "png", "output type[jpg|jpeg|png|gif]")
}

// validateType validates the type of the image
func validateType(t string) error {
	switch t {
	case "jpg", "jpeg", "png", "gif":
		return nil
	default:
		return fmt.Errorf("invalid type: %s", t)
	}
}

// validateArgs validates the arguments
func validateArgs() error {
	flag.Parse()

	if opt.inputType == opt.outputType {
		return fmt.Errorf("input and output type can't be the same")
	}

	if err := validateType(opt.inputType); err != nil {
		return err
	}

	if err := validateType(opt.outputType); err != nil {
		return err
	}

	if flag.Arg(0) == "" {
		return fmt.Errorf("directory is required")
	} else {
		opt.directory = flag.Arg(0)
	}

	return nil
}

func main() {
	if err := validateArgs(); err != nil {
		log.Fatalln(err)
	}

	if err := imgconv.Converter(opt.directory, opt.inputType, opt.outputType); err != nil {
		log.Fatalln(err)
	}
}
