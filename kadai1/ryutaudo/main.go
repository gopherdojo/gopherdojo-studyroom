package main

import (
	"flag"
	"fmt"
	"ryutaudo/imgcvt"
)

var dir, from, to string

func init() {
	flag.StringVar(&dir, "", "", "directory to walk through")
	flag.StringVar(&from, "from", "jpg", "extension of file to convert")
	flag.StringVar(&to, "to", "png", "extension of file to convert to")
}

func main() {
	fmt.Println("converting...")
	params := imgcvt.ConvertParams{From: from, To: to, Dir: dir}
	err := imgcvt.Convert(params)
	fmt.Println(fmt.Errorf("error: %s", err))
}
