package search

import (
	"io/ioutil"
	"path/filepath"
)

// Get a list of files under the directory.
func GetFiles(dir string, ext string) ([]string, error) {
	// Get the files in the target directory
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var targetFiles []string
	for _, file := range files {
		name := file.Name()
		// If the file is directory, add files recursively.
		if file.IsDir() {
			filesInDir, err := GetFiles(filepath.Join(dir, name), ext)
			if err != nil {
				return nil, err
			}
			for _, subFile := range filesInDir {
				targetFiles = append(targetFiles, filepath.Join(name, subFile))
			}
		}
		if filepath.Ext(name) == ext {
			targetFiles = append(targetFiles, name)
		}
	}
	return targetFiles, err
}
