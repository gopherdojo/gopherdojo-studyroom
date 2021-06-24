package listen

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Listen returns a context for keyboad(ctrl + c) interrupt.
func Listen(ctx context.Context, w io.Writer, f func()) (context.Context, func()) {
	ctx, cancel := context.WithCancel(ctx)

	ch := make(chan os.Signal, 2)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch
		_, err := fmt.Fprintln(w, "\n^Csignal : interrupt.")
		if err != nil {
			cancel()
			log.Fatalf("err: listen.Listen: %s\n", err)
		}
		cancel()
		f()
		os.Exit(0)
	}()

	return ctx, cancel
}
