package splitDownload

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// コンテンツのデータサイズを取得
func getContentSize(url string) (size int, err error) {
	// HEADでサイズを調べる
	res, err := http.Head(url)
	if err != nil {
		return 0, err
	}

	// データサイズを取得
	header := res.Header
	cl, ok := header["Content-Length"]
	if !ok {
		return 0, errors.New("Content-Length couldn't be found")
	}
	size, err = strconv.Atoi(cl[0])
	return
}

func Run() {
	log.Println("Run")

	// コンテンツのデータサイズを取得
	url := "https://dumps.wikimedia.org/jawiki/20210101/jawiki-20210101-pages-articles-multistream-index.txt.bz2"
	size, err := getContentSize(url)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("size: %d", size)

	// リクエストとクライアントの作成
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

	// bodyの取得
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
