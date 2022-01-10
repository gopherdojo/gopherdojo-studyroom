package main

import (
	"flag"
	"log"

	"github.com/Kaiton7/gopherdojo-sdutyroom/kadai1/imkaito/imgconv"
)

var fromExt = flag.String("from", "jpg", "The extension to convert from")
var toExt = flag.String("to", "png", "The extension to convert to")

func main() {
	flag.Parse()
	directory := flag.Arg(0)
	converter, err := imgconv.New(directory, *fromExt, *toExt)
	if err != nil {
		log.Fatal(err)
	}
	err = converter.Walk()
	if err != nil {
		log.Fatal(err)
	}

}
