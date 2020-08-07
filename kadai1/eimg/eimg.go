// eimg package encodes image
// - mandatory
//   - set root directory
//     - default setting is directory executed this command
//   - execute recursively
// - optional
//   - arguments
//     - `-f`
//       - file extension before executing
//       - default setting is jpg/jpeg
//     - `-t`
//       - file extension after executing
//       - default setting is png

package eimg

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	version = "0.0.1"
	msg     = "eimg v" + version + ", converts file extension\n"
)

var (
	fr = flag.String("f", "jpeg", "file extension before executing")
	to = flag.String("t", "png", "file extension after executing")
)

// Eimg structs
type Eimg struct {
	RootDir string
	FromExt string
	ToExt   string
}

// New makes Eimg instance.
// Call this function if you use this package.
func New() *Eimg {
	return &Eimg{
		RootDir: ".",
		FromExt: "jpeg",
		ToExt:   "png",
	}
}

// Run converts file extension(from -> to).
func (eimg *Eimg) Run() error {
	if err := eimg.SetParameters(); err != nil {
		return err
	}

	if err := eimg.ConvertExtension(); err != nil {
		return err
	}
	return nil
}

// SetParameters sets parameters for execution.
func (eimg *Eimg) SetParameters() error {
	// Init Parameters
	*fr = "jpeg"
	*to = "png"

	// parse information
	flag.Parse()
	args := flag.Args()

	// set information.
	if fr != nil && *fr != "jpeg" {
		eimg.FromExt = *fr
	}
	if to != nil && *to != "png" {
		eimg.ToExt = *to
	} else {
		eimg.ToExt = "png"
	}

	if args[len(args)-1] != "." {
		if _, err := os.Stat(args[len(args)-1]); err != nil {
			return ErrInvalidPath.WithDebug(err.Error())
		}
		eimg.RootDir = args[len(args)-1]
	}

	return nil
}

// ConvertExtension converts extension by using set parameters.
func (eimg *Eimg) ConvertExtension() error {
	filePaths, err := eimg.GetFilePathsRec(eimg.RootDir)
	if err != nil {
		return ErrInvalidPath.WithDebug(err.Error())
	}
	for _, filePath := range filePaths {
		extension := filepath.Ext(filePath)
		fmt.Printf("Extension: %s\n", filepath.Ext(filePath))
		if extension == "" {
			continue
		}

		// filepath.Ext starts with "."
		// e.g.) filepath.Ext(filePath) => .txt
		if extension[1:] == eimg.FromExt {
			err := eimg.EncodeFile(filePath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// EncodeFile encodes file predefined arguments
func (eimg *Eimg) EncodeFile(filePath string) error {
	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return ErrInvalidPath.WithDebug(err.Error())
	}

	// ref: https://www.yunabe.jp/docs/golang_pitfall.html
	defer func() {
		cerr := file.Close()
		if cerr == nil {
			return
		}
	}()

	img, _, err := image.Decode(file)
	if err != nil {
		return ErrFailedConvert.WithDebug(err.Error())
	}

	out, err := os.Create(filePath)
	if err != nil {
		return ErrInvalidFormat.WithDebug(err.Error())
	}
	defer func() {
		cerr := out.Close()
		if cerr == nil {
			return
		}
	}()

	switch eimg.ToExt {
	case "png":
		err := png.Encode(out, img)
		if err != nil {
			return ErrFailedConvert.WithDebug(err.Error()).WithHint("converted format is png.")
		}
	case "jpg", "jpeg":
		err := jpeg.Encode(out, img, nil)
		if err != nil {
			return ErrFailedConvert.WithDebug(err.Error()).WithHint("converted format is jpeg/jpg.")
		}
	case "gif":
		err = gif.Encode(out, img, nil)
		if err != nil {
			return ErrFailedConvert.WithDebug(err.Error()).WithHint("converted format is gif.")
		}
	}

	// just convert the extension
	fileName := filepath.Base(filePath) + filepath.Ext(filePath)
	// fileName must meet len(fileName) > len(eimg.FromExt)
	if len(fileName) <= len(eimg.FromExt) {
		return ErrInvalidPath.WithDebug(err.Error()).WithHint("A file name might be less than extension")
	}

	newFilePath := filePath[:len(filePath)-len(eimg.FromExt)] + eimg.ToExt
	if err := os.Rename(filePath, newFilePath); err != nil {
		return ErrFailedConvert.WithDebug(err.Error())
	}

	return nil
}

// GetFilePathsRec gets file list recursively
func (eimg *Eimg) GetFilePathsRec(filePath string) ([]string, error) {
	resFilePaths := make([]string, 0)

	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		return nil, ErrInvalidPath.WithDebug(err.Error())
	}

	for _, file := range files {
		nextFilePath := filepath.Join(filePath, file.Name())
		if file.IsDir() {
			nextFiles, err := eimg.GetFilePathsRec(nextFilePath)
			if err != nil {
				return nil, ErrInvalidPath.WithDebug(err.Error())
			}
			resFilePaths = append(resFilePaths, nextFiles...)
		} else {
			resFilePaths = append(resFilePaths, nextFilePath)
		}
	}

	return resFilePaths, nil
}
