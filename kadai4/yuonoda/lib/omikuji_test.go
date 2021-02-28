package omikuji_test

import (
	omiikuji "github.com/yuonoda/gopherdojo-studyroom/kadai4/yuonoda/lib"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"
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

func TestPickResult(t *testing.T) {
	cases := []struct {
		name        string
		time        time.Time
		resultRegex string
	}{
		{
			"normal",
			time.Date(2001, 2, 1, 0, 0, 0, 0, time.Local),
			`(吉|中吉|大吉|凶)`,
		},
		{
			"new_year",
			time.Date(2001, 1, 1, 0, 0, 0, 0, time.Local),
			`(大吉)`,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			regex, err := regexp.Compile(c.resultRegex)
			if err != nil {
				t.Errorf("incorect resultRegex : %s", err)
			}
			result := omiikuji.ExpPickResult(c.time)
			t.Logf("result:%s", result)
			isMatch := regex.MatchString(result)
			if !isMatch {
				t.Errorf("result doesn't match regex, got %s expected %s", result, c.resultRegex)
			}
		})
	}
}
