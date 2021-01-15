package convImages_test

import (
	convImages "github.com/yuonoda/gopherdojo-studyroom/kadai1/yuonoda/lib"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"testing"
)

func TestRun(t *testing.T) {

	cases := []struct {
		name     string
		path     string
		from     string
		to       string
		testImgs []string
		expected []string
	}{
		{"pngToJpg", "test", "png", "jpg", []string{"png.png"}, []string{"png.jpg"}},
		{"jpgToGif", "test", "jpg", "gif", []string{"jpg.jpg", "test2/jpg.jpg"}, []string{"jpg.gif", "test2/jpg.gif"}},
		{"gifToPng", "test", "gif", "png", []string{"png.png"}, []string{"gif.png"}},
	}

	for _, c := range cases {

		t.Run(c.name, func(t *testing.T) {

			// テスト用のディレクトリを作成
			testBuiltTestDir(t, c.testImgs)

			// オプションを指定して実行
			convImages.Run(c.from, c.to, c.path)

			// ファイルが作成されているか確認
			for _, e := range c.expected {
				if _, err := os.Stat("test/" + e); os.IsExist(err) {
					t.Error(err)
				}
			}

			// テストディレクトリの削除
			err := os.RemoveAll("test")
			if err != nil {
				t.Error(err)
			}
		})

	}
}

func testBuiltTestDir(t *testing.T, paths []string) {
	//t.Helper()
	for _, p := range paths {
		// ディレクトリがなければつくる
		p = "test/" + p
		dir, _ := filepath.Split(p)
		buildDirIfNotExist(t, dir)

		// ファイルをつくる
		file, err := os.Create(p)
		if err != nil {
			t.Error(err)
		}

		// 拡張子を判定
		allowedExt := map[string]bool{
			"png": true,
			"jpg": true,
			"gif": true,
		}
		ext := filepath.Ext(p)
		if ext == "" || !allowedExt[ext[1:]] {
			continue
		}
		fmt := ext[1:]

		// データのエンコードと格納
		img := testGenerateImage(t)
		switch fmt {
		case "png":
			err = png.Encode(file, img)
			break
		case "jpg":
			err = jpeg.Encode(file, img, nil)
			break
		case "gif":
			err = gif.Encode(file, img, nil)
			break
		}
		if err != nil {
			t.Error(err)
		}

	}
	return
}

// ディレクトリが存在しなければつくる
func buildDirIfNotExist(t *testing.T, dir string) {
	if dir == "" {
		return
	}
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		err = nil
		err = os.MkdirAll(dir, 0777)
		if err != nil {
			t.Error(err)
		}
	}
}
