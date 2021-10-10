package cvt

import (
	"github.com/gopherdojo/gopherdojo-studyroom/kadai1/pkg/testutil"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"testing"
)

// TestConvert:
func TestConvert(t *testing.T) {
	rootDir := getRootDir()
	vectors := map[string]struct {
		outputDir string
		beforeExt string
		afterExt  string
		srcPaths  []string
		expected  []string
		wantErr   error
	}{
		"OK": {
			beforeExt: ".jpg",
			afterExt:  ".png",
			srcPaths: []string{
				filepath.Join(rootDir, "pkg/testdata/fixture/1000.jpg"),
			},
			expected: []string{
				filepath.Join(rootDir, "pkg/testdata/fixture/1000.png"),
			},
		},
		"CASE_SPECIFIED_OUTPUT_DIR_OK": {
			outputDir: "pkg/testdata/secondfixture",
			beforeExt: ".jpg",
			afterExt:  ".png",
			srcPaths: []string{
				filepath.Join(rootDir, "pkg/testdata/fixture/1000.jpg"),
			},
			expected: []string{
				filepath.Join(rootDir, "pkg/testdata/secondfixture/1000.png"),
			},
		},
	}
	for k, v := range vectors {
		imageCvtGlue := NewImageCvtGlue("", v.outputDir, v.beforeExt, v.afterExt, false)
		err := imageCvtGlue.convert(v.srcPaths)
		if errors.Cause(err) != v.wantErr {
			t.Errorf("test %s, convert() = %v, want %v", k, errors.Cause(err), v.wantErr)
		}
		// 出力先が指定されている際は生成されたディレクトリごと削除
		if v.outputDir != "" {
			testutil.RemoveAllTestSrc(filepath.Join(rootDir, v.outputDir))
		} else {
			// 生成されたテストファイルを削除
			for _, expect := range v.expected {
				testutil.RemoveTestFile(expect)
			}
		}
	}
}

// TestPathWalk:
func TestPathWalk(t *testing.T) {
	var currentDir string
	var err error
	if currentDir, err = os.Getwd(); err != nil {
		t.Errorf("failed get current dir: %v", err)
	}
	vectors := map[string]struct {
		inputDir  string
		beforeExt string
		srcPaths  []string
		expected  []string
		wantErr   error
	}{
		"OK": {
			inputDir:  "internal/cvt/walktest",
			beforeExt: ".jpg",
			srcPaths: []string{
				filepath.Join(currentDir, "walktest/test001.jpg"),
				filepath.Join(currentDir, "walktest/test002.jpg"),
				filepath.Join(currentDir, "walktest/test003.png"),
			},
			expected: []string{
				filepath.Join(currentDir, "walktest/test001.jpg"),
				filepath.Join(currentDir, "walktest/test002.jpg"),
			},
		},
	}
	for k, v := range vectors {
		testutil.PrepareTestSrc("walktest", v.srcPaths)
		imageCvtGlue := NewImageCvtGlue(v.inputDir, "", v.beforeExt, "", false)
		actual, err := imageCvtGlue.pathWalk()
		if errors.Cause(err) != v.wantErr {
			t.Errorf("test %s, pathWalk() = %v, want %v", k, errors.Cause(err), v.wantErr)
		}
		// 実際に取得した配列長とテストケースの配列長が異なる場合
		if len(v.expected) != len(actual) {
			t.Errorf("the length of arrays are different test: %s length of expected %d length of actual %d", k, len(v.expected), len(actual))
		}
		for idx, expected := range v.expected {
			if expected != actual[idx] {
				t.Errorf("test: %s expected %s actual %s", k, expected, actual[idx])
			}
		}
		testutil.RemoveAllTestSrc(filepath.Join(currentDir, "walktest"))
	}
}

// TestGetDstPath:
func TestGetDstPath(t *testing.T) {
	var currentDir string
	var err error
	if currentDir, err = os.Getwd(); err != nil {
		t.Errorf("failed get current dir: %v", err)
	}
	vectors := map[string]struct {
		outputDir string
		afterExt  string
		filePath  string
		expected  string
		wantErr   error
	}{
		"CASE_NOT_SPECIFIED_OUTPUT_DIR_OK": {
			outputDir: "",
			afterExt:  ".png",
			filePath:  filepath.Join(currentDir, "walktest/test003.jpg"),
			expected:  filepath.Join(currentDir, "walktest/test003.png"),
		},
		"CASE_SPECIFIED_OUTPUT_DIR_OK": {
			outputDir: "internal/cvt/testdir",
			afterExt:  ".png",
			filePath:  filepath.Join(currentDir, "walktest/test003.jpg"),
			expected:  filepath.Join(currentDir, "testdir/test003.png"),
		},
	}
	for k, v := range vectors {
		imageCvtGlue := NewImageCvtGlue("", v.outputDir, "", v.afterExt, false)

		actual, err := imageCvtGlue.getDstPath(v.filePath)
		if errors.Cause(err) != v.wantErr {
			t.Errorf("test %s, getDstPath() = %v, want %v", k, errors.Cause(err), v.wantErr)
		}
		if actual != v.expected {
			t.Errorf("test: %s expected %s actual %s", k, v.expected, actual)
		}
		if v.outputDir != "" {
			testutil.RemoveAllTestSrc(filepath.Join(currentDir, "walktest"))
		}
	}
}
