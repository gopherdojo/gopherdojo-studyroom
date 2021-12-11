package omikuji_test

import (
	"net/http"
	"net/http/httptest"
	"omikuji/omikuji"
	"testing"
	"time"
)

func TestHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	omikuji.Handler(w, r)
	rw := w.Result()
	defer rw.Body.Close()

	if rw.StatusCode != http.StatusOK {
		t.Fatal("unexpected status code")
	}
}

func TestOmikujiInNewYear(t *testing.T) {
	cases := map[string]struct {
		date     string
		expected string
	}{
		"january1": {
			"2021-01-01",
			"大吉",
		},
		"january2": {
			"2021-01-02",
			"大吉",
		},
		"january3": {
			"2021-01-03",
			"大吉",
		},
		"otherYear": {
			"2020-01-01",
			"大吉",
		},
	}

	for n, c := range cases {
		c := c
		t.Run(n, func(t *testing.T) {
			time, err := time.Parse(omikuji.Layout, c.date)
			if err != nil {
				t.Fatal("time parse error")
			}

			actual := omikuji.PickOmikuji(time)

			if c.expected != actual {
				t.Fatalf("expected = %s, but got %s", c.expected, actual)
			}
		})
	}
}
