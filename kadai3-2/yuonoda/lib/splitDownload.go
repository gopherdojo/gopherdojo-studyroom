package splitDownload

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func Run() {
	log.Println("Run")

	// リクエストとクライアントの作成
	url := "https://yahoo.co.jp"
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
	log.Printf("body:%s", body)
}
