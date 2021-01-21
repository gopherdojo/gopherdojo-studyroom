package splitDownload

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
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

func getPartialContent(url string, startByte int, endByte int) ([]byte, error) {
	// Rangeヘッダーを作成
	rangeVal := fmt.Sprintf("bytes=%d-%d", startByte, endByte)
	log.Println("rangeVal:", rangeVal)

	// リクエストとクライアントの作成
	r := bytes.NewReader([]byte{})
	req, err := http.NewRequest("GET", url, r)
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}
	req.Header.Set("Range", rangeVal)
	client := &http.Client{}

	// リクエストの実行
	res, err := client.Do(req)
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}

	// bodyの取得
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}
	return body, nil
}

func Run() {
	log.Println("Run")

	// コンテンツのデータサイズを取得
	url := "https://dumps.wikimedia.org/jawiki/20210101/jawiki-20210101-pages-articles-multistream-index.txt.bz2"
	size, err := getContentSize(url)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("size: %d\n", size)

	// ファイルの作成
	_, filename := filepath.Split(url)
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	filePath := homedir + "/Downloads/" + filename
	log.Println(filePath)
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// 1MBごとにアクセス
	singleSize := 1000000
	count := int(math.Ceil(float64(size) / float64(singleSize)))
	log.Printf("count: %d\n", count)
	var fileData []byte
	for i := 0; i < count; i++ {
		// 担当する範囲を決定
		startByte := singleSize * i
		endByte := singleSize*(i+1) - 1

		// レンジごとにリクエスト
		content, err := getPartialContent(url, startByte, endByte)
		if err != nil {
			log.Fatal(err)
		}
		fileData = append(fileData, content...)
	}

	// データの書き込み
	_, err = file.Write(fileData)
	if err != nil {
		log.Fatal(err)
	}
}
