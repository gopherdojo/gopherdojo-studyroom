package lib

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	Path string
	Ext  string
}

const (
	PNG  = ".png"
	JPG  = ".jpg"
	JPEG = ".jpeg"
	GIF  = ".gif"
)

type Files []File

func existDir(dir string) bool {
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

func getFiles(paths []string) Files {
	var fileList []File
	var fileLists Files

	for _, path := range paths {
		fileList = append(fileList, File{
			Path: path,
			Ext:  filepath.Ext(path),
		})
	}

	for _, file := range fileList {
		fileLists = append(fileLists, file)
	}
	return fileLists
}

func (f Files) filter(ext string) Files {
	var fileList Files

	for _, file := range f {
		if file.cmpExt(ext) {
			fileList = append(fileList, file)
		}
	}
	return fileList
}

func (f File) cmpExt(ext string) bool {
	return f.Ext == ext
}

func (f File) convert(ext string) error {
	file, err := os.Open(f.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	imgFile, _, err := image.Decode(file)
	if err != nil {
		imgFile = nil
		return err
	}

	dstFilePath := f.Path[:strings.LastIndex(f.Path, f.Ext)] + ext
	dstFile, err := os.Create(dstFilePath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	switch ext {
	case PNG:
		if err := png.Encode(dstFile, imgFile); err != nil {
			return nil
		}
	case JPG, JPEG:
		if err := jpeg.Encode(dstFile, imgFile, nil); err != nil {
			return nil
		}
	case GIF:
		if err := gif.Encode(dstFile, imgFile, nil); err != nil {
			return nil
		}
	}
	return nil
}
