//指定したファイルを探索する
package findFile
import (
	"io/ioutil"
	"path/filepath"
)
var format string

func Search(path string,input_fmt string) []string {
	//指定した拡張子のファイルのパスを再帰的に検索
    format = "." + input_fmt
	fileList := dirwalk(path)
	return fileList
}

func dirwalk(dir string) []string {
    files, err := ioutil.ReadDir(dir)//ディレクトリ読み込み
    if err != nil {
        //再帰関数でerrorをmain関数まで返す方法わからない
        panic(err)
    }

    var paths []string
    for _, file := range files {
        if file.IsDir() {
            paths = append(paths, dirwalk(filepath.Join(dir, file.Name()))...)
            
            continue
        }
		    if filepath.Ext(file.Name()) == format {
                paths = append(paths, filepath.Join(dir, file.Name()))
            }
			
		
        
    }

    return paths
}