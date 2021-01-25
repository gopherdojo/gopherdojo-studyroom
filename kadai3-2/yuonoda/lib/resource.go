package splitDownload

import (
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type partialContent struct {
	startByte int
	endByte   int
	body      []byte
}

type resource struct {
	url              string
	size             int
	batchSize        int
	batchCount       int
	content          []byte
	partialContentCh chan partialContent
}

func (r *resource) getSize() error {
	log.Printf("resource.getSize()\n")

	// HEADでサイズを調べる
	res, err := http.Head(r.url)
	if err != nil {
		return err
	}

	// データサイズを取得
	header := res.Header
	cl, ok := header["Content-Length"]
	if !ok {
		return errors.New("Content-Length couldn't be found")
	}
	r.size, err = strconv.Atoi(cl[0])
	if err != nil {
		return err
	}
	return nil

}

func (r *resource) getPartialContent(startByte int, endByte int) error {
	log.Printf("resource.getPartialContent(%d, %d)\n", startByte, endByte)
	// Rangeヘッダーを作成
	rangeVal := fmt.Sprintf("bytes=%d-%d", startByte, endByte)

	// リクエストとクライアントの作成
	reader := bytes.NewReader([]byte{})
	req, err := http.NewRequest("GET", r.url, reader)
	if err != nil {
		return err
	}
	req.Header.Set("Range", rangeVal)
	client := &http.Client{}

	res := &http.Response{}
	for i := 0; i < 3; i++ {
		// リクエストの実行
		log.Printf("rangeVal[%d]:%s", i, rangeVal)
		res, err = client.Do(req)
		if err != nil {
			return err
		}

		// ステータスが200系ならループを抜ける
		log.Printf("res.StatusCode:%d\n", res.StatusCode)
		if res.StatusCode >= 200 && res.StatusCode <= 299 {
			break
		}

		// 乱数分スリープ
		rand.Seed(time.Now().UnixNano())
		randFloat := rand.Float64() + 1
		randMs := math.Pow(randFloat, float64(i+1)) * 1000
		sleepTime := time.Duration(randMs) * time.Millisecond
		log.Printf("sleep:%v\n", sleepTime)
		time.Sleep(sleepTime)
	}

	// 正常系レスポンスでないとき
	if res.StatusCode < 200 || res.StatusCode > 299 {
		return errors.New("status code is not 2xx, got " + res.Status)
	}

	// bodyの取得
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return err
	}

	pc := partialContent{startByte: startByte, endByte: endByte, body: body}
	r.partialContentCh <- pc
	return nil
}

func (r *resource) getContent(batchCount int) error {
	log.Println("resource.getContent()")

	// コンテンツのデータサイズを取得
	err := r.getSize()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("r.size: %d\n", r.size)

	// batchCount分リクエスト
	r.batchCount = batchCount
	r.batchSize = int(math.Ceil(float64(r.size) / float64(r.batchCount)))
	r.content = make([]byte, r.size)
	var eg errgroup.Group
	r.partialContentCh = make(chan partialContent, r.batchCount)
	for i := 0; i < r.batchCount; i++ {

		// 担当する範囲を決定
		startByte := r.batchSize * i
		endByte := r.batchSize*(i+1) - 1
		if endByte > r.size {
			endByte = r.size
		}

		// レンジごとにリクエスト
		eg.Go(func() error {
			return r.getPartialContent(startByte, endByte)
		})
	}

	// １リクエストでも失敗すれば終了
	if err := eg.Wait(); err != nil {
		return err
	}

	// 一つのバイト列にマージ
	r.content = make([]byte, r.size)
	for i := 0; i < r.batchCount; i++ {
		log.Println("merging...")
		pc := <-r.partialContentCh
		log.Printf("pc.startByte: %v\n", pc.startByte)
		fillByteArr(r.content[:], pc.startByte, pc.body)
	}

	return nil
}
