package main

import (
	"context"
	"fmt"
	"os"

	"github.com/kimuson13/gopherdojo-studyroom/kimuson13/download"
	"github.com/kimuson13/gopherdojo-studyroom/kimuson13/interrupt"
)

func main() {
	err := setUp(os.Args[1:])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

//setUp function preparate for running in main package
func setUp(args []string) error {
	var options download.Options
	ctx := context.Background()
	ctx, cancel := interrupt.Listen(ctx)
	defer cancel()

	opts, err := options.Parse(args...)
	if err != nil {
		return err
	}

	return download.New(opts).Run(ctx)
}
