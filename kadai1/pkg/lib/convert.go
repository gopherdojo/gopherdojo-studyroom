package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/fs"
	"os"
	"path/filepath"
)

type Convert struct {
	selectedDirectory string
	selectedFileType  string
	convertedFileType string
	stringPath        []string
	stringPathBuff    bytes.Buffer
	args []string

}

func NewConvert(args []string) *Convert {
	return &Convert{
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

var fInit = NewFlg()
func init() {
	flag.StringVar(&fInit.selectedDirectory, "s", "", "ディレクトリを指定")
	flag.StringVar(&fInit.selectedFileType, "f", ".jpg", "変換前のファイルタイプを指定")
	flag.StringVar(&fInit.convertedFileType, "cf", ".png", "変換後のファイルタイプを指定")

}

func (c *Convert) returnFilePath() error {
	err := filepath.Walk(c.selectedDirectory,
		func(paths string, info fs.FileInfo, err error) error {
			if filepath.Ext(paths) == c.selectedFileType {
				c.stringPath = append(c.stringPath, paths)
				c.stringPathBuff.WriteString(paths)
				c.stringPathBuff.WriteString(",")
			}
			return nil
		})
	return err
}

func (c *Convert) replaceExt(filePath, from, to string) string {
	ext := filepath.Ext(filePath)
	if len(from) > 0 && ext != from {
		return filePath
	}
	return filePath[:len(filePath)-len(ext)] + to
}

func makeDir(dirName string) {
	if f, err := os.Stat(dirName); !os.IsExist(err) || f.IsDir() {
		if err = os.Mkdir(dirName, 0777); err == nil {
			os.Chdir("..")
		}
	} else {
		fmt.Println("すでに保存先は存在するのでディレクトリを作成できません")
	}
}
func (c *Convert) convertImage(fn string) error {
	f, err := os.Open(fn)
	err = assert("OS.Open", err)
	defer f.Close()

	p, _ := os.Getwd()
	fmt.Println(p)

	img, _, err := image.Decode(f)
	err = assert("Decode", err)

	makeDir("convert")
	makeDir("convert/jpeg")

	fno := c.replaceExt(fn, ".jpg", ".png")

	fo, err := os.Create(filepath.Base(fno))
	err = assert("OS.Create", err)
	defer fo.Close()

	imageType := filepath.Ext(fn)

	os.Chdir("convert")
	os.Chdir("jpeg")
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
	p, _ = os.Getwd()
	fmt.Println(p)

	return err
}

func NewFlg() *Convert{
	return &Convert{
		selectedDirectory: "",
		selectedFileType:  "",
		convertedFileType: "",
		stringPath:        nil,
		stringPathBuff:    bytes.Buffer{},
	}
}

func (c *Convert) Run() error{
	flag.Parse()
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
