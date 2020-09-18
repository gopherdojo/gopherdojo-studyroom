package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/oamam/gopherdojo-studyroom/kadai1/oamam/imgconv"
)

func main() {
	id := flag.String("id", "", "input directory")
	od := flag.String("od", "", "output directory")
	ie := flag.String("ie", "jpg", "extension of input image")
	oe := flag.String("oe", "png", "extension of output image")
	flag.Parse()

	if err := imgconv.Do(id, od, ie, oe); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully converted!")
}
