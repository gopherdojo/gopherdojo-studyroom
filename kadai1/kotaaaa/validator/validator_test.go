package validator

import (
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
		basePath           string
		imgFormat          string
		convertedImgFormat string
		errRaisedFlag      bool
	}{
		{"Success1", "../testdata", ".png", ".jpg", false},
		{"Success2", "../testdata", ".jpg", ".gif", false},
		{"Success2", "../testdata", ".gif", ".png", false},
		{"Fail1 .svg is not supported.", "../testdata", ".jpg", ".svg", true},
		{"Fail2 path is not correct.", "../not_existed", ".png", ".jpg", true},
	}

	for _, tt := range imgDirDataTests {
		t.Run(tt.caseName, func(t *testing.T) {

			err := ValidateArgs(tt.basePath, tt.imgFormat, tt.convertedImgFormat)
			if errExists(t, err) != tt.errRaisedFlag {
				t.Errorf("error: %#v", err)
			}
		})
	}
}
