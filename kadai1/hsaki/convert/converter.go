package convert

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

const (
	PNG  = "png"
	JPG  = "jpg"
	JPEG = "jpeg"
)

// flagToExtNamesでは、コマンドの-a, -bオプションで与えられた画像形式がどの拡張子に対応するのかをmapで関連づけたものです。
var flagToExtNames map[string][]string = map[string][]string{
	PNG:  {".png"},
	JPG:  {".jpg", ".jpeg"},
	JPEG: {".jpeg", ".jpg"},
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

type Converter struct {
	// srcDirPath, dstDirPath ... それぞれ変換前、変換後の画像を配置するディレクトリの絶対パス
	// bext, aext ... それぞれコマンドオプションで与えられた変換前、変換後の画像形式名
	srcDirPath string
	dstDirPath string
	bext       string
	aext       string
}

// ユーザーがコマンドのフラグで与えられた値の正当性を検証した上で、
// その値を内部にもつConverter構造体を生成するコンストラクタ
func NewConverter(srcDir, dstDir, bExt, aExt string) (*Converter, error) {
	srcDirAbsPath, err := absPath(srcDir)
	if err != nil {
		return nil, &ConvError{Err: err, Code: InValidSrcDirPath, FilePath: srcDir}
	}

	dstDirAbsPath, err := absPath(dstDir)
	if err != nil {
		return nil, &ConvError{Err: err, Code: InValidDstDirPath, FilePath: dstDir}
	}

	if _, ok := flagToExtNames[bExt]; !ok {
		return nil, &ConvError{Err: ErrExt, Code: InValidExt, FilePath: bExt}
	}

	if _, ok := flagToExtNames[aExt]; !ok {
		return nil, &ConvError{Err: ErrExt, Code: InValidExt, FilePath: aExt}
	}

	return &Converter{srcDirPath: srcDirAbsPath, dstDirPath: dstDirAbsPath, bext: bExt, aext: aExt}, nil
}

// レシーバーcが持つ条件で画像変換を実行するメソッド
func (c *Converter) Do() error {
	// -srcで指定したディレクトリ以下の画像に対し
	// 再帰的に画像変換をする
	err := filepath.Walk(c.srcDirPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return &ConvError{Err: err, Code: FileAccessFail, FilePath: path}
			}

			if info.IsDir() {
				return nil
			}

			if contains(flagToExtNames[c.bext], filepath.Ext(path)) {
				err := c.convert(path)
				if err != nil {
					return err
				}
			}
			return nil
		})

	if err != nil {
		return err
	}
	return nil
}

// レシーバーcが持つ条件で、pathで与えられた画像ファイルの変換を行うメソッド
func (c *Converter) convert(path string) error {
	// srcファイルを開いてimage.Image型にデコードする
	file, err := os.Open(path)
	if err != nil {
		return &ConvError{Err: err, Code: FileOpenFail, FilePath: path}
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return &ConvError{Err: err, Code: ImgCreateFail, FilePath: path}
	}

	// 出力先のディレクトリ・ファイルの準備をする
	newFileName, err := c.getOutputFileName(path)
	if err != nil {
		return &ConvError{Err: err, Code: InValidOutputPath, FilePath: path}
	}
	newFileDirName := filepath.Dir(newFileName)
	if err := os.MkdirAll(newFileDirName, 0777); err != nil {
		return &ConvError{Err: err, Code: InValidDstDirPath, FilePath: c.dstDirPath}
	}

	newfile, err := os.Create(newFileName)
	if err != nil {
		return &ConvError{Err: err, Code: FileOutputFail, FilePath: newFileName}
	}
	defer func() {
		err := newfile.Close()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()

	// 指定拡張子で画像をエンコードする
	switch c.aext {
	case PNG:
		err = png.Encode(newfile, img)
		if err != nil {
			return &ConvError{Err: err, Code: FileEncodeFail, FilePath: newFileName}
		}
	case JPG, JPEG:
		err = jpeg.Encode(newfile, img, &jpeg.Options{Quality: 75})
		if err != nil {
			return &ConvError{Err: err, Code: FileEncodeFail, FilePath: newFileName}
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
func (c *Converter) getOutputFileName(path string) (string, error) {
	rel, err := filepath.Rel(c.srcDirPath, path)
	if err != nil {
		return "", err
	}
	fNameWithoutExt := removeFileExt(filepath.Join(c.dstDirPath, rel))

	newExt := flagToExtNames[c.aext][0]

	return fNameWithoutExt + newExt, nil
}
