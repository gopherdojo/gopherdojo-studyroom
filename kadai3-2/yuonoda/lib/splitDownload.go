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

type partialContent struct {
	startByte int
	endByte   int
	body      []byte
}

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

//
func fillByteArr(arr []byte, startAt int, partArr []byte) {
	for i := 0; i < len(partArr); i++ {
		globalIndex := i + startAt
		arr[globalIndex] = partArr[i]
	}
}

// 指定範囲のデータを取得する
func getPartialContent(url string, startByte int, endByte int, fileDataCh chan partialContent) {
	// Rangeヘッダーを作成
	rangeVal := fmt.Sprintf("bytes=%d-%d", startByte, endByte)
	log.Println("rangeVal:", rangeVal)

	// リクエストとクライアントの作成
	r := bytes.NewReader([]byte{})
	req, err := http.NewRequest("GET", url, r)
	if err != nil {
		log.Fatal(err)
		//return nil, err
	}
	req.Header.Set("Range", rangeVal)
	client := &http.Client{}

	// リクエストの実行
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		//return nil, err
	}
	if res.StatusCode < 200 || res.StatusCode > 299 {
		log.Fatal(errors.New("status code is not 2xx, got " + res.Status))
	}

	// bodyの取得
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
		//return nil, err
	}
	pc := partialContent{body: body, startByte: startByte, endByte: endByte}
	fileDataCh <- pc
	return
}

func Run(url string, singleSize int) {
	log.Println("Run")

	// コンテンツのデータサイズを取得
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
	//singleSize := 1000000
	//singleSize = size / 2
	count := int(math.Ceil(float64(size) / float64(singleSize)))
	log.Printf("count: %d\n", count)
	pcch := make(chan partialContent, count)
	for i := 0; i < count; i++ {
		// 担当する範囲を決定
		startByte := singleSize * i
		endByte := singleSize*(i+1) - 1
		if endByte > size {
			endByte = size
		}

		// レンジごとにリクエスト
		go getPartialContent(url, startByte, endByte, pcch)
	}

	// 一つのバイト列にマージ
	fileData := make([]byte, size)
	for i := 0; i < count; i++ {
		pc := <-pcch
		fillByteArr(fileData[:], pc.startByte, pc.body)
	}

	// データの書き込み
	_, err = file.Write(fileData)
	if err != nil {
		log.Fatal(err)
	}
}
