package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/kimuson13/gopherdojo-studyroom/kimuson13/download"
)

func main() {
	flag.Parse()
	url := flag.Arg(0)
	ctx := context.Background()
	err := download.Run(url, ctx)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
