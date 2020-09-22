package util

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, Dirwalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}

func convertImageDir(targetDir, from, to string) {
	filename := "image1.jpg"
	path := fmt.Sprintf("targetDir/%s", filename)
	convertImage(path, from, to)
}

func convertImage(path, from, to string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, err := jpeg.Decode(f)
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	if err := png.Encode(buf, img); err != nil {
		panic(err)
	}

	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Write(buf.Bytes())
}
