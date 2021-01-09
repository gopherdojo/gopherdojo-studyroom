package convImages

import (
	"bytes"
	"errors"
	"flag"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
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

type img struct {
	name  string
	Image image.Image
}

func (i *img) decode(buf *bytes.Reader, imgFmt string) error {
	var err error
	switch imgFmt {
	case "jpg":
		i.Image, err = jpeg.Decode(buf)
		break
	case "png":
		i.Image, err = png.Decode(buf)
		break
	case "gif":
		i.Image, err = gif.Decode(buf)
	default:
		err = errors.New("decode format is incorrect")
	}
	return err
}

func (i *img) encode(buf io.Writer, imgFmt string) error {
	// 変換
	switch imgFmt {
	case "png":
		return png.Encode(buf, i.Image)
	case "gif":
		return gif.Encode(buf, i.Image, nil)
	case "jpg":
		return jpeg.Encode(buf, i.Image, nil)
	default:
		return errors.New("encode format is incorrect")
	}
	return nil
}

// Convert images to specified format recursively.
func Do() {
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
		img := img{nameNoExt, nil}
		buffer := bytes.NewReader(imageBytes)
		err = img.decode(buffer, *fromFmt)
		if err != nil {
			log.Fatal(err)
		}

		// 変換
		newBuf := new(bytes.Buffer)
		err = img.encode(newBuf, *toFmt)
		if err != nil {
			log.Fatal(err)
		}

		// ファイル出力
		newName := img.name + "." + *toFmt
		ioutil.WriteFile(newName, newBuf.Bytes(), 0644)
	}

}
