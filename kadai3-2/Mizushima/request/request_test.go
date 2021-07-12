package request_test

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai3-2/Mizushima/request"
)

var testdataPathMap = map[int][]string{
	0: {"../documents/003", mustGetSize("../documents/003")},
	1: {"../documents/z4d4kWk.jpg", mustGetSize("../documents/z4d4kWk.jpg")},
}

func Test_Request(t *testing.T) {
	t.Helper()

	cases := []struct {
		name     string
		key      int // key for testdataPathMap
		handler  func(t *testing.T, w http.ResponseWriter, r *http.Request, testDataKey int)
		expected *http.Response
	}{
		{
			name:     "case 1",
			key:      0,
			handler:  nonRangeAccessHandler,
			expected: &http.Response{},
		},
		{
			name:     "case 2",
			key:      1,
			handler:  rangeAccessHandler,
			// http.Responseを作る方法を調べる
			expected: &http.Response{
				Status: "206 Partial Content",
				StatusCode: 206,
				Proto: "HTTP/1.1",
			},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			ts, clean := newTestServer(t, c.handler, c.key)
			defer clean()
			actual, err := request.Request(context.Background(), "GET", ts.URL, "", "")
			if err != nil {
				t.Fatal(err)
			}
			fmt.Println("actual:", actual.Header)
			if !reflect.DeepEqual(actual.Header, c.expected.Header) {
				dumped_expected, err := httputil.DumpResponse(c.expected, false)
				if err != nil {
					t.Fatal(err)
				}
				dumped_actual, err := httputil.DumpResponse(actual, false)
				if err != nil {
					t.Fatal(err)
				}
				t.Errorf("expected,\n%vbut got,\n%v", string(dumped_expected), string(dumped_actual))
			}
		})
	}
}

func newTestServer(t *testing.T,
	handler func(t *testing.T, w http.ResponseWriter, r *http.Request, testDataKey int),
	testDataKey int) (*httptest.Server, func()) {

	t.Helper()

	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			handler(t, w, r, testDataKey)
		},
	))

	return ts, func() { ts.Close() }
}

func nonRangeAccessHandler(t *testing.T, w http.ResponseWriter, r *http.Request, testDataKey int) {
	t.Helper()

	body, err := os.ReadFile(testdataPathMap[testDataKey][0])
	if err != nil {
		t.Fatal(err)
	}
	w.Header().Set("Content-Length", testdataPathMap[testDataKey][1])
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, body)
}

func nonRangeAccessTooLateHandler(t *testing.T, w http.ResponseWriter, r *http.Request, testDataKey int) {
	t.Helper()

	body, err := os.ReadFile(testdataPathMap[testDataKey][0])
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(3 * time.Second)
	w.Header().Set("Content-Length", testdataPathMap[testDataKey][1])
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, body)
}

func rangeAccessHandler(t *testing.T, w http.ResponseWriter, r *http.Request, testDataKey int) {
	t.Helper()

	w.Header().Set("Content-Length", testdataPathMap[testDataKey][1])
	w.Header().Set("Access-Range", "bytes")

	rangeHeader := r.Header.Get("Range")

	body := retBody(t, rangeHeader, testdataPathMap[testDataKey][0])
	w.WriteHeader(http.StatusPartialContent)
	fmt.Fprint(w, body)
}

func retBody(t *testing.T, rangeHeader string, testDataPath string) []byte {
	t.Helper()

	b, err := os.ReadFile(testDataPath)
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

	// fmt.Printf("length of b: %d\n", len(b))
	// fmt.Printf("length of b[start:end+1]: %d\n", len(b[start:end+1]))
	return b[start : end+1]
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
		t.Fatal(err)
	}

	out, err := os.Create(dir + "/test")
	if err != nil {
		t.Fatal(err)
	}

	return out,
		func() {
			err = out.Close()
			if err != nil {
				t.Fatal(err)
			}
			err = os.RemoveAll(dir)
			if err != nil {
				t.Fatal(err)
			}
		}
}

func mustGetSize(path string) string {
	
	fileinfo, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}

	return strconv.Itoa(int(fileinfo.Size()))
}

// GetSize returns size from response header.
func getSizeForTest(t *testing.T, r *http.Response) uint {
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
