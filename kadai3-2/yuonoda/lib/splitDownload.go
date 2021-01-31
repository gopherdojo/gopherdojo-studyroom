package splitDownload

import (
	"log"
	"os"
	"path/filepath"
)

func Run(url string, batchCount int, dwDirPath string) string {
	log.Println("Run")

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
	err = r.GetContent(batchCount)
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
