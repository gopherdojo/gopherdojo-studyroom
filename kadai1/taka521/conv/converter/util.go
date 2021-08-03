package converter

import (
	"os"
	"path/filepath"

	"github.com/taka521/gopherdojo-studyroom/kadai1/taka521/conv/constant"
)

// convertedPath は指定されたファイルパスと拡張子を元に、変換後のファイルパスを返却します。
//
// Example:
// 	full, dir, name := convertedPath("/tmp/avatar.png", ExtensionGif)
//
// 	fmt.Print(full) // => "/tmp/converted/avatar.gif"
// 	fmt.Print(dir)  // => "/tmp/converted"
//	fmt.Print(name) // => "avatar.gif"
func convertedPath(filePath string, to constant.Extension) (string, string, string) {
	dir, name := filepath.Split(filePath)
	ext := filepath.Ext(name)

	out := filepath.Join(dir, "converted")
	fileName := name[:len(name)-len(ext)] + "." + string(to)
	fullPath := filepath.Join(out, fileName)

	return fullPath, out, fileName
}

func checkDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.Mkdir(path, 0777)
	}
}
