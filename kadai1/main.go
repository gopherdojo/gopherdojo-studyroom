package main

import (
	"flag"
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
)

func jpg2png(filename string) {
	fs, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	img, err := jpeg.Decode(fs)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fs.Close()

	fs, err = os.OpenFile(filename+".png", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	err = png.Encode(fs, img)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	println(filename)
	fs.Close()
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Fprintln(os.Stderr, "error: need argument")
		return // should return 1
	}

	jpg2png(flag.Args()[0])
}
