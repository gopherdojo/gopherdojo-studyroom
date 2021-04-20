package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

var ts *httptest.Server

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func setUp() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/header.jpg", http.StatusFound)
	})

	mux.HandleFunc("/header.jpg", func(w http.ResponseWriter, r *http.Request) {
		fp := "testdata/header.jpg"
		data, err := os.ReadFile(fp)
		if err != nil {
			panic(err)
		}
		http.ServeContent(w, r, fp, time.Now(), bytes.NewReader(data))
	})

	ts = httptest.NewServer(mux)
	// defer ts.Close()
}

func tearDown() {
	ts.Close()
}

func TestCheck(t *testing.T) {
	p := New()
	p.URL = ts.URL

	if err := p.Check(); err != nil {
		t.Errorf("failed to check header: %s", err)
	}
}
