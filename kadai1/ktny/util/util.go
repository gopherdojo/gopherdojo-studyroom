package util

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func DirWalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	paths := make([]string, 0, len(files))
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, DirWalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}
	return paths
}

func ConvertImage(path, from, to string) {
	var img image.Image
	var err error
	var f *os.File
	buf := new(bytes.Buffer)
	newFilePath := strings.Replace(path, from, to, 1)

	f, err = os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	switch from {
	case "jpg", "jpeg":
		img, err = jpeg.Decode(f)
		if err != nil {
			panic(err)
		}
	case "png":
		img, err = png.Decode(f)
		if err != nil {
			panic(err)
		}
	}

	switch to {
	case "jpg", "jpeg":
		options := &jpeg.Options{Quality: 100}
		if err = jpeg.Encode(buf, img, options); err != nil {
			panic(err)
		}
	case "png":
		if err := png.Encode(buf, img); err != nil {
			panic(err)
		}
	}

	file, err := os.Create(newFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Write(buf.Bytes())

	if err := os.Remove(path); err != nil {
		panic(err)
	}
}
