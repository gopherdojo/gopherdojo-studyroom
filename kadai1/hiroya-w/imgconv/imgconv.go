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

