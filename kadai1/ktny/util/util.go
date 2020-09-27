package util

import (
	"bytes"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type img image.Image

const (
	JPG  = "jpg"
	JPEG = "jpeg"
	PNG  = "png"
	GIF  = "gif"
)

func IsSupportExt(ext string) bool {
	switch ext {
	case JPG, JPEG, PNG, GIF:
		return true
	default:
		return false
	}
}

func DirWalk(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	paths := make([]string, 0, len(files))
	for _, file := range files {
		if file.IsDir() {
			childPaths, err := DirWalk(filepath.Join(dir, file.Name()))
			if err != nil {
				return nil, err
			}
			paths = append(paths, childPaths...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}
	return paths, nil
}

func CanConvertExt(from, path string) bool {
	switch from {
	case JPG, JPEG:
		return strings.HasSuffix(path, toExt(JPG)) || strings.HasSuffix(path, toExt(JPEG))
	case PNG:
		return strings.HasSuffix(path, toExt(PNG))
	case GIF:
		return strings.HasSuffix(path, toExt(GIF))
	default:
		return false
	}
}

func toExt(s string) string {
	return "." + s
}

func ConvertImage(path, from, to string) error {
	var img img
	var err error
	var f *os.File
	buf := new(bytes.Buffer)
	newFilePath := strings.Replace(path, from, to, 1)

	f, err = os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	switch from {
	case JPG, JPEG:
		img, err = jpeg.Decode(f)
		if err != nil {
			return err
		}
	case PNG:
		img, err = png.Decode(f)
		if err != nil {
			return err
		}
	case GIF:
		img, err = gif.Decode(f)
		if err != nil {
			return err
		}
	}

	switch to {
	case JPG, JPEG:
		options := &jpeg.Options{Quality: 100}
		if err = jpeg.Encode(buf, img, options); err != nil {
			return err
		}
	case PNG:
		if err := png.Encode(buf, img); err != nil {
			return err
		}
	case GIF:
		options := &gif.Options{}
		if err := gif.Encode(buf, img, options); err != nil {
			return err
		}
	}

	if err := os.Remove(path); err != nil {
		return err
	}

	file, err := os.Create(newFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	file.Write(buf.Bytes())

	return nil
}
