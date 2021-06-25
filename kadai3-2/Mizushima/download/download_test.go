package download_test

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"testing"

	"github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai3-2/Mizushima/download"
)

var testdataPathMap = map[int]string {
	0 : "../documents/003",
	1 : "../documents/z4d4kWk.jpg",
	2 : "../documents/http.request.txt",
}

func TestDownloader_SingleProcess(t *testing.T) {
	ts, clean := newTestServer(t, nonRangeAccessHandler)
	defer clean()

	// get a url.URL object
	urlObj := getURLObject(t, ts.URL)

	// get a file for output
	output, clean := makeTempFile(t)
	defer clean()

	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	// get the file size to be downloaded.
	size := GetSize(t, resp)

	// this test is non-parallel download.
	part := size
	procs := uint(1)
	isPara := false
	tmpDirName := ""
	ctx := context.Background()

	t.Run("case 1", func(t *testing.T) {
		err := download.Downloader(urlObj, output, size, part, procs, isPara, tmpDirName, ctx)
		if err != nil {
			t.Error(err)
		}
		actual := new(bytes.Buffer).Bytes()
		_, err = output.Read(actual)
		if err != nil {
			t.Error(err)
		}
		expected, err := os.ReadFile(testdataPathMap[0])
		if err != nil {
			t.Error(err)
		}

		if reflect.DeepEqual(actual, expected) {
			t.Errorf("expected %s, but got %s", expected, actual)
		}
	})

}

func newTestServer(t *testing.T, 
	handler func(t *testing.T, w http.ResponseWriter, r *http.Request)) (*httptest.Server, func()) {
	
	t.Helper()

	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			handler(t, w, r)
		},
	))
	
	return ts, func() { ts.Close() }
}

func nonRangeAccessHandler(t *testing.T, w http.ResponseWriter, r *http.Request) {
	t.Helper()

	body, err := os.ReadFile(testdataPathMap[0])
	if err != nil {
		t.Fatal(err)
	}
	w.Header().Set("Content-Length", "311")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, body)
}

func getURLObject(t *testing.T, urlStr string) *url.URL {
	t.Helper()

	urlObj, err := url.ParseRequestURI(urlStr)
	if err != nil {
		t.Error(err)
	}

	return urlObj
}

func makeTempFile(t *testing.T) (*os.File, func()) {
	t.Helper()

	dir, err := ioutil.TempDir("", "test_download")
	if err != nil {
		t.Error(err)
	}

	out, err := os.Create(dir+"/test")
	if err != nil {
		t.Error(err)
	}

	return out, 
		func() {
			err = out.Close()
			if err != nil {
				t.Error(err)
			}
			err = os.RemoveAll(dir)
			if err != nil {
				t.Error(err)
			}
		}
}

// GetSize returns size from response header.
func GetSize(t *testing.T, r *http.Response) uint {
	t.Helper()

	contLen, is := r.Header["Content-Length"]
	// fmt.Println(h)
	if !is {
		t.Errorf("cannot find Content-Length header")
	}
	
	ret, err := strconv.ParseUint(contLen[0], 10, 32)
	if err != nil {
		t.Error(err)
	}
	return uint(ret)
}
