package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"youenn/conversion"
	"youenn/fileutil"
)

var src string
var dst string
var targetPath string

func init() {
	flag.StringVar(&src, "s", "jpeg", "file type before conversion")
	flag.StringVar(&dst, "d", "png", "file type after conversion")
	flag.Parse()
	targetPath = flag.Args()[0]
}

func main() {

	//check command line arguments
	if src == dst {
		log.Fatal("Source file type and destination file type is same.")
	}
	if !conversion.IsSupported(src) {
		log.Fatal("Source file type is not supported")
	}
	if !conversion.IsSupported(dst) {
		log.Fatal("Destination file type is not supported")
	}
	if !fileutil.Exists(targetPath) {
		log.Fatalf("Directory %s does not exit", targetPath)
	}
	if !fileutil.IsDir(targetPath) {
		log.Fatal(targetPath + " is not directory")
	}

	cnt := conversion.WalkConvert(targetPath, src, dst)
	fmt.Println(strconv.Itoa(cnt) + " files have been converted.")
	os.Exit(0)

}
