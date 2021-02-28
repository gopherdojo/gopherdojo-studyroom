package image

import (
	"os"
	"path/filepath"
)

func Search(dir string, srcFormat string) ([]string, error) {
	sources := make([]string, 0)

	if err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == "."+srcFormat {
				sources = append(sources, path)
			}
			return nil
		},
	); err != nil {
		return nil, err
	}

	return sources, nil
}
