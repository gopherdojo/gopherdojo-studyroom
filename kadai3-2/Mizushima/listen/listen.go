package listen

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
)

func Listen(ctx context.Context, w io.Writer, f func()) (context.Context, func()) {
	ctx, cancel := context.WithCancel(ctx)

	ch := make(chan os.Signal, 2)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch
		fmt.Fprintln(w, "\n^Csignal pressed. interrupt.")
		cancel()
		f()
		os.Exit(0)
	}()

	return ctx, cancel
}