package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServerOmikujiHandler(t *testing.T) {
	resp, close := serverTestHelper(t, omikujiHandler, "TestServerOmikujiHandler")
	defer close()

	// 独自テスト
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

func TestServerHelloWorldHandler(t *testing.T) {
	resp, close := serverTestHelper(t, func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "Hello, World!")
	}, "TestServerHelloWorldHandler")
	defer close()
	
	// 独自テスト
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("TestServerHelloWorldHandler: error: %s", err)
	}

	const expected = "Hello, World!\n"
	if string(b) != expected {
		t.Fatalf("expected %s, but got %s", expected, string(b))
	}
}

// serverTestHelper performs a connection test on a test server with 'handler'.
func serverTestHelper(t *testing.T,
	handler func(w http.ResponseWriter, r *http.Request),
	funcName string) (*http.Response, func()) {

	t.Helper()

	ts := httptest.NewServer(router(handler))
	
	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Fatalf("%s: error: %s", funcName, err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("%s: error: unexpected status code: %d", funcName, resp.StatusCode)
	}

	return resp,
	func() {
		ts.Close()
		if err := resp.Body.Close(); err != nil {
			t.Fatalf("resp.Body.Close(): error: %s", err)
		}
	}
}

