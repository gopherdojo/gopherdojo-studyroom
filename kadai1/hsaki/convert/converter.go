package convert

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

// flagToExtNamesでは、コマンドの-a, -bオプションで与えられた画像形式がどの拡張子に対応するのかをmapで関連づけたものです。
var flagToExtNames map[string][]string = map[string][]string{
	"png":  {".png"},
	"jpg":  {".jpg", ".jpeg"},
	"jpeg": {".jpeg", ".jpg"},
}

// []string型の引数slice中に、要素elemがあるかどうかを判定する関数
func contains(slice []string, elem string) bool {
	for _, s := range slice {
		if s == elem {
			return true
		}
	}
	return false
}

type converter struct {
	// srcDirPath, dstDirPath ... それぞれ変換前、変換後の画像を配置するディレクトリの絶対パス
	// bext, aext ... それぞれコマンドオプションで与えられた変換前、変換後の画像形式名
	srcDirPath, dstDirPath, bext, aext string
}

// ユーザーがコマンドのフラグで与えられた値の正当性を検証した上で、
// その値を内部にもつconverter構造体を生成するコンストラクタ
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

// レシーバーcが持つ条件で画像変換を実行するメソッド
func (c *converter) Do() error {
	// -srcで指定したディレクトリ以下の画像に対し
	// 再帰的に画像変換をする
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
			return nil
		})

	if err != nil {
		return err
	}
	return nil
}

// レシーバーcが持つ条件で、pathで与えられた画像ファイルの変換を行うメソッド
func (c *converter) convert(path string) error {
	// srcファイルを開いてimage.Image型にデコードする
	file, err := os.Open(path)
	if err != nil {
		return &ConvError{ErrOpenFile, path}
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return &ConvError{ErrCreateImg, path}
	}

	// 出力先のディレクトリ・ファイルの準備をする
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

	// 指定拡張子で画像をエンコードする
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

// pathで与えられた画像ファイルを画像変換したあとの結果を
// どのパスのファイルに配置するかを得る
//
// (例)
// /User/myname/pic配下の.png画像 → /User/myname/result配下に.jpg変換する場合
//
// 引数path: /User/myname/pic/hoge.png  →  結果返り値: /User/myname/result/hoge.jpg
// 引数path: /User/myname/pic/dir/foo.png  →  結果返り値: /User/myname/result/dir/foo.jpg
func (c *converter) getOutputFileName(path string) (string, error) {
	rel, err := filepath.Rel(c.srcDirPath, path)
	if err != nil {
		return "", err
	}
	fNameWithoutExt := removeFileExt(filepath.Join(c.dstDirPath, rel))

	newExt := flagToExtNames[c.aext][0]

	return fNameWithoutExt + newExt, nil
}
