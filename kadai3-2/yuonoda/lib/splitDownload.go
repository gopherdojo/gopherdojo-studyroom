package splitDownload

import (
	"context"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func Run(url string, batchCount int, dwDirPath string) string {
	log.Println("Run")
	// TODO log.Fatalをやめ、正常系でも異常系でも最後に一時ファイルを削除する

	//　キャンセルコンテクストを定義
	ctx := context.Background()
	cancelCtx, cancel := context.WithCancel(ctx)
	defer cancel() // 何もなければコンテクストを開放

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

	// ファイルの作成
	_, filename := filepath.Split(url)
	dwFilePath := dwDirPath + "/" + filename + ".download"
	finishedFilePath := dwDirPath + "/" + filename
	dwFile, err := os.Create(dwFilePath)
	if err != nil {
		os.Remove(dwFilePath)
		log.Fatal(err)
	}

	// ダウンロード実行
	r := &resource{Url: url}
	err = r.GetContent(batchCount, cancelCtx)
	if err != nil {
		log.Fatal(err)
	}

	// データの書き込み
	_, err = dwFile.Write(r.Content)
	if err != nil {
		os.Remove(dwFilePath)
		log.Fatal(err)
	}
	os.Rename(dwFilePath, finishedFilePath)
	log.Println("download succeeded!")
	return finishedFilePath
}
