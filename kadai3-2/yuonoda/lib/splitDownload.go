package splitDownload

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func Run() {
	log.Println("Run")

	// リクエストとクライアントの作成
	url := "https://dumps.wikimedia.org/jawiki/20210101/jawiki-20210101-pages-articles-multistream-index.txt.bz2"
	r := bytes.NewReader([]byte{})
	req, err := http.NewRequest("GET", url, r)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}

	// リクエストの実行
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	// ファイルの作成
	_, filename := filepath.Split(url)
	file, err := os.Create(filename + ".download")
	if err != nil {
		log.Fatal(err)
	}

	// データの書き込み
	_, err = file.Write(body)
	if err != nil {
		log.Fatal(err)
	}
}
