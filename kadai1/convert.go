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


func newConvert(args []string) *Arguments {
	return &Arguments{
		args: args,
	}
}

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
				fmt.Println(paths)
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


func (c *Arguments) convertImage(fn string) error {
	f, err := os.Open(fn)
	err = assert("OS.Open", err)
	defer f.Close()

	img, _, err := image.Decode(f)
	err = assert("Decode", err)

	fno := c.replaceExt(fn, ".jpg", ".png")

	fo, err := os.Create(fno)
	err = assert("OS.Create", err)
	defer fo.Close()

	imageType := filepath.Ext(fn)

	switch imageType {
	case ".jpeg", ".jpg":
		switch c.convertedFileType {
		case "gif":
			return gif.Encode(fo, img, nil)
		default:
			fmt.Println(imageType)

			return png.Encode(fo, img)
		}

	case "png":
		switch c.convertedFileType {
		case "jpeg", "jpg":
			return jpeg.Encode(fo, img, nil)
		case "gif":
			return gif.Encode(fo, img, nil)
		default:
			fmt.Println("default")
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
