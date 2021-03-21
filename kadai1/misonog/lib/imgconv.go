package lib

import (
	"fmt"
	"strings"
)

func ImgConv(dir string, oldExt string, newExt string) error {
	var fileList Files

	if ok := existDir(dir); !ok {
		return fmt.Errorf("%v: No such file or directory", dir)
	}

	paths := dirWalk(dir)
	fileList = getFiles(paths).filter(oldExt)

	for _, file := range fileList {
		if err := file.convert(newExt); err != nil {
			return fmt.Errorf("error: cannot create %v", file.Path)
		}
	}

	return nil
}

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
