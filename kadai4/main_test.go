package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOmikujiHandler(t *testing.T) {
	testserver := httptest.NewServer(http.HandlerFunc(omikujiHandler))
	defer testserver.Close()
	res, err := http.Get(testserver.URL)
	if err != nil {
		t.Error(err)
	}
	hello, err := ioutil.ReadAll(res.Body)
	fmt.Println("ページの内容：", string(hello))
	defer res.Body.Close()
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Error("a response code is not 200")
	}
}
