package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai4/Mizushima/fortune"
	"golang.org/x/net/nettest"
)

func TestRouternOmikujiHandler(t *testing.T) {
	resp, close := serverTestHelper(t, fortune.OmikujiHandler, "TestServerOmikujiHandler")
	defer close()

	// 独自テスト
	var res fortune.Res
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&res); err != nil && err != io.EOF {
		t.Fatal(err)
	}

	s := []string{"大吉", "中吉", "小吉", "凶"}
	if res.Result == "" || !contains(s, res.Result) {
		t.Fatal("Error: json that the handler returned is invalid")
	}
}

func TestRouterHelloWorldHandler(t *testing.T) {
	resp, close := serverTestHelper(t, func(w http.ResponseWriter, r *http.Request) {
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

	// create the new test server
	ts := httptest.NewServer(router(handler))

	// get the response.
	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Fatalf("%s: error: %s", funcName, err)
	}

	// check the status code in the response.
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

func contains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

func TestListen(t *testing.T) {

	doneCh := make(chan struct{})
	osExit = func(code int) { doneCh <- struct{}{} }

	testListener, err := nettest.NewLocalListener("tcp")
	if err != nil {
		t.Fatal(err)
	}
	_, cancel := listen(context.Background(), testListener)
	defer cancel()

	proc, err := os.FindProcess(os.Getpid())
	if err != nil {
		t.Fatal(err)
	}

	if err := proc.Signal(os.Interrupt); err != nil {
		t.Fatal(err)
	}

	<-doneCh
}
