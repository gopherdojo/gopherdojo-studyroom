package downloader_test

import (
	"os"
	"reflect"
	"testing"

	downloader "github.com/yuonoda/gopherdojo-studyroom/kadai3-2/yuonoda/downloader"
)

func TestRun(t *testing.T) {
	cases := []struct {
		name         string
		url          string
		expectedSize int64
		concurrency  int
	}{
		{
			name:         "basic",
			url:          "https://dumps.wikimedia.org/jawiki/20210101/jawiki-20210101-pages-articles-multistream-index.txt.bz2",
			expectedSize: int64(25802009),
			concurrency:  3,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// ダウンロードパスを指定
			homedir, err := os.UserHomeDir()
			if err != nil {
				t.Error(err)
			}
			dwDirPath := homedir + "/Downloads"

			// ダウンロード
			filePath := downloader.Run(c.url, c.concurrency, dwDirPath)
			file, err := os.Open(filePath)
			if err != nil {
				t.Error(err)
			}

			// サイズを取得
			info, err := file.Stat()
			if err != nil {
				t.Error(err)
			}
			dwSize := info.Size()

			// サイズを比較
			if dwSize != c.expectedSize {
				t.Errorf("Size doesn't match, got %d but expexted %d", dwSize, c.expectedSize)
			}

			// ファイルを削除
			err = os.Remove(filePath)
			if err != nil {
				t.Error(err)
			}
		})
	}

}

func TestFillByteArr(t *testing.T) {
	cases := []struct {
		name        string
		arr         []byte
		startAt     int
		partArr     []byte
		expectedArr []byte
	}{
		{
			name:        "basic",
			arr:         []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			startAt:     3,
			partArr:     []byte{4, 5, 6},
			expectedArr: []byte{0, 0, 0, 4, 5, 6, 0, 0, 0, 0},
		},
		{
			name:        "basic2",
			arr:         []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			startAt:     8,
			partArr:     []byte{9, 10},
			expectedArr: []byte{0, 0, 0, 0, 0, 0, 0, 0, 9, 10},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			downloader.ExportedFillByteArr(c.arr[:], c.startAt, c.partArr)
			if !reflect.DeepEqual(c.expectedArr, c.arr) {
				t.Error("Array does not match")
			}
		})
	}

}
