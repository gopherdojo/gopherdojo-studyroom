package cli

import (
	"flag"
	"fmt"

	"io"
	"io/fs"
	"path/filepath"

	"github.com/kynefuk/gopherdojo-studyroom/kadai1/converter"
)

// Exit codes are values that represents an exit code
const (
	ExitOk  = 0
	ExitErr = 1
)

// CLI is a struct of cli
type CLI struct {
	OutStream io.Writer
	ErrStream io.Writer
}

// Run is a main func of CLI
func (c *CLI) Run(args []string) int {
	var (
		targetDir string
		fromExt   string
		toExt     string
	)

	flag.CommandLine.SetOutput(c.ErrStream)
	flag.Usage = func() {
		fmt.Fprintf(c.ErrStream, helpText)
	}

	flag.StringVar(&targetDir, "d", "", "specify target directory")
	flag.StringVar(&fromExt, "f", converter.ExtJPG, "specify \"fromExt\"")
	flag.StringVar(&toExt, "t", converter.ExtPNG, "specify \"toExt\"")

	flag.Parse()

	if ok := converter.IsConvertible(fromExt); !ok {
		fmt.Printf("fromExt format is not convertible: %v\n", fromExt)
		return ExitErr
	}

	if ok := converter.IsConvertible(toExt); !ok {
		fmt.Printf("toExt format is not convertible: %v\n", toExt)
		return ExitErr
	}

	// walking directory to collect target image files
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

	// converting
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
