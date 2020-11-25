package converter

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// JPEG jpeg
const (
	JPEG = ".jpeg"
	JPG  = ".jpg"
	PNG  = ".png"
	GIF  = ".gif"
)

// Converter is image convertor.
// Src: Source directory. Required value.
// Dst: Destination directory.
// From: Extension before convert.
// To: Extension after converted.
// IsDebug: Converter debug flag.
type Converter struct {
	Src     string
	Dst     string
	From    string
	To      string
	IsDebug bool
}

// Validate validates flags
func (c *Converter) Validate() error {
	c.debugf("Flags : %+v\n", c)

	// 指定されたディレクトリは存在するか
	if f, err := os.Stat(c.Src); os.IsNotExist(err) || !f.IsDir() {
		return fmt.Errorf("%s directory does not exist", c.Src)
	}

	// -from, -toはサポート対象のものを指定しているか
	tmp := []*string{&c.From, &c.To}
	for _, ext := range tmp {
		if !strings.Contains(*ext, ".") {
			*ext = "." + *ext
		}
		if ok := isSupported(*ext); !ok {
			return fmt.Errorf("%s is not supported", *ext)
		}
	}

	// -fromと-toが同じ値ではないか
	if c.From == c.To {
		return fmt.Errorf("-from and -to are same. -from %s, -to %s", c.From, c.To)
	}
	if (c.From == JPEG || c.From == JPG) && (c.To == JPEG || c.To == JPG) {
		return fmt.Errorf("-from and -to are same. -from %s, -to %s", c.From, c.To)
	}

	return nil
}

func isSupported(ext string) bool {
	switch ext {
	case JPEG, JPG, PNG, GIF:
		return true
	default:
		return false
	}
}

// DoConvert converts image's extension from c.From to c.To.
func (c *Converter) DoConvert() error {

	// 出力先パスを絶対パスに変える
	if !filepath.IsAbs(c.Dst) {
		abs, err := filepath.Abs(c.Dst)
		if err != nil {
			return err
		}
		c.Dst = abs
	}
	c.debugf("output root path : %s\n", c.Dst)

	// 処理対象のディレクトリに移動する。処理が終われば元の場所に戻る。
	prevDir, err := filepath.Abs(".")
	c.debugf("current dir : %s\n", prevDir)
	if err != nil {
		return err
	}
	os.Chdir(c.Src)
	defer os.Chdir(prevDir)

	// 処理対象のディレクトリ名取得
	workDir, err := os.Getwd()
	if err != nil {
		return err
	}
	srcRoot := filepath.Base(workDir)
	c.debugf("src dir name : %s\n", srcRoot)

	// 再帰的に処理を実行する
	err = filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == c.From {
				c.debugf("found. %s\n", path)

				// 対象のファイルに対して処理を実行する
				err := c.convert(path, srcRoot)
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

func (c *Converter) convert(inputfile string, root string) error {
	c.debugf("target file path=%v\n", inputfile)

	// 出力先ディレクトリの作成
	dst := filepath.Dir(filepath.Join(c.Dst, root, inputfile))
	c.debugf("output path=%v\n", dst)
	err := os.MkdirAll(dst, os.ModeDir)
	if err != nil {
		return err
	}

	// ファイルのデコード
	img, err := c.decode(inputfile)
	if err != nil {
		return err
	}

	// 出力ファイル名取得
	extlen := len(inputfile) - len(filepath.Ext(inputfile))
	filename := filepath.Base(inputfile[:extlen]) + c.To
	outputfile := filepath.Join(dst, filename)

	// ファイルを出力先にコピー
	err = c.encode(outputfile, img)
	if err != nil {
		return err
	}

	fmt.Printf("[INFO]conversion complete. converted file from %s to %s\n", inputfile, outputfile)
	return nil
}

func (c Converter) decode(input string) (image.Image, error) {
	in, err := os.Open(input)
	defer in.Close()
	if err != nil {
		return nil, err
	}

	switch c.From {
	case PNG:
		c.debugf("decode %s file %s\n", PNG, input)
		return png.Decode(in)
	case GIF:
		c.debugf("decode %s file %s\n", GIF, input)
		return gif.Decode(in)
	default:
		c.debugf("decode %s file %s\n", JPEG, input)
		return jpeg.Decode(in)
	}
}

func (c Converter) encode(output string, m image.Image) error {
	newfile, err := os.Create(output)
	defer newfile.Close()
	if err != nil {
		return err
	}

	switch c.To {
	case JPG, JPEG:
		c.debugf("encode %s file and output to %s\n", JPEG, newfile.Name())
		options := &jpeg.Options{Quality: 100}
		return jpeg.Encode(newfile, m, options)
	case GIF:
		c.debugf("encode %s file and output to %s\n", GIF, newfile.Name())
		options := &gif.Options{NumColors: 256}
		return gif.Encode(newfile, m, options)
	default:
		c.debugf("encode %s file and output to %s\n", PNG, newfile.Name())
		return png.Encode(newfile, m)
	}
}

func (c Converter) debugf(format string, a ...interface{}) {
	c.debug(func() string {
		return fmt.Sprintf(format, a...)
	})
}

func (c Converter) debug(msg func() string) {
	if c.IsDebug {
		fmt.Print("[Debug]", msg())
	}
}
