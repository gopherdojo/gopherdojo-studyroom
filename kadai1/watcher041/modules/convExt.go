package modules

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func ConvExt() {

	// オプションの初期設定をする
	beforeExt := flag.String("beforeExt", "jpg", "変換前のオプション")
	afterExt := flag.String("afterExt", "png", "変換後のオプション")

	// 画像が存在するパスを入力する
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("変換する画像が存在する相対パスを入力してください >>")
	if !scanner.Scan() {
		fmt.Println("Please input source image file path.")
		return
	}

	// 画像があるパスを指定する。
	inputPath := scanner.Text()
	err := filepath.Walk(inputPath, func(filename string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		ext := filepath.Ext(filename)
		if ext == "."+*beforeExt {

			// 対象の画像をオープンする
			imageFile, err := os.Open(filename)
			assert(err, "Invalid image file path "+filename)
			defer imageFile.Close()

			// ファイルオブジェクトを画像オブジェクトに変換
			imgData, _, err := image.Decode(imageFile)
			assert(err, "Failed to convert file to image.")

			// 変換後の拡張子ごとに使用する関数を変更する
			switch *afterExt {
			case "png":
				filename = strings.Replace(filename, "."+*beforeExt, ".png", 1)
				outputFile, err := os.Create(filename)
				assert(err, "Failed to create destination path.")
				defer outputFile.Close()
				png.Encode(outputFile, imgData)
			case "jpg":
				filename = strings.Replace(filename, "."+*beforeExt, ".jpg", 1)
				outputFile, err := os.Create(filename)
				assert(err, "Failed to create destination path.")
				defer outputFile.Close()
				jpeg.Encode(outputFile, imgData, nil)
			}

		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(*beforeExt)
	fmt.Println(*afterExt)

}

// errorオブジェクトをチェックし、nilの場合例外を送出
func assert(err error, msg string) {
	if err != nil {
		panic(err.Error() + ":" + msg)
	}
}
