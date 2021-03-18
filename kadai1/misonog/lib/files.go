package lib

import (
	"log"
	"os"
	"path/filepath"
)

type File struct {
	Path string
	Ext  string
}

func ExistDir(dir string) bool {
	if f, err := os.Stat(dir); os.IsNotExist(err) || !f.IsDir() {
		return false
	}
	return true
}

func dirWalk(dir string) []string {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirWalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}

func getFileStruct(paths []string) []File {
	var fileList []File

	for _, path := range paths {
		fileList = append(fileList, File{
			Path: path,
			Ext:  filepath.Ext(path),
		})
	}
	return fileList
}
