package download_test

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai3-2/Mizushima/download"
)

var testdataPathMap = map[int][]string {
	0 : {"../documents/003","311"},
	1 : {"../documents/z4d4kWk.jpg", "146515"},
	// 2 : "../documents/http.request.txt",
}

func TestDownloader_SingleProcess(t *testing.T) {
	t.Helper()

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

		expected, err := os.ReadFile(testdataPathMap[0][0])
		if err != nil {
			t.Error(err)
		}

		if reflect.DeepEqual(actual, expected) {
			t.Errorf("expected %s, but got %s", expected, actual)
		}
	})
}

func TestDownloader_SingleProcessTimeout(t *testing.T) {
	t.Helper()

	ts, clean := newTestServer(t, nonRangeAccessTooLateHandler)
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
	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cancel()

	t.Run("case 1", func(t *testing.T) {
		actual := download.Downloader(urlObj, output, size, part, procs, isPara, tmpDirName, ctx)
		if err != nil {
			t.Error(err)
		}
		expected := fmt.Errorf("request.Request err: Get \"%s\": %w", urlObj, context.DeadlineExceeded)
		if actual.Error() != expected.Error() {
			t.Errorf("expected %s, \nbut got %s", expected, actual)
		}
	})
}

func TestDownloader_ParallelProcess(t *testing.T) {
	t.Helper()

	ts, clean := newTestServer(t, rangeAccessHandler)
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
	procs := uint(runtime.NumCPU())
	part := size / procs
	isPara := true
	tmpDirName := "test"
	ctx := context.Background()

	t.Run("case 1", func(t *testing.T) {
		if err := os.Mkdir(tmpDirName, 0775); err != nil {
			t.Error(err)
		}
		err := download.Downloader(urlObj, output, size, part, procs, isPara, tmpDirName, ctx)
		if err != nil {
			t.Errorf("err: %w", err)
		}
		defer func() {
			err := os.RemoveAll(tmpDirName)
			if err != nil {
				t.Error(err)
			}
		}()

		// make expected data
		byteData, err := os.ReadFile(testdataPathMap[1][0])
		if err != nil {
			t.Error(err)
		}

		for i := uint(0); i < procs; i++ {
			var start, end uint
			if i == 0 {
				start = 0
			} else {
				start = i * part + uint(1)
			}

			if end == procs -1 {
				end = size
			} else {
				end = (i + 1) * part
			}

			expected := byteData[start:end+1]
			fmt.Printf("length of expected: %d\n", len(expected))

			actual, err := os.ReadFile(tmpDirName+"/"+strconv.Itoa(int(i)))
			if err != nil {
				t.Error(err)
			}
			fmt.Printf("length of actual: %d\n", len(actual))

			if !reflect.DeepEqual(actual, expected) {
				t.Error("expected is not equal to actual")
			}
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

	body, err := os.ReadFile(testdataPathMap[0][0])
	if err != nil {
		t.Fatal(err)
	}
	w.Header().Set("Content-Length", testdataPathMap[0][1])
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, body)
}

func nonRangeAccessTooLateHandler(t *testing.T, w http.ResponseWriter, r *http.Request) {
	t.Helper()

	body, err := os.ReadFile(testdataPathMap[0][0])
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(3 * time.Second)
	w.Header().Set("Content-Length", testdataPathMap[0][1])
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, body)
}

func rangeAccessHandler(t *testing.T, w http.ResponseWriter, r *http.Request) {
	t.Helper()

	w.Header().Set("Content-Length", testdataPathMap[1][1])
	w.Header().Set("Access-Range", "bytes")
	
	rangeHeader := r.Header.Get("Range")

	body := retBody(t, rangeHeader, testdataPathMap[1][0])
	w.WriteHeader(http.StatusPartialContent)
	fmt.Fprint(w, body)
}

func retBody(t *testing.T, rangeHeader string, testdataPath string) []byte {
	b, err := os.ReadFile(testdataPath)
	if err != nil {
		t.Fatal(err)
	}

	if rangeHeader == "" {
		return b
	}

	rangeVals := strings.Split(rangeHeader, "=")
	if rangeVals[0] != "bytes" {
		t.Fatal(errors.New("err : Range header expected \"bytes\""))
	}

	rangeBytes := strings.Split(rangeVals[1], "-")
	start, err := strconv.Atoi(rangeBytes[0])
	if err != nil {
		t.Fatal(err)
	}

	end, err := strconv.Atoi(rangeBytes[1])
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("length of b: %d\n", len(b))
	fmt.Printf("length of b[start:end+1]: %d\n", len(b[start:end+1]))
	return b[start:end+1]
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
