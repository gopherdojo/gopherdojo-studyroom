package interrupt

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Listen function check canceled or not
func Listen(ctx context.Context) (context.Context, func()) {
	ctx, cancel := context.WithCancel(ctx)
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		select {
		case <-ch:
			fmt.Println("interrupt")
			if f, err := os.Stat("tempdir"); os.IsExist(err) || f.IsDir() {
				if err := os.RemoveAll("tempdir"); err != nil {
					log.Fatal("err:", err)
				}
			}
			cancel()
		}
	}()

	return ctx, cancel
}
