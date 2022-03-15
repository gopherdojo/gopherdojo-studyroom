package imgconv_test

import (
	"io/fs"
	"path/filepath"
	"testing"

	"example.com/ex01/imgconv"
)

func TestRun(t *testing.T) {
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
				"../testdata/Convert_/normal/test3.png",
				"../testdata/Convert_/normal/test4.png",
			},
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
				"../testdata/Convert_/inSubDir/test3.png",
				"../testdata/Convert_/inSubDir/test4.png",
				"../testdata/Convert_/inSubDir/subdir/test3.png",
				"../testdata/Convert_/inSubDir/subdir/test4.png",
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := imgconv.Run(tt.dirs, tt.inExt, tt.outExt); (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
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
		if filepath.Ext(file) == ".png" {
			if ok, _ := isPng(file); !ok {
				t.Errorf("assertImageFiles : %s is not png", file)
				return
			}
		} else if filepath.Ext(file) == ".gif" {
			if ok, _ := isGif(file); !ok {
				t.Errorf("assertImageFiles : %s is not gif", file)
				return
			}
		} else if filepath.Ext(file) == ".jpg" || filepath.Ext(file) == ".jpeg" {
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
