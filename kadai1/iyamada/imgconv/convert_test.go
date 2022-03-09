package imgconv_test

import (
	"fmt"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/fs"
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

func TestExportConvert(t *testing.T) {
	tests := []struct {
		name    string
		arg1    string
		arg2    string
		wantErr bool
	}{
		// TODO: Add error test cases.
		{"jpg to png", "../testdata/ExportConvert/jpg1.jpg", "../testdata/ExportConvert/jpg1.png", false},
		{"jpg to jpeg", "../testdata/ExportConvert/jpg2.jpg", "../testdata/ExportConvert/jpg2.jpeg", false},
		{"jpg to gif", "../testdata/ExportConvert/jpg3.jpg", "../testdata/ExportConvert/jpg3.gif", false},
		{"jpeg to gif", "../testdata/ExportConvert/jpeg1.jpeg", "../testdata/ExportConvert/jpeg1.png", false},
		{"jpeg to jpg", "../testdata/ExportConvert/jpeg2.jpeg", "../testdata/ExportConvert/jpeg2.jpg", false},
		{"jpeg to gif", "../testdata/ExportConvert/jpeg3.jpeg", "../testdata/ExportConvert/jpeg3.gif", false},
		{"png to jpg", "../testdata/ExportConvert/png1.png", "../testdata/ExportConvert/png1.jpg", false},
		{"png to jpeg", "../testdata/ExportConvert/png2.png", "../testdata/ExportConvert/png2.jpeg", false},
		{"png to gif", "../testdata/ExportConvert/png3.png", "../testdata/ExportConvert/png3.gif", false},
		{"gif to jpg", "../testdata/ExportConvert/gif1.gif", "../testdata/ExportConvert/gif1.jpg", false},
		{"gif to jpeg", "../testdata/ExportConvert/gif2.gif", "../testdata/ExportConvert/gif2.jpeg", false},
		{"gif to png", "../testdata/ExportConvert/gif3.gif", "../testdata/ExportConvert/gif3.png", false},
		{"secret", "../testdata/ExportConvert/.secret.jpg", "../testdata/ExportConvert/.secret.png", false},
		// {"ascii", "../testdata/ExportConvert/txt.jpg", "../testdata/ExportConvert/txt.png", true},
		// {"invalid image info", "../testdata/ExportConvert/png.jpg", "../testdata/ExportConvert/png.png", true},
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
			if strings.HasSuffix(tt.arg2, ".png") {
				if ok, err := isPng(tt.arg2); !ok {
					t.Errorf("ExportConvert() can't convert: error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			} else if strings.HasSuffix(tt.arg2, ".gif") {
				if ok, err := isGif(tt.arg2); !ok {
					t.Errorf("ExportConvert() can't convert: error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			} else if strings.HasSuffix(tt.arg2, ".jpg") || strings.HasSuffix(tt.arg2, ".jpeg") {
				if ok, err := isJpg(tt.arg2); !ok {
					t.Errorf("ExportConvert() can't convert: error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
		})
	}
}

func TestConvert(t *testing.T) {
	tests := []struct {
		name         string
		dirs         []string
		inExt        string
		outExt       string
		wantErr      bool
		expDirStruct []string
		createFiles  []string
	}{
		// TODO: Add test cases.
		{"normal", []string{"../testdata/Convert_/normal"}, "jpg", "png", true,
			[]string{
				"../testdata/Convert_/normal/test1.png",
				"../testdata/Convert_/normal/test2.gif",
				"../testdata/Convert_/normal/test3.jpeg",
				"../testdata/Convert_/normal/test4.jpg",
				"../testdata/Convert_/normal/test4.png"},
			[]string{
				"../testdata/Convert_/normal/test4.png"},
		},
		{"subdir", []string{"../testdata/Convert_/inSubDir"}, "jpg", "png", true,
			[]string{
				"../testdata/Convert_/inSubDir/test1.png",
				"../testdata/Convert_/inSubDir/test2.gif",
				"../testdata/Convert_/inSubDir/test3.jpeg",
				"../testdata/Convert_/inSubDir/test4.jpg",
				"../testdata/Convert_/inSubDir/test4.png",
				"../testdata/Convert_/inSubDir/subdir/test1.png",
				"../testdata/Convert_/inSubDir/subdir/test2.gif",
				"../testdata/Convert_/inSubDir/subdir/test3.jpeg",
				"../testdata/Convert_/inSubDir/subdir/test4.jpg",
				"../testdata/Convert_/inSubDir/subdir/test4.png",
			},
			[]string{
				"../testdata/Convert_/inSubDir/test4.png",
				"../testdata/Convert_/inSubDir/subdir/test4.png",
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := imgconv.Convert(tt.dirs, tt.inExt, tt.outExt); (err != nil) != tt.wantErr {
				t.Errorf("ValidateArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assertDirStruct(t, tt.dirs[0], tt.expDirStruct)
			assertImageFiles(t, tt.createFiles)
		})
	}
}

func assertImageFiles(t *testing.T, files []string) {
	t.Helper()
	for _, file := range files {
		if strings.HasSuffix(file, ".png") {
			if ok, _ := isPng(file); !ok {
				t.Errorf("assertImageFiles : %s is not png", file)
				return
			}
		} else if strings.HasSuffix(file, ".gif") {
			if ok, _ := isGif(file); !ok {
				t.Errorf("assertImageFiles : %s is not gif", file)
				return
			}
		} else if strings.HasSuffix(file, ".jpg") || strings.HasSuffix(file, ".jpeg") {
			if ok, _ := isJpg(file); !ok {
				t.Errorf("assertImageFiles : %s is not jpg", file)
				return
			}
		}
	}
}

func assertDirStruct(t *testing.T, testDir string, expDir []string) {
	t.Helper()
	res := make(map[string]int)
	filepath.WalkDir(testDir, func(path string, info fs.DirEntry, err error) error {
		if _, ok := res[path]; !ok {
			res[path] = 1
		}
		return nil
	})
	for _, file := range expDir {
		if _, ok := res[file]; !ok {
			t.Errorf("assertDirStruct : %s is not created", file)
			return
		}
	}
}
