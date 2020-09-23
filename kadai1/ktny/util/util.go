package util

import (
	"bytes"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	ImageExtensions = []string{"jpg", "jpeg", "png", "gif"}
)

func GetExtension(s string) string {
	t := strings.Split(s, ".")
	extension := t[len(t)-1]
	return extension
}

func Contains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

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

	newFilePath := strings.Replace(path, from, to, 1)

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
