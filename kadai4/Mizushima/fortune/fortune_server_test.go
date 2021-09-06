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

func TestOmikujiHandler(t *testing.T) {
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
