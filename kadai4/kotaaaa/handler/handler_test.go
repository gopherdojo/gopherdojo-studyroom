package handler

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	cases := []struct {
		name           string
		date           string
		expectedResult bool
		expected       string
		statusCode     int
	}{
		{name: "No date", date: "", statusCode: 200},
		{name: "12/31", date: "2020-12-31", statusCode: 200},
		{name: "New Year's Day", date: "2021-01-01", expectedResult: true, expected: "大吉", statusCode: 200},
		{name: "1/2", date: "2021-01-02", expectedResult: true, expected: "大吉", statusCode: 200},
		{name: "1/3", date: "2021-01-03", expectedResult: true, expected: "大吉", statusCode: 200},
		{name: "3/9", date: "2021-03-09", statusCode: 200},
		{name: "Not found a]", date: "2021-03-09", statusCode: 200},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/draw?p="+c.date, nil)
			httpHandler(w, r)
			rw := w.Result()
			defer rw.Body.Close()
			if rw.StatusCode != c.statusCode {
				t.Fatal("Unexpected status code")
			}

			var res ResModel
			dec := json.NewDecoder(rw.Body)
			if err := dec.Decode(&res); err != nil {
				t.Fatal(err)
			}
			if c.expectedResult && c.expected != res.Result {
				t.Errorf("Result is illigal. Expected %v, Result %v", c.expected, res.Result)
			}
		})
	}
}
