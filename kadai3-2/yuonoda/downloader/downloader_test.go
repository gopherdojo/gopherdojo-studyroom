package downloader_test

import (
	"os"
	"testing"

	"github.com/yuonoda/gopherdojo-studyroom/kadai3-2/yuonoda/downloader"
	"github.com/yuonoda/gopherdojo-studyroom/kadai3-2/yuonoda/terminate"
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
			err = d.Download(ctx, c.concurrency, dwDirPath)
			if err != nil {
				t.Error(err)
			}
		})
	}

}
