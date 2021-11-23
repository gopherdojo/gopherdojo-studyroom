package imgconv_test

import (
	"errors"
	"flag"
	"fmt"
	"strings"
	"testing"
	"time"

	ic "github.com/exaream/gopherdojo-studyroom/kadai1/exaream/imgconv"
	"github.com/exaream/gopherdojo-studyroom/kadai1/exaream/slices"
)

const (
	// Extensions
	extJpg  = "jpg"
	extJpeg = "jpeg"
	extPng  = "png"
	extGif  = "gif"
	extTiff = "tiff"
	extTif  = "tif"
	extBmp  = "bmp"
	extTxt  = "txt"

	// Arguments
	argSrcExt         = "src-ext"
	argDstExt         = "dst-ext"
	argSrcDir         = "src-dir"
	argDstDir         = "dst-dir"
	argFileDeleteFlag = "delete"

	// Directories
	testDataDir  = "../testdata"
	srcDir       = testDataDir + "/src"
	dstDir       = testDataDir + "/dst"
	testPathJpg  = testDataDir + "/src/sample1.jpg"
	testPathJpeg = testDataDir + "/src/sample2.jpeg"
	testPathPng  = testDataDir + "/src/sample3.png"
	testPathGif  = testDataDir + "/src/sample4.gif"
	testPathTiff = testDataDir + "/src/sample5.tiff"
	testPathTif  = testDataDir + "/src/sample6.tif"
	testPathBmp  = testDataDir + "/src/sample7.bmp"

	errMsgSrcExtMustBeSelected = "the src-ext must be selected from: jpg jpeg png gif tif tiff bmp"
	errMsgDstExtMustBeSelected = "the dst-ext must be selected from: jpg jpeg png gif tif tiff bmp"
	errMsgSrcDirDoesNotExist   = "the src-dir does not exist:"
	errMsgSrcDirMustBeDir      = "the src-dir must be a directory:"
	errMsgSrcDirMustNotBeEmpty = "the src-dir must not be empty"
	errMstDstDirMustNotBeEmpty = "the dst-dir must not be empty"

	// Other
	strEmpty = ""
	strTrue  = "true"
	strFalse = "false"
)

var randomDir string = testDataDir + "/" + fmt.Sprint(time.Now().UnixNano())

func TestValidateArgsNormal(t *testing.T) {
	tests := []struct {
		name           string
		srcExt         string
		dstExt         string
		srcDir         string
		dstDir         string
		fileDeleteFlag string
	}{
		// -src-ext=.jpg and -src-ext=.*
		{"1", extJpg, extJpg, srcDir, dstDir, strTrue},
		{"2", extJpg, extJpeg, srcDir, dstDir, strTrue},
		{"3", extJpg, extPng, srcDir, dstDir, strTrue},
		{"4", extJpg, extGif, srcDir, dstDir, strTrue},
		// -src-ext=.jpeg and -src-ext=.*
		{"5", extJpeg, extJpg, srcDir, dstDir, strTrue},
		{"6", extJpeg, extJpeg, srcDir, dstDir, strTrue},
		{"7", extJpeg, extPng, srcDir, dstDir, strTrue},
		{"8", extJpeg, extGif, srcDir, dstDir, strTrue},
		// -src-ext=.png and -src-ext=.*
		{"9", extPng, extJpg, srcDir, dstDir, strTrue},
		{"10", extPng, extJpeg, srcDir, dstDir, strTrue},
		{"11", extPng, extPng, srcDir, dstDir, strTrue},
		{"12", extPng, extGif, srcDir, dstDir, strTrue},
		// -src-ext=.gif and -src-ext=.*
		{"13", extGif, extJpg, srcDir, dstDir, strTrue},
		{"14", extGif, extJpeg, srcDir, dstDir, strTrue},
		{"15", extGif, extPng, srcDir, dstDir, strTrue},
		{"16", extGif, extGif, srcDir, dstDir, strTrue},
		// -src-dir=existing directory and -dst-dir=existing directory
		{"17", extJpg, extPng, srcDir, srcDir, strTrue},
		{"18", extJpg, extPng, dstDir, dstDir, strTrue},
		{"19", extJpg, extPng, dstDir, srcDir, strTrue},
		// -src-dir=existing directory and -dst-dir=non-existent directory
		{"20", extJpg, extPng, srcDir, randomDir, strTrue},
		// -delete=false
		{"21", extJpg, extPng, srcDir, dstDir, strFalse},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			setFlagComamandLine(t, argSrcExt, test.srcExt)
			setFlagComamandLine(t, argDstExt, test.dstExt)
			setFlagComamandLine(t, argSrcDir, test.srcDir)
			setFlagComamandLine(t, argDstDir, test.dstDir)
			setFlagComamandLine(t, argFileDeleteFlag, test.fileDeleteFlag)
			flag.Parse()
			if err := ic.ValidateArgs(); err != nil {
				t.Errorf("%v", err)
			}
		})
	}
}

