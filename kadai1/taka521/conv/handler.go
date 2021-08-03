package conv

import (
	"fmt"
	"image"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/taka521/gopherdojo-studyroom/kadai1/taka521/conv/constant"
	"github.com/taka521/gopherdojo-studyroom/kadai1/taka521/conv/converter"
)

func Handle(input HandlerInput) error {
	// validation
	if err := input.Validate(); err != nil {
		return fmt.Errorf("%w", err)
	}

	// get file path list
	paths, err := getFIlePaths(input)
	if err != nil {
		return fmt.Errorf("ファイルの一覧取得に失敗しました。: %w", err)
	}

	// convert
	success := 0
	for _, path := range paths {
		err := converter.Convert(path, constant.Extension(input.To))
		if err != nil {
			fmt.Printf("Error: %q の変換に失敗しました。: %v\n", path, err)
			continue
		}
		success++
	}

	fmt.Printf("変換処理が終了しました。[成功: %v 件, 失敗: %v 件]\n", success, len(paths)-success)
	return nil
}

// getFIlePaths は画像ファイルのパスを一覧取得します。
func getFIlePaths(input HandlerInput) ([]string, error) {
	files := make([]string, 0)
	err := filepath.Walk(input.Dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 画像ファイル以外はスキップ
		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return nil
		}
		defer f.Close()

		_, format, err := image.DecodeConfig(f)
		if err != nil {
			return nil
		}

		// 変換対象の拡張子以外はスキップ
		if format != input.From {
			return nil
		}

		files = append(files, path)
		return nil
	})

	return files, err
}
