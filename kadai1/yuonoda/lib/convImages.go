package convImages

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var fromFmt = flag.String("from", "jpg", "Specify a format of your original image")
var toFmt = flag.String("to", "png", "Specify a format which you want to convert the image to")
var path = flag.String("path", ".", "Path to a directory in which images will be converted recursively")

func getImagePaths(path string, fmt string) ([]string, error) {
	log.Println("getImagePaths(path string, fmt string) ([]string, error) ")
	// ディレクトリ内を再帰的に探索
	var paths []string
	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}

			// 対象の拡張子か査定
			ext := filepath.Ext(path)
			if ext == "" || ext[1:] != fmt { // TODO jpgの表記
				return nil
			}
			paths = append(paths, path)
			return nil
		})

	// エラーがあれば、エラーを返す
	if err != nil {
		return nil, err
	}
	return paths, nil
}

// Convert images to specified format recursively.
func Do() {
	log.Println("Do()")
	flag.Parse()

	// 指定フォーマットの画像の一覧を取得
	paths, err := getImagePaths(*path, *fromFmt)
	if err != nil {
		log.Fatal(err)
	}

	// 画像パスをループさせて一括変換
	for _, path := range paths {
		log.Printf("converting %s to %s", path, *toFmt)

		// 元画像を読み込み
		imageBytes, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}

		// 画像をデコード
		nameNoExt := strings.TrimSuffix(path, filepath.Ext(path))
		c := imgConverter{nameNoExt, nil}
		buffer := bytes.NewReader(imageBytes)
		err = c.decode(buffer, *fromFmt)
		if err != nil {
			log.Fatal(err)
		}

		// 画像のエンコード
		newBuf := new(bytes.Buffer)
		err = c.encode(newBuf, *toFmt)
		if err != nil {
			log.Fatal(err)
		}

		// ファイル出力
		newName := c.name + "." + *toFmt
		ioutil.WriteFile(newName, newBuf.Bytes(), 0644)
	}

}
