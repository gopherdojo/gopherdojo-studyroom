package termination

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
)

var cleanFns []func()
var osExit = os.Exit

// Listen listens signal
func Listen(ctx context.Context, w io.Writer) (context.Context, func()) {
	ctx, cancel := context.WithCancel(ctx)

	ch := make(chan os.Signal, 2)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch
		fmt.Fprintln(w, "Ctrl+C pressed in Terminal")
		cancel()
		for _, f := range cleanFns {
			f()
		}
		osExit(0)
	}()

	return ctx, cancel
}

// CleanFunc registers clean function
func CleanFunc(f func()) {
	cleanFns = append(cleanFns, f)
}
