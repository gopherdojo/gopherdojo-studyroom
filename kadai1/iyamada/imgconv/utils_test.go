package imgconv_test

import (
	"testing"

	"example.com/ex01/imgconv"
)

func TestExportHasValidFileExtent(t *testing.T) {
	tests := []struct {
		name string
		in1  string
		in2  string
		out  bool
	}{
		// TODO: Add test cases.
		{"good", "hoge.jpg", "jpg", true},
		{"good", "hoge.jpg", "jpeg", true},
		{"good", "hoge.png", "png", true},
		{"fail", "hoge.txt", "png", false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := imgconv.ExportHasValidFileExtent(tt.in1, tt.in2)
			if got != tt.out {
				t.Errorf("Parse() got = %v, expected %v", got, tt.out)
				return
			}
		})
	}
}
