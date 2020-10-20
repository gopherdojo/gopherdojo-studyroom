/*
Conversion は 画像の拡張子の変更を行うためのパッケージです。

*/
package conversion

import (
	"errors"
	"image"
	"image/jpeg"
	"os"
)

const (
	JPEG = "jpeg"
	JPG  = "jpg"
	GIF  = "gif"
	PNG  = "png"
)

// -e で指定した拡張子が対応しているか、判断します。
func ExtensionCheck(ext string) error {
	switch ext {
	case JPEG, JPG, GIF, PNG:
		return nil
	default:
		return errors.New("指定できない拡張子です" + ext)
	}
}

// -f で指定したファイルが存在するか、判断します。
func FilepathCheck(imagepath string) error {
	switch imagepath {
	case "":
		return errors.New("ファイルの指定がされてません" + imagepath)
	default:
		if f, err := os.Stat(imagepath); os.IsNotExist(err) || f.IsDir() {
			return errors.New("ファイルが存在しません" + imagepath)
		} else {
			return nil
		}
	}
}

func FileExtCheck(imagepath string) error {
	switch imagepath {
	case "." + JPEG, "." + JPG, "." + GIF, "." + PNG:
		return nil
	default:
		return errors.New("指定したファイルが対応していません。：" + imagepath)
	}
}

func FileExtension(extension string, imagepath string, dirpath string) error {
	exFile, err := os.Open(imagepath)
	defer exFile.Close()
	if err != nil {
		return errors.New("os.Create失敗")
	}
	output, err := os.Create(imagepath)
	defer output.Close()
	if err != nil {
		return errors.New("output失敗")
	}
	img, _, err := image.Decode(exFile)
	if err != nil {
		return errors.New("image.Decode失敗")
	}

	switch extension {
	case "." + JPEG:
		err = jpeg.Encode(exFile, img, nil)
		return nil
	}
	return nil
}
