package convert

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

var flagToExtNames map[string][]string = map[string][]string{
	"png":  {".png"},
	"jpg":  {".jpg", ".jpeg"},
	"jpeg": {".jpeg", ".jpg"},
}

func contains(slice []string, elem string) bool {
	for _, s := range slice {
		if s == elem {
			return true
		}
	}
	return false
}

type converter struct {
	srcDirPath, dstDirPath, bext, aext string
}

func NewConverter(srcDir, dstDir, bExt, aExt string) *converter {
	return &converter{absPath(srcDir), absPath(dstDir), bExt, aExt}
}

func (c *converter) Do() {
	err := filepath.Walk(c.srcDirPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
				return err
			}

			if info.IsDir() {
				return nil
			}

			if contains(flagToExtNames[c.bext], filepath.Ext(path)) {
				c.convert(path)
			}
			return nil
		})

	if err != nil {
		fmt.Println(err)
	}
}

func (c *converter) convert(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("1err: ", err)
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("2err: ", err)
		return
	}

	newFileName := c.getOutputFileName(path)
	newFileDirName := filepath.Dir(newFileName)
	if err := os.MkdirAll(newFileDirName, 0777); err != nil {
		fmt.Println("3err: ", err)
	}

	newfile, _ := os.Create(newFileName)
	defer newfile.Close()

	switch c.aext {
	case "png":
		png.Encode(newfile, img)
	case "jpg", "jpeg":
		jpeg.Encode(newfile, img, &jpeg.Options{Quality: 75})
	}
}

// 入力パスから出力パス(拡張子あり)を返す
func (c *converter) getOutputFileName(path string) string {
	rel, _ := filepath.Rel(c.srcDirPath, path)
	fNameWithoutExt := removeFileExt(filepath.Join(c.dstDirPath, rel))

	newExt := flagToExtNames[c.aext][0]

	return fNameWithoutExt + newExt
}
