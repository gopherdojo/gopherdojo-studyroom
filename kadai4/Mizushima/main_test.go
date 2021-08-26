package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBootServer(t *testing.T) {
	ts:= httptest.NewServer(router(omikujiHandler))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := resp.Body.Close(); err!= nil {
			t.Fatal(err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code: %d", resp.StatusCode)
	}

	var res Res
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&res); err != nil && err != io.EOF {
		t.Fatal(err)
	}

	s := []string{"大吉", "中吉", "小吉", "凶"}
	if res.Result == "" || !contains(s, res.Result) {
		t.Fatal("Error: json that the handler returned is invalid")
	}
}