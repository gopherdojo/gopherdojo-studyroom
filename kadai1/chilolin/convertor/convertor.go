package convertor

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// Do is convert image format from src format to dest format
func Do(dir string, srcFormat string, destFormat string) error {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Fprintln(os.Stderr, "このディレクトリにファイルは存在しません")
			return err
		}

		// JPEGファイルの拡張子を .jpg に統一する
		if filepath.Ext(path) == ".jpeg" {
			idx := strings.Index(path, ".jpeg")
			path = path[:idx] + ".jpg"
		}

		if filepath.Ext(path) == "."+srcFormat {
			// TODO
			// 1. 変換前の指定形式のファイルを開く
			sf, err := os.Open(path)
			if err != nil {
				fmt.Fprintln(os.Stderr, "このファイルは開けません")
				return err
			}
			defer sf.Close()

			// 2. ファイルデータを画像データに変換する
			si, _, err := image.Decode(sf)
			if err != nil {
				fmt.Fprintln(os.Stderr, "変換に失敗しました")
				return err
			}

			// 3. 変換後の指定形式のファイルを作成
			idx := strings.Index(filepath.Base(path), "."+srcFormat)
			df, err := os.Create(filepath.Base(path)[:idx] + "." + destFormat)
			if err != nil {
				fmt.Fprintln(os.Stderr, "アウトプットファイルを開けません")
				return err
			}
			defer df.Close()

			// 4. 出力する
			switch destFormat {
			case "jpg":
				jpeg.Encode(df, si, nil)
			case "png":
				png.Encode(df, si)
			case "gif":
				gif.Encode(df, si, nil)
			}

			fmt.Fprintf(os.Stdin, "%s -> %s\n", path, filepath.Base(path)[:idx]+"."+destFormat)
		}

		return nil
	})

	return err
}
