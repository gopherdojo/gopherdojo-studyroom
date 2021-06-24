package download_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestDownloader_SingleProcess(t *testing.T) {
	
}

func newTestServer(t *testing.T, 
	handler func(t *testing.T, w http.ResponseWriter, r *http.Request)) (*httptest.Server, func()) {
	
	t.Helper()

	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			handler(t, w, r)
		},
	))
	
	return ts, func() { ts.Close()}
}

func nonRangeAccessHandler(t *testing.T, w http.ResponseWriter, r *http.Request) {
	t.Helper()

	body, err := os.ReadFile("../003")
	if err != nil {
		t.Fatal(err)
	}
	w.Header().Set("Content-Length", "311")
	w.WriteHeader(http.StatusFound)
	fmt.Fprint(w, body)

}
