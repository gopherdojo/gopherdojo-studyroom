package splitDownload

import (
	"bytes"
	"context"
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
	StartByte int
	EndByte   int
	Body      []byte
}

type resource struct {
	Url              string
	Size             int
	BatchSize        int
	BatchCount       int
	Content          []byte
	PartialContentCh chan partialContent
	Http             http.Client
}

func (r *resource) GetSize() error {
	log.Printf("resource.getSize()\n")

	// HEADでサイズを調べる
	res, err := r.Http.Head(r.Url)
	if err != nil {
		return err
	}

	// データサイズを取得
	header := res.Header
	cl, ok := header["Content-Length"]
	if !ok {
		return errors.New("Content-Length couldn't be found")
	}
	r.Size, err = strconv.Atoi(cl[0])
	if err != nil {
		return err
	}
	return nil

}

func (r *resource) GetPartialContent(startByte int, endByte int, ctx context.Context) error {
	log.Printf("resource.getPartialContent(%d, %d)\n", startByte, endByte)
	// Rangeヘッダーを作成
	rangeVal := fmt.Sprintf("bytes=%d-%d", startByte, endByte)

	// リクエストとクライアントの作成
	reader := bytes.NewReader([]byte{})
	req, err := http.NewRequest("GET", r.Url, reader)
	if err != nil {
		return err
	}
	req.Header.Set("Range", rangeVal)
	client := &http.Client{}

	res := &http.Response{}
	const retryCount = 3
	for i := 0; i < retryCount; i++ {

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
	log.Println("start reading")
	bodyCh := make(chan []byte)
	go func() {
		body, _ := ioutil.ReadAll(res.Body)
		bodyCh <- body
	}()

	// 中止になったらBodyの読み込みを中止
	var body []byte
	select {
	case <-ctx.Done():
		log.Println("canceled reading body")
		return ctx.Err()
	case body = <-bodyCh:
		log.Println("finished reading body")
	}

	defer res.Body.Close()
	if err != nil {
		return err
	}

	pc := partialContent{StartByte: startByte, EndByte: endByte, Body: body}
	r.PartialContentCh <- pc
	return nil
}

func (r *resource) GetContent(batchCount int, ctx context.Context) error {
	log.Println("resource.getContent()")

	// コンテンツのデータサイズを取得
	err := r.GetSize()
	if err != nil {
		return err
	}
	log.Printf("r.size: %d\n", r.Size)

	// batchCount分リクエスト
	r.BatchCount = batchCount
	r.BatchSize = int(math.Ceil(float64(r.Size) / float64(r.BatchCount)))
	r.Content = make([]byte, r.Size)
	var eg *errgroup.Group
	eg, ctx = errgroup.WithContext(ctx)
	r.PartialContentCh = make(chan partialContent, r.BatchCount)
	for i := 0; i < r.BatchCount; i++ {

		// 担当する範囲を決定
		startByte := r.BatchSize * i
		endByte := r.BatchSize*(i+1) - 1
		if endByte > r.Size {
			endByte = r.Size
		}

		// TODO Channelを返すようにして、中断時に終了できるようにする
		// レンジごとにリクエスト
		eg.Go(func() error {
			return r.GetPartialContent(startByte, endByte, ctx)
		})
	}

	// １リクエストでも失敗すれば終了
	if err := eg.Wait(); err != nil {
		return err
	}

	// 一つのバイト列にマージ
	r.Content = make([]byte, r.Size)
	for i := 0; i < r.BatchCount; i++ {
		log.Println("merging...")
		pc := <-r.PartialContentCh
		log.Printf("pc.startByte: %v\n", pc.StartByte)
		fillByteArr(r.Content[:], pc.StartByte, pc.Body)
	}

	return nil
}
