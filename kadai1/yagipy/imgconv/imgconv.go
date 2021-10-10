// Package imgconv
/*
Abstract

convert image
 */
package imgconv

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

type ImgConv struct {
	Ext string
	FilePath string
}

func NewImgConv(filePath string, ext string) (*ImgConv, error) {
	return &ImgConv{Ext: ext, FilePath: filePath}, nil
}

func (conv *ImgConv)Do() error {
	file, err := os.Open(conv.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	newFileName := conv.FilePath[:len(conv.FilePath)-len(filepath.Ext(conv.FilePath))+1] + conv.Ext
	newFile, err := os.Create(newFileName)
	if err != nil {
		return err
	}
	defer newFile.Close()

	switch conv.Ext {
	case "jpg", "jpeg":
		err := jpeg.Encode(newFile, img, &jpeg.Options{})
		if err != nil {
			return err
		}
	case "png":
		err := png.Encode(newFile, img)
		if err != nil {
			return err
		}
	case "gif":
		err := gif.Encode(newFile, img, nil)
		if err != nil {
			return err
		}
	default:
		errors.New("not support extension")
	}

	err = os.Remove(conv.FilePath)
	if err != nil {
		return err
	}

	return err
}
