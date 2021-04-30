package main

import (
	"bytes"
	"fmt"
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

func TestDownload(t *testing.T) {
	p := New()
	p.URL = ts.URL
	p.TargetDir = "testdata/test_download"
	p.Utils = &Data{
		filename: "header.jpg",
	}

	if err := p.Check(); err != nil {
		t.Errorf("failed to check header: %s", err)
	}

	if err := p.Download(); err != nil {
		t.Errorf("failed to download: %s", err)
	}

	for i := 0; i < p.Procs; i++ {
		filename := fmt.Sprintf("testdata/test_download/header.jpg.%d.%d", p.Procs, i)
		_, err := os.Stat(filename)
		if err != nil {
			t.Errorf("file not exist: %s", err)
		}
	}
}

// utils.goにあるメソッドをテストするのは違和感があるがこのファイルの中でテストを行う
func TestMergeFiles(t *testing.T) {
	p := New()
	p.URL = ts.URL
	p.Utils = &Data{
		filename:     "header.jpg",
		dirname:      "testdata/test_download",
		fullfilename: "testdata/test_download/header.jpg",
	}

	if err := p.MergeFiles(p.Procs); err != nil {
		t.Errorf("failed to MergeFiles: %s", err)
	}
}
