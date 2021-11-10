package search

import (
	"sort"
	"testing"
)

func TestGetFiles(t *testing.T) {

	ext := ".jpg"
	files, err := GetFiles("../testdata/", ext)
	if err != nil {
		t.Errorf("GetFiles doesn't return correct value")
	}
	if len(files) != 2 {
		t.Errorf("Invalid result. Expected: number of file:2 Actual file: %v", len(files))
	}
	sort.Strings(files)
	if files[0] == "owl.jpg" {
		t.Errorf("This is not expected file. %v", files[0])
	}
	if files[1] == "owl3.jpg" {
		t.Errorf("This is not expected file. %v", files[1])
	}
}
