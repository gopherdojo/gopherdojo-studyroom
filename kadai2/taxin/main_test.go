package main

import (
	"flag"
	"testing"
)

func errExists(t *testing.T, err error) bool {
	t.Helper()
	if err != nil {
		return true
	}
	return false
}

func TestValidateArgs(t *testing.T) {
	imgDirDataTests := []struct {
		caseName           string
		imgFormat          string
		convertedImgFormat string
		dirPath            string
		errExistsFlag      bool
	}{
		{"case1", "png", "jpg", "testdata", false},
		{"case2", "jpg", "svg", "testdata", true},
		{"case3", "png", "jpg", "testdatadata", true},
		{"case4", "jpg", "gif", "testdata", false},
	}

	for _, tt := range imgDirDataTests {
		t.Run(tt.caseName, func(t *testing.T) {
			flag.CommandLine.Set("fmt", tt.imgFormat)
			flag.CommandLine.Set("outfmt", tt.convertedImgFormat)
			flag.CommandLine.Set("dir", tt.dirPath)

			err := validateArgs()
			if errExists(t, err) != tt.errExistsFlag {
				t.Errorf("error: %#v", err)
			}
		})
	}
}
