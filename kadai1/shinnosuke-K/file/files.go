package file

import (
	"image"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type File struct {
	Dir       string // absolute directory path which files exists
	Name      string // File Name
	Extension string
}

var fileList = make([]File, 0)

func GetImgFiles(path string, beforeEx string) []File {

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.IsDir() {
			GetImgFiles(filepath.Join(path, f.Name()), beforeEx)

		} else if filepath.Ext(f.Name())[1:] == beforeEx {
			fileList = append(fileList, File{
				Dir:       path,
				Name:      f.Name(),
				Extension: filepath.Ext(f.Name()),
			})
		}
	}
	return fileList
}

func ExistDir(path string) bool {
	if f, err := os.Stat(path); os.IsNotExist(err) || !f.IsDir() {
		return false
	} else {
		return true
	}
}

func DecodeToImg(dir string, name string) (image.Image, error) {
	file, err := os.Open(filepath.Join(dir, name))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	imgFile, _, err := image.Decode(file)
	return imgFile, err
}
