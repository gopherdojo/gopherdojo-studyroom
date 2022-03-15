package imgconv_test

import (
	"fmt"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"example.com/ex01/imgconv"
)

func isPng(image string) (bool, error) {
	f, err := os.Open(image)
	if err != nil {
		return false, err
	}
	_, err = png.Decode(f)
	if err != nil {
		return false, fmt.Errorf("%s is not png\n", image)
	}
	defer f.Close()
	defer os.Remove(image)
	return true, nil
}

func isGif(image string) (bool, error) {
	f, err := os.Open(image)
	if err != nil {
		return false, err
	}
	_, err = gif.Decode(f)
	if err != nil {
		return false, fmt.Errorf("%s is not gif\n", image)
	}
	defer f.Close()
	defer os.Remove(image)
	return true, nil
}

func isJpg(image string) (bool, error) {
	f, err := os.Open(image)
	if err != nil {
		return false, err
	}
	_, err = jpeg.Decode(f)
	if err != nil {
		return false, fmt.Errorf("%s is not jpg\n", image)
	}
	defer f.Close()
	defer os.Remove(image)
	return true, nil
}

func replaceExt(path, ext string) string {
	return strings.TrimSuffix(path, filepath.Ext(path)) + "." + ext
}

func TestExportConvert(t *testing.T) {
	tests := []struct {
		name    string
		arg1    string
		arg2    string
		wantErr bool
	}{
		// TODO: Add error test cases.
		// {"jpg to png", "../testdata/ExportConvert/jpg1.jpg", "../testdata/ExportConvert/jpg1.png", false},
		// {"jpg to jpeg", "../testdata/ExportConvert/jpg2.jpg", "../testdata/ExportConvert/jpg2.jpeg", false},
		// {"jpg to gif", "../testdata/ExportConvert/jpg3.jpg", "../testdata/ExportConvert/jpg3.gif", false},
		// {"jpeg to gif", "../testdata/ExportConvert/jpeg1.jpeg", "../testdata/ExportConvert/jpeg1.png", false},
		// {"jpeg to jpg", "../testdata/ExportConvert/jpeg2.jpeg", "../testdata/ExportConvert/jpeg2.jpg", false},
		// {"jpeg to gif", "../testdata/ExportConvert/jpeg3.jpeg", "../testdata/ExportConvert/jpeg3.gif", false},
		// {"png to jpg", "../testdata/ExportConvert/png1.png", "../testdata/ExportConvert/png1.jpg", false},
		// {"png to jpeg", "../testdata/ExportConvert/png2.png", "../testdata/ExportConvert/png2.jpeg", false},
		// {"png to gif", "../testdata/ExportConvert/png3.png", "../testdata/ExportConvert/png3.gif", false},
		// {"gif to jpg", "../testdata/ExportConvert/gif1.gif", "../testdata/ExportConvert/gif1.jpg", false},
		// {"gif to jpeg", "../testdata/ExportConvert/gif2.gif", "../testdata/ExportConvert/gif2.jpeg", false},
		// {"gif to png", "../testdata/ExportConvert/gif3.gif", "../testdata/ExportConvert/gif3.png", false},
		// {"secret", "../testdata/ExportConvert/.secret.jpg", "../testdata/ExportConvert/.secret.png", false},
		{"jpg to png", "../testdata/ExportConvert/jpg1.jpg", "png", false},
		{"jpg to jpeg", "../testdata/ExportConvert/jpg2.jpg", "jpeg", false},
		{"jpg to gif", "../testdata/ExportConvert/jpg3.jpg", "gif", false},
		{"jpeg to gif", "../testdata/ExportConvert/jpeg1.jpeg", "png", false},
		{"jpeg to jpg", "../testdata/ExportConvert/jpeg2.jpeg", "jpg", false},
		{"jpeg to gif", "../testdata/ExportConvert/jpeg3.jpeg", "gif", false},
		{"png to jpg", "../testdata/ExportConvert/png1.png", "jpg", false},
		{"png to jpeg", "../testdata/ExportConvert/png2.png", "jpeg", false},
		{"png to gif", "../testdata/ExportConvert/png3.png", "gif", false},
		{"gif to jpg", "../testdata/ExportConvert/gif1.gif", "jpg", false},
		{"gif to jpeg", "../testdata/ExportConvert/gif2.gif", "jpeg", false},
		{"gif to png", "../testdata/ExportConvert/gif3.gif", "png", false},
		{"secret", "../testdata/ExportConvert/.secret.jpg", "png", false},
		// {"invalid format", "../testdata/ExportConvert/.secret.jpg", "bmp", true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := imgconv.ExportConvert(tt.arg1, tt.arg2)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExportConvert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			path := replaceExt(tt.arg1, tt.arg2)
			if strings.HasSuffix(path, ".png") {
				if ok, err := isPng(path); !ok {
					t.Errorf("ExportConvert() can't convert: error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			} else if strings.HasSuffix(path, ".gif") {
				if ok, err := isGif(path); !ok {
					t.Errorf("ExportConvert() can't convert: error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			} else if strings.HasSuffix(path, ".jpg") || strings.HasSuffix(path, ".jpeg") {
				if ok, err := isJpg(path); !ok {
					t.Errorf("ExportConvert() can't convert: error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
		})
	}
}
