package search

import (
	"io/ioutil"
	"log"
	"path/filepath"
)

// Get a list of files under the directory.
func GetFiles(dir string, ext string) []string {
	// Get the files in the target directory
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	var arr []string
	for _, file := range files {
		name := file.Name()
		// If the file is directory, add files recursively.
		if file.IsDir() {
			for _, subFile := range GetFiles(dir+name, ext) {
				arr = append(arr, name+"/"+subFile)
			}
		}
		if filepath.Ext(name) == ext {
			arr = append(arr, name)
		}
	}
	return arr
}
