package files_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/exaream/gopherdojo-studyroom/kadai1/exaream/files"
)

const (
	// Paths
	testDataDir  = "../testdata"
	srcDir       = testDataDir + "/src"
	testPathJpg  = srcDir + "/sample1.jpg"
	testPathJpeg = srcDir + "/sample2.jpeg"
	testPathPng  = srcDir + "/sample3.png"
	testPathGif  = srcDir + "/sample4.gif"
	testPathTif  = srcDir + "/sample5.tif"
	testPathTiff = srcDir + "/sample6.tiff"
	testPathBmp  = srcDir + "/sample7.bmp"

	// Extensions
	extJpg  = "jpg"
	extJpeg = "jpeg"
	extPng  = "png"
	extGif  = "gif"
	extTif  = "tif"
	extTiff = "tiff"
	extBmp  = "bmp"
	extNon  = "ext-non"

	// Other
	strEmpty = ""
)

var randomDir string = testDataDir + "/" + fmt.Sprint(time.Now().UnixNano())

func TestOpenFileNormal(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		path string
	}{
		{"1", testPathJpg},
		{"2", testPathJpeg},
		{"3", testPathPng},
		{"4", testPathGif},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if _, err := files.Open(test.path); err != nil {
				t.Errorf("failed to open %v", test.path)
			}
		})
	}
}

func TestOpenFileAbnormal(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		path string
	}{
		{"1", randomDir},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if _, err := files.Open(test.path); os.IsExist(err) {
				t.Errorf("The path must not exist: %s", test.path)
			}
		})
	}
}

func TestGetFileStem(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		path string
		stem string
	}{
		// Normal
		{"1", testPathJpg, "sample1"},
		{"2", testPathJpeg, "sample2"},
		{"3", testPathPng, "sample3"},
		{"4", testPathGif, "sample4"},
		{"5", testPathTif, "sample5"},
		{"6", testPathTiff, "sample6"},
		{"7", testPathBmp, "sample7"},
		// Abnormal
		{"8", extNon, extNon},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if files.Stem(test.path) != test.stem {
				t.Errorf("The file stem must be %s", test.stem)
			}
		})
	}
}

func TestGetFormattedFileExt(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		path string
		ext  string
	}{
		// Normal
		{"1", testPathJpg, extJpg},
		{"2", testPathJpeg, extJpeg},
		{"3", testPathPng, extPng},
		{"4", testPathGif, extGif},
		{"5", testPathTif, extTif},
		{"6", testPathTiff, extTiff},
		{"7", testPathBmp, extBmp},
		// Abnormal
		{"8", extNon, strEmpty},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if files.Ext(test.path) != test.ext {
				t.Errorf("The file extension must be %s", test.ext)
			}
		})
	}
}

func TestGetMimeType(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		path     string
		mimeType string
	}{
		// .tif and .tiff cannot be detected as image/tiff
		// when using http.DetectContentType in files.GetMimeType()
		{"1", testPathJpg, "image/jpeg"},
		{"2", testPathJpeg, "image/jpeg"},
		{"3", testPathPng, "image/png"},
		{"4", testPathGif, "image/gif"},
		{"5", testPathTif, "application/octet-stream"},
		{"6", testPathTiff, "application/octet-stream"},
		{"7", testPathBmp, "image/bmp"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			mimeType, err := files.MimeType(test.path)
			if err != nil {
				t.Error(err)
			}
			if mimeType != test.mimeType {
				t.Errorf("The file's MIME Type must be %s", mimeType)
			}
		})
	}
}
