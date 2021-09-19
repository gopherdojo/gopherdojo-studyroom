package imgconv

import (
	"fmt"
	"os"
	"path/filepath"
)

func Converter(directory, inputType, outputType string) error {
	fileNames, err := getFiles(directory, inputType)
	if err != nil {
		return err
	}

	for _, file := range fileNames {
		fmt.Println(file)
	}
	return nil
}

func getFiles(directory, inputType string) ([]string, error) {
	var fileNames []string

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
			fileNames = append(fileNames, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return fileNames, nil
}

