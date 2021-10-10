// Package file_searcher
/*
Abstract

Search files recursively
 */
package file_searcher

import (
	"io/ioutil"
	"path/filepath"
)

type FileSearcher struct {
	Dir string
	Ext string
}

func NewFileSearcher(ext string, dir string) (*FileSearcher, error) {
	return &FileSearcher{Dir: dir, Ext: ext}, nil
}

func (searcher *FileSearcher) Do() ([]string, error) {
	return do(searcher.Dir, searcher.Ext)
}

func do(dir string, extension string) ([]string, error) {
	var paths []string
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			subpaths, err := do(filepath.Join(dir, file.Name()), extension)
			paths = append(paths, subpaths...)
			if err != nil {
				return nil, err
			}
			continue
		}
		ext := filepath.Ext(file.Name())

		if ext != "" && ext[1:] == extension {
			fullpath := filepath.Join(dir, file.Name())
			paths = append(paths, fullpath)
		}
	}
	return paths, err
}
