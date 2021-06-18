package convert

import (
	"convert"
	"path/filepath"
	"testing"
)

func TestAllfile(t *testing.T) {
	t.Helper()
	var s []string
	output, _ := GetAllFile("./img", s)
	if len(output) != 15 {
		t.Errorf("Image file does not complete readed,expected 15,but %v ", len(output)) //画像の枚数が違う
	}
}

func TestTableAllfile(t *testing.T) {
	t.Helper()
	var s []string
	var tests = []struct {
		pathinput string
		expect    int
	}{
		{"./img/img_test1", 3},
		{"./img/img_test2", 4},
		{"./img/img_test3", 5},
	}
	for _, te := range tests {
		output, _ := GetAllFile(te.pathinput, s)
		if len(output) != te.expect {
			t.Errorf("Image file does not complete readed,expected %v,but %v ", tests[1], len(output)) //画像の枚数が違う
		}
	}
}

func TestConv(t *testing.T) {
	t.Helper()
	err := convert.Conv("./img/1.0.jpg", "./img/1.0.png")
	if err != nil {
		t.Errorf("error happened %v", err)
	}
	matches, _ := filepath.Glob("./img/*.png")
	if matches == nil {
		t.Errorf("Image file does not converted %v", matches)
	}
}
