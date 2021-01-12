package convImages

import (
	"flag"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"testing"
)

func TestDo(t *testing.T) {
	cases := []struct {
		path     string
		from     string
		to       string
		expected []string
	}{
		{"test1", "png", "jpg", []string{"test1/png.jpg"}},
		{"test1", "jpg", "gif", []string{"test2/jpg.gif"}},
		{"test1", "gif", "png", []string{"test1/gif.png"}},
	}

	for _, c := range cases {
		// テスト用のディレクトリを作成
		testBuiltTestDir(t)

		// オプションを指定して実行
		flag.CommandLine.Set("path", c.path)
		flag.CommandLine.Set("from", c.from)
		flag.CommandLine.Set("to", c.to)
		Do()

		// ファイルが作成されているか確認
		for _, e := range c.expected {
			if _, err := os.Stat(e); os.IsExist(err) {
				t.Error(err)
			}
		}

		// テストディレクトリの削除
		err := os.RemoveAll("test1")
		if err != nil {
			t.Error(err)
		}
	}
}

func testBuiltTestDir(t *testing.T) {
	t.Helper()

	// テストディレクトリの作成
	err := os.MkdirAll("test1/test2", 0777)
	if err != nil {
		t.Error(err)
	}

	// 作成する画像のリスト
	testImages := []struct {
		fmt  string
		path string
	}{
		{"png", "test1/png.png"},
		{"jpg", "test1/gif.gif"},
		{"gif", "test1/test2/jpg.jpg"},
	}

	// テスト画像の作成
	testImage := testGenerateImage(t)
	for _, i := range testImages {

		// ファイルの作成
		file, err := os.Create(i.path)
		if err != nil {
			t.Error(err)
		}

		// データのエンコードと格納
		switch i.fmt {
		case "png":
			err = png.Encode(file, testImage)
			break
		case "jpg":
			err = jpeg.Encode(file, testImage, nil)
			break
		case "gif":
			err = gif.Encode(file, testImage, nil)
			break
		}
		if err != nil {
			t.Error(err)
		}
	}
	return

}
