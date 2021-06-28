package request_test

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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
	0: {"../documents/003", "311"},
	1: {"../documents/z4d4kWk.jpg", "146515"},
	// 2 : "../documents/http.request.txt",
}

func Test_Request(t *testing.T) {
	t.Helper()

	cases := []struct {
		name string
		key int // key for testdataPathMap
		handler http.HandlerFunc
		expected *http.Response
	}{
		{
			name: "case 1",
			key: 0,
			handler: nonRangeAccessHandler,
			expected: &http.Response{},
		},
		{
			name: "case 2",
			key: 1,
			handler: rangeAccessHandler,
			expected: &http.Response{},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			ts, clean := newTestServer(t, c.handler, c.key)
			actual, err := request.Request(context.BackGround(), "GET", ts.URL, "", "")
			if reflect.DeepEqual(actual, c.expected) {
				t.Errorf("expected %v, but got %v", c.expected, actual)
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