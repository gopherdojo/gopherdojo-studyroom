package image

import (
	"os"
	"path/filepath"
)

func Search(dir string, srcFormat string) ([]string, error) {
	const max int = 1024
	sources := make([]string, 0, max)

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
