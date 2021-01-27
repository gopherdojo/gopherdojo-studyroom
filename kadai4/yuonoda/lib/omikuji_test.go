package omikuji_test

import (
	omiikuji "github.com/yuonoda/gopherdojo-studyroom/kadai4/yuonoda/lib"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	// テストリクエストを送信
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	omiikuji.ExpHandler(w, r)

	// レスポンスを判定
	rw := w.Result()
	defer rw.Body.Close()
	if rw.StatusCode != http.StatusOK {
		t.Fatal("unexpected status code")
	}
}
