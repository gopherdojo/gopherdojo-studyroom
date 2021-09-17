package main

import (
	"flag"
	"fmt"
	"imgconv/imgconv"
	"log"
)

var (
	inputType  string
	outputType string
	directory  string
)

func init() {
	flag.StringVar(&inputType, "input-type", "jpg", "input type")
	flag.StringVar(&outputType, "output-type", "png", "output type")
}

func validateType(t string) error {
	switch t {
	case "jpg":
		fallthrough
	case "png":
		return nil
	default:
		return fmt.Errorf("invalid type: %s", t)
	}
}

func validateArgs() error {
	flag.Parse()

	if inputType == outputType {
		return fmt.Errorf("input and output type can't be the same")
	}

	if err := validateType(inputType); err != nil {
		return err
	}

	if err := validateType(outputType); err != nil {
		return err
	}

	if flag.Arg(0) == "" {
		return fmt.Errorf("directory is required")
	} else {
		directory = flag.Arg(0)
	}

	return nil
}

func main() {
	if err := validateArgs(); err != nil {
		log.Fatalln(err)
	}
}
