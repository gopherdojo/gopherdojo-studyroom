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

func NewConverter(srcDir, dstDir, bExt, aExt string) (*converter, error) {
	srcDirAbsPath, err := absPath(srcDir)
	if err != nil {
		return nil, &ConvError{ErrSrcDirPath, srcDir}
	}

	dstDirAbsPath, err := absPath(dstDir)
	if err != nil {
		return nil, &ConvError{ErrDstDirPath, dstDir}
	}

	if _, ok := flagToExtNames[bExt]; !ok {
		return nil, &ConvError{ErrExt, bExt}
	}

	if _, ok := flagToExtNames[aExt]; !ok {
		return nil, &ConvError{ErrExt, aExt}
	}

	return &converter{srcDirAbsPath, dstDirAbsPath, bExt, aExt}, nil
}

func (c *converter) Do() error {
	err := filepath.Walk(c.srcDirPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return &ConvError{ErrAccessFile, path}
			}

			if info.IsDir() {
				return nil
			}

			if contains(flagToExtNames[c.bext], filepath.Ext(path)) {
				c.convert(path)
			}
			fmt.Println(path)
			//return errors.New("fooooo!")
			return nil
		})

	if err != nil {
		return err
	}
	return nil
}

func (c *converter) convert(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return &ConvError{ErrOpenFile, path}
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return &ConvError{ErrCreateImg, path}
	}

	newFileName, err := c.getOutputFileName(path)
	if err != nil {
		return &ConvError{ErrOutputPath, path}
	}
	newFileDirName := filepath.Dir(newFileName)
	if err := os.MkdirAll(newFileDirName, 0777); err != nil {
		return &ConvError{ErrDstDirPath, c.dstDirPath}
	}

	newfile, err := os.Create(newFileName)
	if err != nil {
		return &ConvError{ErrOutputFile, newFileName}
	}
	defer newfile.Close()

	switch c.aext {
	case "png":
		err = png.Encode(newfile, img)
		if err != nil {
			return &ConvError{ErrEncodeFile, newFileName}
		}
	case "jpg", "jpeg":
		err = jpeg.Encode(newfile, img, &jpeg.Options{Quality: 75})
		if err != nil {
			return &ConvError{ErrEncodeFile, newFileName}
		}
	}
	return nil
}

// 入力パスから出力パス(拡張子あり)を返す
func (c *converter) getOutputFileName(path string) (string, error) {
	rel, err := filepath.Rel(c.srcDirPath, path)
	if err != nil {
		return "", err
	}
	fNameWithoutExt := removeFileExt(filepath.Join(c.dstDirPath, rel))

	newExt := flagToExtNames[c.aext][0]

	return fNameWithoutExt + newExt, nil
}
