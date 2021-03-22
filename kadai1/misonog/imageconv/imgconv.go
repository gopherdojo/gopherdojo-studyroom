package imageconv

import (
	"fmt"
	"strings"
)

// ImgConvはディレクトリを引数にとり、指定した拡張子の画像ファイルを変換する
func ImgConv(dir string, oldExt string, newExt string) error {
	var fileList Files

	if ok := existDir(dir); !ok {
		return fmt.Errorf("%v: No such file or directory", dir)
	}

	vaildOldExt, err := validateExtArg(oldExt)
	if err != nil {
		return err
	}
	validNewExt, err := validateExtArg(newExt)
	if err != nil {
		return err
	}

	paths := dirWalk(dir)
	fileList = getFiles(paths).filter(vaildOldExt)

	for _, file := range fileList {
		if err := file.convert(validNewExt); err != nil {
			return fmt.Errorf("error: cannot create %v", file.Path)
		}
	}

	return nil
}

// validateExtArgは引数として渡された拡張子が正しいかと"."がなければ付与を行う
func validateExtArg(extArg string) (string, error) {
	var collectExt = []string{PNG, JPG, JPEG, GIF}
	ext := strings.ToLower(extArg)

	if !strings.HasPrefix(ext, ".") {
		ext = "." + ext
	}

	for _, i := range collectExt {
		if ext == i {
			return ext, nil
		}
	}
	err := fmt.Errorf("%v is not a supported extension", extArg)
	return "", err
}
