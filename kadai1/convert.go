package main

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/fs"
	"os"
	"path/filepath"
)

func assert(msg string, err error) error {
	if err != nil {
		fmt.Println(msg, "error")
		return err
	}
	return err
}

func (c *Arguments) returnFilePath() error {
	err := filepath.Walk(c.selectedDirectory,
		func(paths string, info fs.FileInfo, err error) error {
			if filepath.Ext(paths) == c.selectedFileType {
				c.stringPath = append(c.stringPath, paths)
			}
			return nil
		})
	return err
}

func (c *Arguments) replaceExt(filePath, from, to string) string {
	ext := filepath.Ext(filePath)
	if len(from) > 0 && ext != from {
		return filePath
	}
	return filePath[:len(filePath)-len(ext)] + to
}

func (c *Arguments) caseConvert(from string, to string, fn string, img image.Image) error {
	fno := c.replaceExt(fn, from, "-out"+to)
	fo, err := os.Create(fno)
	defer fo.Close()
	err = assert("OS.Create", err)

	switch to {
	case ".gif":
		return gif.Encode(fo, img, nil)

	case ".png":

		return png.Encode(fo, img)
	case ".jpg":
		return jpeg.Encode(fo, img, nil)
	}

	return err
}

func (c *Arguments) convertImage(fn string) error {
	f, err := os.Open(fn)
	err = assert("OS.Open", err)
	defer f.Close()

	img, _, err := image.Decode(f)
	err = assert("Decode", err)

	imageType := filepath.Ext(fn)

	switch imageType {

	case ".jpg":
		switch c.convertedFileType {
		case ".gif":
			err = c.caseConvert(".jpg",".gif",fn,img)
		default:
			err = c.caseConvert(".jpg",".png",fn,img)
		}

	case ".jpeg":
		switch c.convertedFileType {
		case ".gif":
			err = c.caseConvert(".jpg",".gif",fn,img)
		default:
			err = c.caseConvert(".jpg",".png",fn,img)
		}

	case ".png":
		switch c.convertedFileType {
		case ".jpeg":
			err = c.caseConvert(".png",".jpeg",fn,img)
		case ".jpg":
			err = c.caseConvert(".png",".jpg",fn,img)
		case ".gif":
			err = c.caseConvert(".png",".gif",fn,img)
		}
	}
	return err
}

func (c *Arguments) Run() error {
	err := c.returnFilePath()
	err = assert("returnFilePath", err)
	var cmd []string

	cmd = append(cmd, c.args...)

	for _, v := range c.stringPath {
		err := c.convertImage(v)
		err = assert("convertImage", err)
	}
	return nil
}
