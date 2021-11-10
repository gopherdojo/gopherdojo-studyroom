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
		{"png to jpg", "../testdata", ".png", ".jpg", false},
		{"jpg to gif", "../testdata", ".jpg", ".gif", false},
		{"gif to png", "../testdata", ".gif", ".png", false},
		{".svg is not supported.", "../testdata", ".jpg", ".svg", true},
		{"path is not correct.", "../not_existed", ".png", ".jpg", true},
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
