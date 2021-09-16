package main

import (
	"flag"
	"fmt"
	"log"
)

var (
	inputType  string
	outputType string
)

func init() {
	flag.StringVar(&inputType, "input-type", "jpg", "input type")
	flag.StringVar(&outputType, "output-type", "png", "output type")
}

func validateArgs() error {
	flag.Parse()
	if flag.Arg(0) == "" {
		return fmt.Errorf("directory is required")
	}
	return nil
}

func main() {
	if err := validateArgs(); err != nil {
		log.Fatalln(err)
	}
}
