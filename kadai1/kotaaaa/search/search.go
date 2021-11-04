package search

import (
	"io/ioutil"
	"log"
	"path/filepath"
)

// ディレクトリ以下のファイル一覧を取得します。
func GetFiles(dir string, ext string) []string {
	//対象ディレクトリのファイルを取得する
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	var arr []string
	for _, file := range files {
		name := file.Name()
		// ディレクトリであれば、再帰的にファイル名を追加する
		if file.IsDir() {
			for _, subFile := range GetFiles(dir+name, ext) {
				arr = append(arr, name+"/"+subFile)
			}
		}
		if filepath.Ext(name) == ext {
			arr = append(arr, name)
		}
	}
	return arr
}
