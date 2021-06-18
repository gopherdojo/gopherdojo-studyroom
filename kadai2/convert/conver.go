package convert

import (
	"fmt"
	"image"
	"image/gif"
	_ "image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
)

//再帰的にファイルを読み込む
func GetAllFile(pathname string, ipt []string) ([]string, error) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("Failed to read dir:", err)
		return ipt, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullDir := pathname + "/" + fi.Name()
			ipt, err = GetAllFile(fullDir, ipt)
			if err != nil {
				fmt.Println("Failed to read dir:", err)
				return ipt, err
			}
		} else {
			fullName := pathname + "/" + fi.Name()
			ipt = append(ipt, fullName)
		}
	}
	return ipt, nil
}

//画像の形式を変換する
func Conv(ipt, opt string) (s string, err error) {
	s1 := "right"
	file, err := os.Open(ipt)
	if err != nil {
		return s1, err
	}
	//assert(err, "Invalid image file path ")
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return s1, err
	}
	//assert(err, "Failed to convert file to image.")

	out, err := os.Create(opt)
	if err != nil {
		return s1, err
	}
	defer out.Close()

	switch filepath.Ext(opt) {
	case ".png":
		err = png.Encode(out, img)
	case ".gif":
		err = gif.Encode(out, img, nil)
	}
	if err != nil {
		return s1, err
	}
	return s1, nil
}

//func assert(err error, msg string) {
//	if err != nil {
//		panic(err.Error() + ":" + msg)
//	}
//}
