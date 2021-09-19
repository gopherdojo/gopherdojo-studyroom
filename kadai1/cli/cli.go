package cli

import (
	"flag"
	"fmt"

	"io"
	"io/fs"
	"path/filepath"

	"github.com/kynefuk/gopherdojo-studyroom/kadai1/converter"
)

const (
	ExitOk  = 0
	ExitErr = 1
)

type CLI struct {
	OutStream io.Writer
	ErrStream io.Writer
}

func (c *CLI) Run(args []string) int {
	var (
		targetDir string
		fromExt   string
		toExt     string
	)

	flags := flag.NewFlagSet("image-convert", flag.ContinueOnError)
	flags.SetOutput(c.ErrStream)
	flags.Usage = func() {
		fmt.Fprintf(c.ErrStream, helpText)
	}

	flags.StringVar(&targetDir, "d", "", "specify target directory")
	flags.StringVar(&fromExt, "f", converter.ExtJPG, "specify \"fromExt\"")
	flags.StringVar(&toExt, "t", converter.ExtPNG, "specify \"toExt\"")

	if err := flags.Parse(args[1:]); err != nil {
		fmt.Printf("failed to parse flags: %v", err)
		return ExitErr
	}

	// 指定された拡張子が変換可能かチェック
	if ok := converter.IsConvertible(fromExt); !ok {
		fmt.Printf("fromExt format is not convertible: %v\n", fromExt)
		return ExitErr
	}

	if ok := converter.IsConvertible(toExt); !ok {
		fmt.Printf("toExt format is not convertible: %v\n", toExt)
		return ExitErr
	}

	// 指定されたディレクトリをWalkして、fromExtの拡張子に該当する画像ファイルをスライスに入れる
	var targetFiles []string
	err := filepath.WalkDir(targetDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q, error: %v\n", path, err)
			return err
		}

		if ext := filepath.Ext(path); ext == fromExt {
			targetFiles = append(targetFiles, path)
			return nil
		}
		return nil
	})

	if err != nil {
		fmt.Printf("failed to walk dir: %v\n", err)
		return ExitErr
	}

	// スライスの中身を1つずつ変換する
	for _, f := range targetFiles {
		con := converter.Converter{FromExt: fromExt, ToExt: toExt, TargetFilePath: f}

		if err := con.Convert(); err != nil {
			fmt.Printf("failed to convert: %v\n", err)
			continue
		}
	}
	return ExitOk
}

var helpText = `Usage: image-convert [options...]
image-convert
Options:
-d
	specify target dir in which images will be converted
-f
	specify image ext which convert from
-t
	specify image ext which convert to
`
