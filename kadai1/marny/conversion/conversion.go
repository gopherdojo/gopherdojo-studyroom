/*
Conversion は 画像の拡張子の変更を行うためのパッケージです。

*/
package conversion

import (
	"errors"
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
		return errors.New("指定できない拡張子です")
	}
}

// -f で指定したファイルが存在するか、判断します。
func FilepathCheck(filepath string) error {
	switch filepath {
	case "":
		return errors.New("ファイルの指定がされてません")
	default:
		if f, err := os.Stat(filepath); os.IsNotExist(err) || f.IsDir() {
			return errors.New("ファイルが存在しません")
		} else {
			return nil
		}
	}
}
