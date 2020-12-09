package gen_imags

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

func getImagePathes(path string, fmt string) ([]string, error) {
	// ディレクトリ内を再帰的に探索
	var pathes []string
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
			pathes = append(pathes, path)
			return nil
		})

	// エラーがあれば、エラーを返す
	if err != nil {
		return []string{}, err
	}
	return pathes, nil
}

func decodeImg(buf *bytes.Reader, fmt string) (image.Image, error) {
	var img image.Image
	var err error
	switch fmt {
	case "jpg":
		img, err = jpeg.Decode(buf)
		break
	case "png":
		img, err = png.Decode(buf)
		break
	case "gif":
		img, err = gif.Decode(buf)
	default:
		err = errors.New("decode format is incorrect")
	}
	return img, err
}

func encodeImg(buf io.Writer, img image.Image, fmt string) error {
	// 変換
	var err error
	switch fmt {
	case "png":
		if err != png.Encode(buf, img) {
			return err
		}
		break
	case "gif":
		if err != gif.Encode(buf, img, nil) {
			return err
		}
		break
	case "jpg":
		if err != jpeg.Encode(buf, img, nil) {
			return err
		}
		break
	default:
		return errors.New("encode format is incorrect")
		break
	}
	return nil
}

func Do() {
	flag.Parse()

	// 指定フォーマットの画像の一覧を取得
	pathes, err := getImagePathes(*path, *fromFmt)
	if err != nil {
		log.Fatal(err)
	}

	// 画像パスをループさせて一括変換
	for _, path := range pathes {
		log.Printf("converting %s to %s", path, *toFmt)

		// 元画像を読み込み
		imageBytes, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}

		// 画像をデコード
		buffer := bytes.NewReader(imageBytes)
		img, err := decodeImg(buffer, *fromFmt)
		if err != nil {
			log.Fatal(err)
		}

		// 変換
		nameNoExt := strings.TrimSuffix(path, filepath.Ext(path))
		newBuf := new(bytes.Buffer)
		encodeImg(newBuf, img, *toFmt)

		// ファイル出力
		newName := nameNoExt + "." + *toFmt
		ioutil.WriteFile(newName, newBuf.Bytes(), 0644)
	}

}
