/*
Conversion は 画像の拡張子の変更を行うためのパッケージです。

*/
package conversion

import (
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
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

	output, err := os.Create(dirpath)
	defer output.Close()
	if err != nil {
		return errors.New("output失敗")
	}

	img, _, Err := image.Decode(exFile)
	if Err != nil {
		return errors.New("Decode失敗")
	}

	switch extension {
	case JPEG, JPG:
		err = jpeg.Encode(output, img, nil)
		if err != nil {
			return errors.New("Encode失敗")
		}
		fmt.Println("変換成功")
		return nil
	case GIF:
		err = gif.Encode(output, img, nil)
		if err != nil {
			return errors.New("Encode失敗")
		}
		fmt.Println("変換成功")
		return nil
	case PNG:
		err = png.Encode(output, img)
		if err != nil {
			return errors.New("Encode失敗")
		}
		fmt.Println("変換成功")
		return nil
	}
	return nil
}
