package conv

import (
	"io/ioutil"
	"log"
	"path/filepath"
)

type File struct {
	Dir       string // absolute directory path which files exists
	Name      string // File Name
	Extension string
}

var fileList = make([]File, 0)

func getFiles(path string) []File {

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.IsDir() {
			getFiles(filepath.Join(path, f.Name()))
		} else {
			fileList = append(fileList, File{
				Dir:       path,
				Name:      f.Name(),
				Extension: filepath.Ext(f.Name())[1:],
			})
		}
	}
	return fileList
}
