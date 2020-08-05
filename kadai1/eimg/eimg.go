// eimg package encode image
// - mandatory
//   - set root directory
//     - default setting is directory executed this command
//   - execute recursively
// - optional
//   - arguments
//     - `-f`
//       - file extension before executing
//       - default setting is jpg
//     - `-t`
//       - file extension after executing
//       - default setting is png

package eimg

import (
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	version = "0.0.1"
	msg     = "eimg v" + version + ", converts file extension\n"
)

// Eimg structs
type Eimg struct {
	RootDir string
	From    string
	To      string
}

// New for eimg package
func New() *Eimg {
	return &Eimg{
		RootDir: ".",
		From:    "jpg",
		To:      "png",
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
	// parse information
	fr := flag.String("f", "jpg", "file extension before executing")
	to := flag.String("t", "png", "file extension after executing")

	flag.Parse()
	args := flag.Args()

	// set information.
	if *fr != "jpg" {
		eimg.From = *fr
	}
	if *to != "png" {
		eimg.To = *to
	}

	// use default setting.
	if len(args) == 0 {
		return nil
	}

	if args[0] != "." {
		if _, err := os.Stat(args[0]); err != nil {
			return ErrInvalidPath.WithDebug(err.Error())
		}
		eimg.RootDir = args[0]
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
		if extension == "" {
			continue
		}

		// # TODO
		// # impelemnt encoding
		// # ref: https://rennnosukesann.hatenablog.com/entry/2019/08/14/175308
		if extension == eimg.From {
            
		}
	}
	return nil
}

// GetFilePathsRec gets file list recursively
func (eimg *Eimg) GetFilePathsRec(filePath string) ([]string, error) {
	// folder has likely more than 5 files...?
	resFilePaths := make([]string, 5)

	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		return nil, ErrInvalidPath.WithDebug(err.Error())
	}

	for _, file := range files {
		nextFilePath := filepath.Join(eimg.RootDir, file.Name())
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
