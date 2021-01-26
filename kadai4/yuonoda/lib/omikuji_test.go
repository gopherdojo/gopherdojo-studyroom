package omikuji_test

import (
	omiikuji "github.com/yuonoda/gopherdojo-studyroom/kadai4/yuonoda/lib"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	omiikuji.ExpHandler(w, r)
	rw := w.Result()
	defer rw.Body.Close()
	if rw.StatusCode != http.StatusOK {
		t.Fatal("unexpected status code")
	}
	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatal("unexpected error")
	}
	const expected = `{result:"大吉”}`
	if s := string(b); s != expected {
		t.Fatalf("unexpected response: %s", s)
	}
}
