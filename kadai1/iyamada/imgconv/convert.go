package imgconv

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
)

type myImage image.Image

func encodeImage(file io.Writer, img myImage, ext string) (err error) {
	switch ext {
	case ".jpg", ".jpeg":
		err = jpeg.Encode(file, img, nil)
	case ".png":
		err = png.Encode(file, img)
	case ".gif":
		err = gif.Encode(file, img, nil)
	default:
		return fmt.Errorf("error: %s is cannot encode", ext)
	}
	return err
}

func writeImage(path string, img myImage) (err error) {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	if err := encodeImage(f, img, filepath.Ext(path)); err != nil {
		return err
	}
	defer func() {
		err = f.Close()
	}()
	return nil
}

func decodeImage(file io.Reader, ext string) (img myImage, err error) {
	switch ext {
	case ".jpg", ".jpeg":
		img, err = jpeg.Decode(file)
	case ".png":
		img, err = png.Decode(file)
	case ".gif":
		img, err = gif.Decode(file)
	default:
		return nil, fmt.Errorf("error: %s is cannot decode", ext)
	}
	return img, err
}

func readImage(path string) (img myImage, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	img, err = decodeImage(f, filepath.Ext(path))
	defer func() {
		err = f.Close()
	}()
	return img, err
}

func convert(inPath string, outExt string) (err error) {
	inImg, err := readImage(inPath)
	if err != nil {
		return err
	}
	outPath := genOutPath(inPath, outExt)
	if err := writeImage(outPath, inImg); err != nil {
		return err
	}
	return nil
}
