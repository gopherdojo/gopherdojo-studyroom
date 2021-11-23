package slices_test

import (
	"testing"

	"github.com/exaream/gopherdojo-studyroom/kadai1/exaream/slices"
)

func TestSlices(t *testing.T) {
	list := []string{"apple", "orange", "banana"}
	cases := map[string]struct {
		list []string
		str  string
		want bool
	}{
		"exit":     {list: list, str: "apple", want: true},
		"not exit": {list: list, str: "tomato", want: false},
	}
	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			if slices.Contains(tt.list, tt.str) != tt.want {
				t.Error()
			}
		})
	}
}
