package downloading

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"split_download/opt"
	"strconv"
	"strings"
	"testing"
	"time"
)

var registeredTestdatum = map[string]string{
	"foo.png":   readTestdata("foo.png"),
	"a.txt":     readTestdata("a.txt"),
	"empty.txt": readTestdata("empty.txt"),
}

var currentTestdataName string

func TestDownloading_Download_Success(t *testing.T) {
	cases := map[string]struct {
		parallelism         int
		currentTestdataName string
	}{
		"normal":                      {parallelism: 3, currentTestdataName: "foo.png"},
		"parallelism < 1":             {parallelism: 0, currentTestdataName: "a.txt"},
		"contentLength < parallelism": {parallelism: 4, currentTestdataName: "a.txt"},
	}

	for n, c := range cases {
		c := c
		t.Run(n, func(t *testing.T) {
			parallelism := c.parallelism
			currentTestdataName = c.currentTestdataName

			output, clean := createTempOutput(t)
			defer clean()

			ts, clean := newTestServer(t, normalHandler)
			defer clean()

			err := newDownloader(t, output, ts, parallelism).Download(context.Background())
			if err != nil {
				t.Fatalf("err %s", err)
			}
		})
	}

}

func TestDownloading_Download_NoContent(t *testing.T) {
	expected := errNoContent

	currentTestdataName = "empty.txt"

	output, clean := createTempOutput(t)
	defer clean()

	ts, clean := newTestServer(t, normalHandler)
	defer clean()

	actual := newDownloader(t, output, ts, 1).Download(context.Background())
	if actual != expected {
		t.Errorf(`unexpected error: expected: "%s" actual: "%s"`, expected, actual)
	}
}

func TestDownloading_Download_AcceptRangesHeaderNotFound(t *testing.T) {
	expected := errResponseDoesNotIncludeAcceptRangesHeader

	output, clean := createTempOutput(t)
	defer clean()

	ts, clean := newTestServer(t, func(t *testing.T, w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "") })
	defer clean()

	actual := newDownloader(t, output, ts, 8).Download(context.Background())
	if actual != expected {
		t.Errorf(`unexpected error: expected: "%s" actual: "%s"`, expected, actual)
	}
}

func TestDownloading_Download_AcceptRangesHeaderSupportsBytesOnly(t *testing.T) {
	expected := errValueOfAcceptRangesHeaderIsNotBytes

	output, clean := createTempOutput(t)
	defer clean()

	ts, clean := newTestServer(t, func(t *testing.T, w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Accept-Ranges", "none")
		fmt.Fprint(w, "")
	})
	defer clean()

	actual := newDownloader(t, output, ts, 8).Download(context.Background())
	if actual != expected {
		t.Errorf(`unexpected error: expected: "%s" actual: "%s"`, expected, actual)
	}
}

func TestDownloading_Download_BadRequest(t *testing.T) {
	expected := "unexpected status code: 400"

	output, clean := createTempOutput(t)
	defer clean()

	ts, clean := newTestServer(t, func(t *testing.T, w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Accept-Ranges", "bytes")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "bad request")
	})
	defer clean()

	err := newDownloader(t, output, ts, 8).Download(context.Background())
	if err == nil {
		t.Fatalf("unexpectedly err is nil")
	}
	actual := err.Error()
	if actual != expected {
		t.Errorf(`unexpected error: expected: "%s" actual: "%s"`, expected, actual)
	}
}

func TestDownloading_Download_RenameError(t *testing.T) {
	currentTestdataName = "foo.png"

	ts, clean := newTestServer(t, normalHandler)
	defer clean()

	err := newDownloader(t, "/non/existent/path", ts, 1).Download(context.Background())
	if err == nil {
		t.Fatal("unexpectedly err is nil")
	}

	if !regexp.MustCompile("/non/existent/path: no such file or directory").MatchString(err.Error()) {
		t.Errorf("unexpectedly not matched: %s", err.Error())
	}
}

