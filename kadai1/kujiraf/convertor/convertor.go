package imgconv

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

// Convertor is image convertor.
// Src: Source directory. Required value.
// Dst: Destination directory.
// From: Extension before convert.
// To: Extension after converted.
// IsDebug: Convertor debug flag.
type Convertor struct {
	Src     string
	Dst     string
	From    string
	To      string
	IsDebug bool
}

// Validate validates flags
func (c *Convertor) Validate() error {
	c.debugf("Flags : %+v\n", c)

	// -inが指定されているか
	if c.Src == "" {
		return fmt.Errorf("-in is required")
	}

	// -from, -toはサポート対象のものを指定しているか
	attachDotIfNotPresent(&c.From)
	if ok := isSupported(c.From); !ok {
		return fmt.Errorf("-from dose not support %s", c.From)
	}
	attachDotIfNotPresent(&c.To)
	if ok := isSupported(c.To); !ok {
		return fmt.Errorf("-to dose not support %s", c.To)
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

func attachDotIfNotPresent(ext *string) {
	if !strings.Contains(*ext, ".") {
		*ext = "." + *ext
	}
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
func (c Convertor) DoConvert() error {
	// 再帰的に処理を実行する
	err := filepath.Walk(c.Src,
		func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == c.From {

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

func (c *Convertor) convert(inputfile string) error {
	c.debugf("target file path=%v\n", inputfile)

	// 出力先ディレクトリの作成
	dst := filepath.Dir(filepath.Join(c.Dst, inputfile))
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

func (c Convertor) decode(input string) (image.Image, error) {
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

func (c Convertor) encode(output string, m image.Image) error {
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

func (c Convertor) debugf(format string, a ...interface{}) {
	c.debug(func() string {
		return fmt.Sprintf(format, a...)
	})
}

func (c Convertor) debug(msg func() string) {
	if c.IsDebug {
		fmt.Print("[Debug]", msg())
	}
}
