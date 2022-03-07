package convExt

import (
	"bufio"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func ConvExt(beforeExt string, afterExt string) error {

	// 変換前後の拡張子が同じ場合は何もせずに終了する
	if beforeExt == afterExt {
		err := errors.New("this is errors.New sample.")
		return err
	}

	// 画像が存在するパスを入力する
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("変換する画像が存在する相対パスを入力してください >>")
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// 画像があるパスを指定する。
	inputPath := scanner.Text()
	err := filepath.Walk(inputPath, func(filename string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		ext := filepath.Ext(filename)
		if ext == "."+beforeExt {

			// 対象の画像をオープンする
			imageFile, err := os.Open(filename)
			if err != nil {
				return err
			}
			defer imageFile.Close()
			if err := imageFile.Sync(); err != nil {
				return err
			}

			// ファイルオブジェクトを画像オブジェクトに変換
			imgData, _, err := image.Decode(imageFile)
			if err != nil {
				return err
			}

			// 変換後の拡張子ごとに使用する関数を変更する
			switch afterExt {
			case "png":
				filename = strings.Replace(filename, "."+beforeExt, ".png", 1)
				outputFile, err := os.Create(filename)
				if err != nil {
					return err
				}
				defer outputFile.Close()
				if err := outputFile.Sync(); err != nil {
					return err
				}
				err = png.Encode(outputFile, imgData)
			case "jpg":
				filename = strings.Replace(filename, "."+beforeExt, ".jpg", 1)
				outputFile, err := os.Create(filename)
				if err != nil {
					return err
				}
				defer outputFile.Close()
				if err := outputFile.Sync(); err != nil {
					return err
				}
				err = jpeg.Encode(outputFile, imgData, nil)
			}
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	return nil
}
