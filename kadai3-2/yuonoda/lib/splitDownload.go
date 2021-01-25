package splitDownload

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Run(url string, batchCount int) {
	log.Println("Run")

	// ファイルの作成
	_, filename := filepath.Split(url)
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	dwFilePath := homedir + "/Downloads/" + filename + ".download"
	log.Println(dwFilePath)
	dwFile, err := os.Create(dwFilePath)
	if err != nil {
		os.Remove(dwFilePath)
		log.Fatal(err)
	}

	// ダウンロード実行
	r := &resource{url: url}
	err = r.getContent(batchCount)
	if err != nil {
		log.Fatal(err)
	}

	// データの書き込み
	_, err = dwFile.Write(r.content)
	if err != nil {
		os.Remove(dwFilePath)
		log.Fatal(err)
	}
	os.Rename(dwFilePath, strings.Trim(dwFilePath, ".download"))

	log.Println("download succeeded!")
}
