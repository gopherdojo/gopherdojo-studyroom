package splitDownload

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Run(url string, batchCount int, dwDirPath string) string {
	log.Println("Run")

	// ファイルの作成
	_, filename := filepath.Split(url)
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	if dwDirPath == "" {
		dwDirPath = homedir + "/Downloads"
	}
	dwFilePath := dwDirPath + "/" + filename + ".download"
	log.Println(dwFilePath)
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
	finishedFilePath := strings.Trim(dwFilePath, ".download")
	os.Rename(dwFilePath, finishedFilePath)

	log.Println("download succeeded!")
	return finishedFilePath
}
