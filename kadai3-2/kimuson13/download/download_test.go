package download_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/kimuson13/gopherdojo-studyroom/kimuson13/download"
)

var currentTestdataName string
var registeredTestdatum = map[string][]byte{
	"foo.png":   readTestdata("foo.png"),
	"empty.txt": readTestdata("empty.txt"),
}

func TestDownloader_Run_Success(t *testing.T) {
	cases := map[string]struct {
		parallelism         int
		timeout             int
		output              string
		currentTestDataName string
	}{
		"normal": {
			parallelism:         3,
			timeout:             30,
			output:              "output.txt",
			currentTestDataName: "foo.png",
		},
		"highparallelism": {
			parallelism:         100,
			timeout:             30,
			output:              "output.txt",
			currentTestDataName: "foo.png",
		},
		"lowparallelism": {
			parallelism:         2,
			timeout:             30,
			output:              "output.txt",
			currentTestDataName: "foo.png",
		},
		"shortname": {
			parallelism:         3,
			timeout:             30,
			output:              "a",
			currentTestDataName: "foo.png",
		},
	}

	for n, c := range cases {
		c := c
		t.Run(n, func(t *testing.T) {
			currentTestdataName = c.currentTestDataName

			output, clean := createTempOutput(t, c.output)
			defer clean()

			ts, closefunc := newTestServer(t, normalHandler)
			defer closefunc()

			opt := &download.Options{
				Parallel: c.parallelism,
				Timeout:  c.timeout,
				Filename: output,
				URL:      ts.URL,
			}

			downloader := download.New(opt)

			err := downloader.Run(context.Background())
			if err != nil {
				t.Fatalf("err: %s", err)
			}

			before, err := os.Stat(path.Join("_testdata", currentTestdataName))
			if err != nil {
				panic(err)
			}

			after, err := os.Stat(output)
			if err != nil {
				panic(err)
			}

			if after.Name() != c.output {
				t.Fatalf("downloading file name is %v, but expected is %v", after.Name(), c.output)
			}

			if before.Size() != after.Size()-1 {
				t.Fatalf("it is not same %d and %d", before.Size(), after.Size()-1)
			}
		})
	}
}

func TestDownloader_With_Errors(t *testing.T) {
	cases := map[string]struct {
		parallelism         int
		timeout             int
		expected            error
		handler             func(t *testing.T, w http.ResponseWriter, r *http.Request)
		currentTestDataName string
	}{
		"parallelValidationError0": {
			parallelism:         0,
			timeout:             30,
			expected:            download.ErrValidateParallelism,
			handler:             normalHandler,
			currentTestDataName: "foo.png",
		},
		"parallelValidationError1": {
			parallelism:         1,
			timeout:             30,
			expected:            download.ErrValidateParallelism,
			handler:             normalHandler,
			currentTestDataName: "foo.png",
		},
		"timeoutValidationError": {
			parallelism:         3,
			timeout:             0,
			expected:            download.ErrValidateTimeout,
			handler:             normalHandler,
			currentTestDataName: "foo.png",
		},
		"noContent-LengthError": {
			parallelism:         3,
			timeout:             30,
			expected:            download.ErrNoContentLength,
			handler:             normalHandler,
			currentTestDataName: "empty.txt",
		},
		"notIncludeRange-AccessError": {
			parallelism:         3,
			timeout:             30,
			expected:            download.ErrNotIncludeRangeAccess,
			handler:             notIncludeAcceptRangeHandler,
			currentTestDataName: "foo.png",
		},
		"notContentInAccpetRange": {
			parallelism:         3,
			timeout:             30,
			expected:            download.ErrNotContent,
			handler:             notcontentAcceptRangeHeaderHandler,
			currentTestDataName: "foo.png",
		},
	}

	for n, c := range cases {
		c := c
		currentTestdataName = c.currentTestDataName
		t.Run(n, func(t *testing.T) {
			testForErrors(t, c.parallelism, c.timeout, c.expected, c.handler)
		})
	}
}

func testForErrors(t *testing.T, parallel, timeout int, expected error, handler func(t *testing.T, w http.ResponseWriter, r *http.Request)) {
	t.Helper()

	output, clean := createTempOutput(t, "output.txt")
	defer clean()

	ts, closefunc := newTestServer(t, handler)
	defer closefunc()

	opt := &download.Options{
		Parallel: parallel,
		Timeout:  timeout,
		Filename: output,
		URL:      ts.URL,
	}

	downloader := download.New(opt)

	actual := downloader.Run(context.Background())
	if actual != expected {
		t.Errorf("expected: %v, but got: %v", expected, actual)
	}
}

func newTestServer(t *testing.T, handler func(t *testing.T, w http.ResponseWriter, r *http.Request)) (*httptest.Server, func()) {
	t.Helper()

	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			handler(t, w, r)
		},
	))

	return ts, func() { ts.Close() }
}

func normalHandler(t *testing.T, w http.ResponseWriter, r *http.Request) {
	t.Helper()

	w.Header().Set("Accept-Ranges", "bytes")

	rangeHeader := r.Header.Get("Range")

	body := func() []byte {
		if rangeHeader == "" {
			return registeredTestdatum[currentTestdataName]
		}

		eqlSplitVals := strings.Split(rangeHeader, "=")
		if eqlSplitVals[0] != "bytes" {
			t.Fatalf("err: %s", eqlSplitVals[1])
		}

		c := strings.Split(eqlSplitVals[1], "-")

		min, err := strconv.Atoi(c[0])
		if err != nil {
			t.Fatalf("err: %s", err)
		}

		max, err := strconv.Atoi(c[1])
		if err != nil {
			t.Fatalf("err: %s", err)
		}

		return registeredTestdatum[currentTestdataName][min : max+1]
	}()

	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(body)))

	w.WriteHeader(http.StatusPartialContent)
	w.Write(body)
}

func notIncludeAcceptRangeHandler(t *testing.T, w http.ResponseWriter, r *http.Request) {
	t.Helper()

	body := registeredTestdatum[currentTestdataName]
	w.Write(body)
}

func notcontentAcceptRangeHeaderHandler(t *testing.T, w http.ResponseWriter, r *http.Request) {
	t.Helper()

	w.Header().Set("Accept-Ranges", "none")

	body := registeredTestdatum[currentTestdataName]
	w.Write(body)
}

func readTestdata(filename string) []byte {
	b, err := ioutil.ReadFile(path.Join("_testdata", filename))
	if err != nil {
		panic(err)
	}

	return b
}

func createTempOutput(t *testing.T, name string) (string, func()) {
	t.Helper()

	dir, err := ioutil.TempDir("", "parallel-download")
	if err != nil {
		panic(err)
	}

	return filepath.Join(dir, name), func() { os.RemoveAll(dir) }
}
