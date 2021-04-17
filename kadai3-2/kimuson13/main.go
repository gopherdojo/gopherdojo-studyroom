package main

import (
	"context"
	"fmt"
	"os"

	"github.com/kimuson13/gopherdojo-studyroom/kimuson13/download"
	"github.com/kimuson13/gopherdojo-studyroom/kimuson13/interrupt"
)

func main() {
	err := setup(os.Args[1:])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func setup(args []string) error {
	var options download.Options
	ctx := context.Background()
	ctx, cancel := interrupt.Listen(ctx)
	defer cancel()

	opts, err := options.Parse(args)
	if err != nil {
		return err
	}

	return download.New(opts).Run(ctx)
}
