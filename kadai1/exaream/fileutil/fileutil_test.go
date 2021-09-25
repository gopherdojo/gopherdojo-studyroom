package fileutil_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"assignment/fileutil"
)

const (
	// Paths
	testDataDir  = "../testdata"
	testPathJpg  = testDataDir + "/src/sample1.jpg"
	testPathJpeg = testDataDir + "/src/sample2.jpeg"
	testPathPng  = testDataDir + "/src/sample3.png"
	testPathGif  = testDataDir + "/src/sample4.gif"
	testPathTif  = testDataDir + "/src/sample5.tif"
	testPathTiff = testDataDir + "/src/sample6.tiff"
	testPathBmp  = testDataDir + "/src/sample7.bmp"

	// Extensions
	extJpg  = ".jpg"
	extJpeg = ".jpeg"
	extPng  = ".png"
	extGif  = ".gif"
	extTif  = ".tif"
	extTiff = ".tiff"
	extBmp  = ".bmp"
	extNon  = "ext-non"

	// Other
	emptyStr = ""
)

var randomDir string = testDataDir + "/" + fmt.Sprint(time.Now().UnixNano())

func TestOpenFileNormal(t *testing.T) {
	testCases := []struct {
		name string
		path string
	}{
		{"1", testPathJpg},
		{"2", testPathJpeg},
		{"3", testPathPng},
		{"4", testPathGif},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			if _, err := fileutil.OpenFile(testCase.path); err != nil {
				t.Errorf("failed to open %v", testCase.path)
			}
		})
	}
}

func TestOpenFileAbnormal(t *testing.T) {
	testCases := []struct {
		name string
		path string
	}{
		{"1", randomDir},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			if _, err := fileutil.OpenFile(testCase.path); os.IsExist(err) {
				t.Errorf("The path must not exist: %s", testCase.path)
			}
		})
	}
}

func TestGetFileStem(t *testing.T) {
	testCases := []struct {
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

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			if fileutil.GetFileStem(testCase.path) != testCase.stem {
				t.Errorf("The file stem must be %s", testCase.stem)
			}
		})
	}
}

func TestGetFormattedFileExt(t *testing.T) {
	testCases := []struct {
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
		{"8", extNon, emptyStr},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			if fileutil.GetFormattedFileExt(testCase.path) != testCase.ext {
				t.Errorf("The file extension must be %s", testCase.ext)
			}
		})
	}
}

func TestFormatFileExt(t *testing.T) {
	testCases := []struct {
		name     string
		value    string
		expected string
	}{
		// Normal
		{"1", extJpg, extJpg},
		{"2", extJpeg, extJpeg},
		{"3", extPng, extPng},
		{"4", extGif, extGif},
		{"5", extTif, extTif},
		{"6", extTiff, extTiff},
		{"7", extBmp, extBmp},
		// Normal(uppercase)
		{"8", ".JPG", extJpg},
		{"9", ".JPEG", extJpeg},
		{"10", ".PNG", extPng},
		{"11", ".GIF", extGif},
		{"12", ".TIF", extTif},
		{"13", ".TIFF", extTiff},
		{"14", ".BMP", extBmp},
		// Abnormal
		{"15", ".JpG", extJpg},
		{"16", ".jPEg", extJpeg},
		{"17", ".Png", extPng},
		{"18", ".giF", extGif},
		{"19", ".tIF", extTif},
		{"20", ".TiFf", extTiff},
		{"21", ".bMp", extBmp},
		{"22", emptyStr, emptyStr},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			if fileutil.FormatFileExt(testCase.value) != testCase.expected {
				t.Errorf("The file extension must convert from %s to %s", testCase.value, testCase.expected)
			}
		})
	}
}

func TestGetMimeType(t *testing.T) {
	testCases := []struct {
		name     string
		path     string
		mimeType string
	}{
		// .tif and .tiff cannot be detected as image/tiff
		// when using http.DetectContentType in fileutil.GetMimeType()
		{"1", testPathJpg, "image/jpeg"},
		{"2", testPathJpeg, "image/jpeg"},
		{"3", testPathPng, "image/png"},
		{"4", testPathGif, "image/gif"},
		{"5", testPathTif, "application/octet-stream"},
		{"6", testPathTiff, "application/octet-stream"},
		{"7", testPathBmp, "image/bmp"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			mimeType, err := fileutil.GetMimeType(testCase.path)
			if err != nil {
				t.Error(err)
			}
			if mimeType != testCase.mimeType {
				t.Errorf("The file's MIME Type must be %s", mimeType)
			}
		})
	}
}
