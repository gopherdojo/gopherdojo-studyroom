package lib

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

type Flag struct {
	selectedDirectory string
	selectedFileType  string
	convertedFileType string
	stringPath        []string
	stringPathBuff    bytes.Buffer
}

func assert(msg string, err error) error {
	if err != nil {
		fmt.Println(msg, "error")
		return err
	}
	return err
}
func (flg *Flag) init() {
	flag.StringVar(&flg.selectedDirectory, "s", "", "ディレクトリを指定")
	flag.StringVar(&flg.selectedFileType, "f", ".jpg", "変換前のファイルタイプを指定")
	flag.StringVar(&flg.convertedFileType, "cf", ".png", "変換後のファイルタイプを指定")

}

func (flg *Flag) returnFilePath() error {
	err := filepath.Walk(flg.selectedDirectory,
		func(paths string, info fs.FileInfo, err error) error {
			if filepath.Ext(paths) == flg.selectedFileType {
				flg.stringPath = append(flg.stringPath, paths)
				flg.stringPathBuff.WriteString(paths)
				flg.stringPathBuff.WriteString(",")
			}
			return nil
		})
	return err
}

func (flg *Flag) replaceExt(filePath, from, to string) string {
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
func (flg *Flag) convertImage(fn string) error {
	f, err := os.Open(fn)
	err = assert("OS.Open", err)
	defer f.Close()

	p, _ := os.Getwd()
	fmt.Println(p)

	img, _, err := image.Decode(f)
	err = assert("Decode", err)

	makeDir("convert")
	makeDir("convert/jpeg")

	fno := flg.replaceExt(fn, ".jpg", ".png")

	fo, err := os.Create(filepath.Base(fno))
	err = assert("OS.Create", err)
	defer fo.Close()

	imageType := filepath.Ext(fn)

	os.Chdir("convert")
	os.Chdir("jpeg")
	switch imageType {
	case ".jpeg", ".jpg":
		switch flg.convertedFileType {
		case "gif":
			return gif.Encode(fo, img, nil)
		default:
			fmt.Println(imageType)

			return png.Encode(fo, img)
		}

	case "png":
		switch flg.convertedFileType {
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

func NewFlg() *Flag{
	return &Flag{
		selectedDirectory: "",
		selectedFileType:  "",
		convertedFileType: "",
		stringPath:        nil,
		stringPathBuff:    bytes.Buffer{},
	}
}

func (flg *Flag) Exec() {
	flag.Parse()
	err := flg.returnFilePath()
	err = assert("returnFilePath", err)

	for _, v := range flg.stringPath {
		err := flg.convertImage(v)
		err = assert("convertImage", err)

	}

}
