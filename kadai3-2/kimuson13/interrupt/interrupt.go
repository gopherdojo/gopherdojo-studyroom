package interrupt

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Listen(ctx context.Context) (context.Context, func()) {
	ctx, cancel := context.WithCancel(ctx)
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		select {
		case <-ch:
			fmt.Println("interrupt")
			if err := os.RemoveAll("tempdir"); err != nil {
				log.Fatal("Error:", err)
			}
			cancel()
		}
	}()

	return ctx, cancel
}
