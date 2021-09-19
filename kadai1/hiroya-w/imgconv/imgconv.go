package imgconv

import (
	"fmt"
	"os"
	"path/filepath"
)

func Converter(directory, inputType, outputType string) error {
	imgPaths, err := getFiles(directory, inputType)
	if err != nil {
		return err
	}

	for _, path := range imgPaths {
		fmt.Println(path)
	}
	return nil
}

func getFiles(directory, inputType string) ([]string, error) {
	var imgPaths []string

	if f, err := os.Stat(directory); err != nil {
		return nil, err
	} else if !f.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", directory)
	}

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == "."+inputType {
			imgPaths = append(imgPaths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return imgPaths, nil
}

