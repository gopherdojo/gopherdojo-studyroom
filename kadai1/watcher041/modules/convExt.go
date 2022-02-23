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
	flag.Parse()

	// 変換前後の拡張子が同じ場合は何もせずに終了する
	if *beforeExt == *afterExt {
		fmt.Println("変換する拡張子が前後で同じものです！別々の拡張子に指定してください…")
		return
	}

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
			if err != nil {
				return err
			}
			defer imageFile.Close()

			// ファイルオブジェクトを画像オブジェクトに変換
			imgData, _, err := image.Decode(imageFile)
			if err != nil {
				return err
			}

			// 変換後の拡張子ごとに使用する関数を変更する
			switch *afterExt {
			case "png":
				filename = strings.Replace(filename, "."+*beforeExt, ".png", 1)
				outputFile, err := os.Create(filename)
				if err != nil {
					return err
				}
				defer outputFile.Close()
				png.Encode(outputFile, imgData)
			case "jpg":
				filename = strings.Replace(filename, "."+*beforeExt, ".jpg", 1)
				outputFile, err := os.Create(filename)
				if err != nil {
					return err
				}
				defer outputFile.Close()
				jpeg.Encode(outputFile, imgData, nil)
			}

		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

}
