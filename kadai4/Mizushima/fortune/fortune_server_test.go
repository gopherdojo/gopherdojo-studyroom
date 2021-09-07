package fortune_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai4/Mizushima/fortune"
)

func TestOmikujiHandlerTheFirstThreeDays(t *testing.T) {
	cases := map[string]struct {
		date     time.Time
		expected string
	}{
		"the 1st of the first three days of the new year.": {
			date:     time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local),
			expected: "大吉",
		},
		"the 2nd of the first three days of the new year.": {
			date:     time.Date(2021, 1, 2, 0, 0, 0, 0, time.Local),
			expected: "大吉",
		},
		"the 3rd of the first three days of the new year.": {
			date:     time.Date(2021, 1, 3, 0, 0, 0, 0, time.Local),
			expected: "大吉",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/", nil)

			fortune.SetTime(c.date)

			fortune.OmikujiHandler(w, r)
			rw := w.Result()
			defer func() {
				if err := rw.Body.Close(); err != nil {
					t.Fatal(err)
				}
			}()

			if rw.StatusCode != http.StatusOK {
				t.Fatalf("unexpected status code: %d", rw.StatusCode)
			}

			var res fortune.Res
			dec := json.NewDecoder(rw.Body)
			if err := dec.Decode(&res); err != nil && err != io.EOF {
				t.Fatal(err)
			}

			if res.Result != c.expected {
				t.Fatalf("expected %s, but got %s", res.Result, c.expected)
			}
		})
	}
}

func TestOmikujiHandlerTheOtherDays(t *testing.T) {
	cases := map[string]struct {
		date     time.Time
		results  [4]string
		expected string
	}{
		"the other day 1": {
			date:     time.Date(2021, 1, 4, 0, 0, 0, 0, time.Local),
			results:  [4]string{"a", "a", "a", "a"},
			expected: "a",
		},
		"the other day 2": {
			date:     time.Date(2021, 8, 25, 0, 0, 0, 0, time.Local),
			results:  [4]string{"b", "b", "b", "b"},
			expected: "b",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/", nil)

			fortune.SetTime(c.date)
			fortune.SetResOmikuji(c.results)

			fortune.OmikujiHandler(w, r)
			rw := w.Result()
			defer func() {
				if err := rw.Body.Close(); err != nil {
					t.Fatal(err)
				}
			}()

			if rw.StatusCode != http.StatusOK {
				t.Fatalf("unexpected status code: %d", rw.StatusCode)
			}

			var res fortune.Res
			dec := json.NewDecoder(rw.Body)
			if err := dec.Decode(&res); err != nil && err != io.EOF {
				t.Fatal(err)
			}

			if res.Result != c.expected {
				t.Fatalf("expected %s, but got %s", res.Result, c.expected)
			}
		})
	}
}

func contains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}
