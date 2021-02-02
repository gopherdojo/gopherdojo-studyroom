package terminate

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Listen() context.Context {
	//　キャンセルコンテクストを定義
	ctx := context.Background()
	cancelCtx, cancel := context.WithCancel(ctx)

	// 中断シグナルがきたらキャンセル処理
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		select {
		case <-c:
			log.Println("interrupted")
			cancel()
		}
	}()
	return cancelCtx

}
