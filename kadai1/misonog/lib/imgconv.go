package lib

import "fmt"

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
