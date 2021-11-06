package search

import (
	"path/filepath"
	"testing"
)

func TestGetFiles(t *testing.T) {

	extentions := []string{".txt", ".png", ".jpg", ".gif"}
	for _, ext := range extentions {
		files := GetFiles("../testdata/", ext)
		for _, file := range files {
			if filepath.Ext(file) != ext {
				t.Errorf("Invalid result. Expected: %v Actual file: %v", ext, file)
			}
		}
	}
}
