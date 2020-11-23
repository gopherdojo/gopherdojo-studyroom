package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Convertor flags パッケージをimgconvにしよう
type Convertor struct {
	src     string
	dst     string
	from    string
	to      string
	isDebug bool
}

var cvt *Convertor

func init() {
	cvt = &Convertor{}
	flag.StringVar(&cvt.src, "in", "", "対象のディレクトリ")
	flag.StringVar(&cvt.dst, "out", "./output", "出力対象のディレクトリ")
	flag.StringVar(&cvt.from, "f", "jpg", "変換前の拡張子")
	flag.StringVar(&cvt.to, "t", "png", "変換後の拡張子")
	flag.BoolVar(&cvt.isDebug, "debug", false, "デバッグログを出したい場合はtrueをセットする")
}

// 終了コード
const (
	ExitCodeOK = 0
	ExitError  = 1
)

func main() {
	flag.Parse()
	cvt.debugf("Flags : %+v\n", cvt)

	// 入力チェック
	if msg, ok := cvt.validate(); !ok {
		fmt.Fprintln(os.Stderr, msg)
		os.Exit(ExitError)
	}

	// 拡張子に'.'がついていない場合、'.'を付与する
	cvt.attachDot()

	// 変換を実行する
	if err := cvt.doConvert(); err != nil {
		fmt.Fprintln(os.Stderr, "[ERROR]"+err.Error())
		os.Exit(ExitError)
	}
}

func (c Convertor) doConvert() error {
	// 再帰的に処理を実行する
	err := filepath.Walk(c.src,
		func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == c.from {

				// 対象のファイルに対して処理を実行する
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

func (c *Convertor) convert(path string) error {
	c.debugf("target file path=%v\n", path)

	// 出力先ディレクトリの作成
	dst := filepath.Dir(filepath.Join(c.dst, path))
	c.debugf("output path=%v\n", dst)
	err := os.MkdirAll(dst, os.ModeDir)
	if err != nil {
		return err
	}

	// ファイルのデコード
	in, err := os.Open(path)
	defer in.Close()
	if err != nil {
		return err
	}
	c.debugf("open input file : %s\n", in.Name())

	img, err := c.decode(in)
	if err != nil {
		return err
	}
	c.debugf("decoded jpg file\n")

	// 出力ファイル名取得
	extlen := len(path) - len(filepath.Ext(path))
	filename := filepath.Base(path[:extlen]) + c.to
	dst = filepath.Join(dst, filename)

	// ファイルを出力先にコピー
	out, err := os.Create(dst)
	defer out.Close()
	if err != nil {
		return err
	}
	c.debugf("created output file : %s\n", out.Name())

	err = c.encode(out, img)
	if err != nil {
		return err
	}
	c.debugf("encoded png file\n")
	return nil
}

func (c Convertor) decode(in *os.File) (image.Image, error) {
	switch c.from {
	case ".png":
		fmt.Println("decode png")
		return png.Decode(in)
	default:
		return jpeg.Decode(in)
	}
}

func (c Convertor) encode(w io.Writer, m image.Image) error {
	switch c.to {
	case ".jpg":
		fmt.Println("encode jpg")
		// options := &jpeg.Options{Quality: 100}
		return jpeg.Encode(w, m, nil)
	default:
		return png.Encode(w, m)
	}
}

func (c *Convertor) attachDot() {
	if !strings.Contains(c.from, ".") {
		c.from = "." + c.from
	}
	if !strings.Contains(c.to, ".") {
		c.to = "." + c.to
	}
	c.debugf("attached dot. c.target=%s, c.ext=%s\n", c.from, c.to)
}

func (c Convertor) validate() (string, bool) {
	if c.src == "" {
		return "-inフラグの入力は必須です", false
	}
	return "", true
}

func (c Convertor) debugf(format string, a ...interface{}) {
	c.debug(func() string {
		return fmt.Sprintf(format, a...)
	})
}

func (c Convertor) debug(msg func() string) {
	if c.isDebug {
		fmt.Print("[Debug]", msg())
	}
}