func TestValidateArgsAbnormal(t *testing.T) {
	tests := []struct {
		name           string
		srcExt         string
		dstExt         string
		srcDir         string
		dstDir         string
		fileDeleteFlag string
		errMsg         string
	}{
		// -src-ext
		{"1", extTxt, extPng, srcDir, dstDir, strTrue, errMsgSrcExtMustBeSelected},
		{"2", strEmpty, extPng, srcDir, dstDir, strTrue, errMsgSrcExtMustBeSelected},
		// -dst-ext
		{"3", extJpg, extTxt, srcDir, dstDir, strTrue, errMsgDstExtMustBeSelected},
		{"4", extJpg, strEmpty, srcDir, dstDir, strTrue, errMsgDstExtMustBeSelected},
		// -src-ext and -dst-ext
		{"5", extTxt, extTxt, srcDir, dstDir, strTrue, errMsgSrcExtMustBeSelected},
		{"6", strEmpty, strEmpty, srcDir, dstDir, strTrue, errMsgSrcExtMustBeSelected},
		// -src-dir
		{"7", extJpg, extPng, randomDir, dstDir, strTrue, errMsgSrcDirDoesNotExist},
		{"8", extJpg, extPng, testPathJpg, dstDir, strTrue, errMsgSrcDirMustBeDir},
		{"9", extJpg, extPng, strEmpty, dstDir, strTrue, errMsgSrcDirMustNotBeEmpty},
		// -dst-dir
		{"10", extJpg, extPng, srcDir, strEmpty, strTrue, errMstDstDirMustNotBeEmpty},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			setFlagComamandLine(t, argSrcExt, test.srcExt)
			setFlagComamandLine(t, argDstExt, test.dstExt)
			setFlagComamandLine(t, argSrcDir, test.srcDir)
			setFlagComamandLine(t, argDstDir, test.dstDir)
			flag.Parse()
			switch test.name {
			case "1", "2", "3", "4", "5", "6":
				if err := ic.ValidateArgs(); err.Error() != test.errMsg {
					t.Error(err)
				}
			case "7", "8", "9", "10":
				if err := ic.ValidateArgs(); !strings.Contains(err.Error(), test.errMsg) {
					t.Error(err)
				}
			}
		})
	}
}

func TestNewConverter(t *testing.T) {
	tests := []struct {
		name           string
		srcExt         string
		dstExt         string
		srcDir         string
		dstDir         string
		fileDeleteFlag bool
	}{
		{"1", extJpg, extJpeg, srcDir, dstDir, false},
		{"2", extPng, extJpg, srcDir, srcDir, true},
		{"3", extJpg, extGif, dstDir, dstDir, false},
		{"4", extTif, extJpg, dstDir, srcDir, true},
		{"5", extJpg, extTiff, srcDir, srcDir, false},
		{"6", extBmp, extJpg, dstDir, dstDir, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			conv := ic.NewConverter(test.srcExt, test.dstExt, test.srcDir, test.dstDir, test.fileDeleteFlag)
			if test.srcExt != ic.ExportSrcExt(conv) || test.dstExt != ic.ExportDstExt(conv) ||
				test.srcDir != ic.ExportSrcDir(conv) || test.dstDir != ic.ExportDstDir(conv) ||
				test.fileDeleteFlag != ic.ExportFileDeleteFlag(conv) {
				t.Errorf("the values do not match: %v", *conv)
			}
		})
	}
}

func TestContainsStringInSlice(t *testing.T) {
	t.Parallel()
	slice := []string{"Hello", "Good morning", "Good afternoon"}
	sliceStr := `"` + strings.Join(slice, `" "`) + `"`
	tests := []struct {
		name     string
		value    string
		expected bool
	}{
		{"1", "Hello", true},
		{"2", "Good afternoon", true},
		{"3", "hello", false},
		{"4", "morning", false},
		{"5", "Good evening", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if slices.Contains(slice, test.value) != test.expected {
				t.Errorf(`value=%v, slice={%v}`, test.value, sliceStr)
			}
		})
	}
}

func TestGetType(t *testing.T) {
	t.Parallel()
	type customType int
	sampleCustomType := customType(1)
	sampleArray := [...]int{10, 20, 30, 40, 50}
	sampleSlice := []int{1, 2, 3}
	sampleErr := errors.New("error!")

	tests := []struct {
		name     string
		value    interface{}
		expected interface{}
	}{
		{"bool", true, "bool"},
		{"string", "Hello", "string"},
		{"int", 1, "int"},
		{"float", 1.0, "float64"},
		{"complex", (0 + 0i), "complex128"},
		{"array", sampleArray, "[5]int"},
		{"slice", sampleSlice, "[]int"},
		{"error", sampleErr, "*errors.errorString"},
		{"custom_type", sampleCustomType, "imgconv_test.customType"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if actual := ic.ExportGetType(test.value); actual != test.expected {
				t.Errorf("The %v must be %v type", test.value, test.expected)
			}
		})
	}
}

func setFlagComamandLine(t *testing.T, name string, value string) {
	t.Helper()
	if err := flag.CommandLine.Set(name, value); err != nil {
		t.Fatalf("Faild to set a flag of command line:name=%v, value=%v", name, value)
	}
}
