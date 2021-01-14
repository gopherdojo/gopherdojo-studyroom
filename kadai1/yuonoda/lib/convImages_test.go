package convImages_test

import (
	convImages "github.com/yuonoda/gopherdojo-studyroom/kadai1/yuonoda/lib"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	cases := []struct {
		name     string
		path     string
		from     string
		to       string
		expected []string
	}{
		{"pngToJpg", "test1", "png", "jpg", []string{"test1/png.jpg"}},
		{"jpgToGif", "test1", "jpg", "gif", []string{"test2/jpg.gif"}},
		{"gifToPng", "test1", "gif", "png", []string{"test1/gif.png"}},
	}

	for _, c := range cases {

		t.Run(c.name, func(t *testing.T) {

			// テスト用のディレクトリを作成
			testBuiltTestDir(t)

			// オプションを指定して実行
			convImages.Run(c.from, c.to, c.path)

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
		})

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
	cases := []struct {
		name string
		fmt  string
		path string
	}{
		{"encodePng", "png", "test1/png.png"},
		{"encodeJpg", "jpg", "test1/gif.gif"},
		{"encodeGif", "gif", "test1/test2/jpg.jpg"},
	}

	// テスト画像の作成
	testImage := testGenerateImage(t)
	for _, c := range cases {

		// ファイルの作成
		file, err := os.Create(c.path)
		if err != nil {
			t.Error(err)
		}

		// データのエンコードと格納
		switch c.fmt {
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