func TestDownloading_getContentLength_DoError(t *testing.T) {
	expected := `Head "A": unsupported protocol scheme ""`

	ts, clean := newTestServer(t, noopHandler)
	defer clean()

	d := newDownloader(t, "path/to/output", ts, 2)

	u, err := url.Parse("A")
	if err != nil {
		t.Fatalf("err %s", err)
	}

	d.url = u

	_, err = d.getContentLength(context.Background())

	actual := err.Error()
	if actual != expected {
		t.Errorf(`unexpected error: expected: "%s" actual: "%s"`, expected, actual)
	}
}

func TestDownloading_partialDownload_osCreateError(t *testing.T) {
	ts, clean := newTestServer(t, func(t *testing.T, w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusPartialContent)
	})
	defer clean()

	_, err := newDownloader(t, "", ts, 2).partialDownload(context.Background(), "bytes=0-1", "non/existent/path")
	if !regexp.MustCompile("no such file or directory").MatchString(err.Error()) {
		t.Errorf("unexpectedly not matched: %s", err.Error())
	}
}

func TestDownloading_concat_osCreateError(t *testing.T) {
	ts, clean := newTestServer(t, noopHandler)
	defer clean()

	d := newDownloader(t, "", ts, 2)
	_, err := d.concat(map[int]string{}, "non/existent/path")

	if !regexp.MustCompile("no such file or directory").MatchString(err.Error()) {
		t.Errorf("unexpectedly not matched: %s", err.Error())
	}
}

func TestDownloading_concat_osOpenError(t *testing.T) {
	ts, clean := newTestServer(t, noopHandler)
	defer clean()

	dir, err := ioutil.TempDir("", "parallel-download")
	if err != nil {
		t.Fatalf("err %s", err)
	}
	defer os.RemoveAll(dir)

	_, err = newDownloader(t, "", ts, 2).concat(map[int]string{0: "non/existent/path"}, dir)

	if !regexp.MustCompile("no such file or directory").MatchString(err.Error()) {
		t.Errorf("unexpectedly not matched: %s", err.Error())
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

	rangeHdr := r.Header.Get("Range")

	body := func() string {
		if rangeHdr == "" {
			return registeredTestdatum[currentTestdataName]
		}

		eqlSplitVals := strings.Split(rangeHdr, "=")
		if eqlSplitVals[0] != "bytes" {
			t.Fatalf("err %s", eqlSplitVals[0])
		}

		c := strings.Split(eqlSplitVals[1], "-")

		min, err := strconv.Atoi(c[0])
		if err != nil {
			t.Fatalf("err %s", err)
		}

		max, err := strconv.Atoi(c[1])
		if err != nil {
			t.Fatalf("err %s", err)
		}

		return registeredTestdatum[currentTestdataName][min : max+1]
	}()

	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(body)))

	w.WriteHeader(http.StatusPartialContent)

	fmt.Fprint(w, body)
}

func noopHandler(t *testing.T, w http.ResponseWriter, r *http.Request) {}

func newDownloader(t *testing.T, output string, ts *httptest.Server, parallelism int) *Downloader {
	t.Helper()

	opts := &opt.Options{
		Parallelism: parallelism,
		Output:      output,
		URL:         mustParseRequestURI(t, ts.URL),
		Timeout:     60 * time.Second,
	}

	return NewDownloader(ioutil.Discard, opts)
}

func mustParseRequestURI(t *testing.T, s string) *url.URL {
	t.Helper()

	u, err := url.ParseRequestURI(s)
	if err != nil {
		t.Fatalf("err %s", err)
	}

	return u
}

func readTestdata(filename string) string {
	b, err := ioutil.ReadFile(path.Join("testdata", filename))
	if err != nil {
		panic(err)
	}
	return string(b)
}

func createTempOutput(t *testing.T) (string, func()) {
	t.Helper()

	dir, err := ioutil.TempDir("", "parallel-download")
	if err != nil {
		panic(err)
	}

	return filepath.Join(dir, "output.txt"), func() { os.RemoveAll(dir) }
}