package main

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

const testDir = "testdata/test_download"

var (
	dir      string
	mkdirErr error
	ts       *httptest.Server
)

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

	dir, mkdirErr = os.MkdirTemp(testDir, "")
	if mkdirErr != nil {
		panic(mkdirErr)
	}
}

func tearDown() {
	ts.Close()
	os.RemoveAll(dir)
}

func TestCheck(t *testing.T) {
	p := New()
	p.URL = ts.URL

	if err := p.Check(context.Background(), dir); err != nil {
		t.Errorf("failed to check header: %s", err)
	}
}

func TestDownload(t *testing.T) {
	p := New()
	p.URL = ts.URL
	p.TargetDir = testDir
	p.filename = "header.jpg"

	err := p.Check(context.Background(), dir)
	if err != nil {
		t.Errorf("failed to check header: %s", err)
	}

	if err := p.Download(context.Background()); err != nil {
		t.Errorf("failed to download: %s", err)
	}

	for i := 0; i < p.Procs; i++ {
		filename := fmt.Sprintf(p.dirname+"/header.jpg.%d.%d", p.Procs, i)
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
	p.TargetDir = testDir
	p.filename = "header.jpg"
	p.fullfilename = "testdata/test_download/header.jpg"

	err := p.Check(context.Background(), dir)
	if err != nil {
		t.Errorf("failed to check header: %s", err)
	}

	if err := p.Download(context.Background()); err != nil {
		t.Errorf("failed to download: %s", err)
	}

	if err := mergeFiles(p.Procs, p.filename, p.dirname, p.fullfilename); err != nil {
		t.Errorf("failed to MergeFiles: %s", err)
	}
}
