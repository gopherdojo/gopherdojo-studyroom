package downloader_test

import (
	"github.com/yuonoda/gopherdojo-studyroom/kadai3-2/yuonoda/terminate"
	"os"
	"testing"

	downloader "github.com/yuonoda/gopherdojo-studyroom/kadai3-2/yuonoda/downloader"
)

func TestDownload(t *testing.T) {
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
			ctx := terminate.Listen()
			d := downloader.Downloader{Url: c.url}
			filePath := d.Download(ctx, c.concurrency, dwDirPath)
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
