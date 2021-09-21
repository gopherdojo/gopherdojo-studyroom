package imgconv

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"kadai1/fileutil"
)

const (
	// Extensions
	extJpg        = ".jpg"
	extJpeg       = ".jpeg"
	extPng        = ".png"
	extGif        = ".gif"
	defaultSrcExt = extJpg
	defaultDstExt = extPng

	// Directories
	defaultSrcDir = "./testdata/src"
	defaultDstDir = "./testdata/dst"

	// Flag
	defaultFileDeleteFlag = false
)

type Converter struct {
	srcExt         string
	dstExt         string
	srcDir         string
	dstDir         string
	fileDeleteFlag bool
}

var (
	// Extensions' list
	extList    []string = []string{extJpg, extJpeg, extPng, extGif}
	extListStr string   = strings.Join(extList, " ")

	// Arguments
	SrcExt         *string = flag.String("src-ext", defaultSrcExt, "Source extension (choices "+extListStr+")")
	DstExt         *string = flag.String("dst-ext", defaultDstExt, "Destination extension (choices "+extListStr+")")
	SrcDir         *string = flag.String("src-dir", defaultSrcDir, "Source directory")
	DstDir         *string = flag.String("dst-dir", defaultDstDir, "Destination directory")
	FileDeleteFlag *bool   = flag.Bool("delete", defaultFileDeleteFlag, "File delete flag")
)

// Validate arguments
func ValidateArgs() error {
	if !containsStringInSlice(extList, *SrcExt) {
		return fmt.Errorf("the selected \"src-ext\" does not exist in: %v", extListStr)
	}

	if !containsStringInSlice(extList, *DstExt) {
		return fmt.Errorf("the selected \"dst-ext\" does not exist in: %v", extListStr)
	}

	dir, err := os.Open(*SrcDir)
	if err != nil {
		return fmt.Errorf("faild to open the directory of the \"src-dir\": %w", err)
	}

	dirInfo, err := dir.Stat()
	if err != nil {
		return fmt.Errorf("failed to get the directory info of the \"src-dir\": %w", err)
	}

	if !dirInfo.IsDir() {
		return errors.New("the \"src-dir\" must be an existing directory")
	}

	if err := dir.Close(); err != nil {
		return fmt.Errorf("failed to close the directory of \"src-dir\": %w", err)
	}

	if getType(*FileDeleteFlag) != "bool" {
		return errors.New("the \"delete\" option must be true or false")
	}
	return nil
}

// Get new converter
func NewConverter(srcExt string, dstExt string, srcDir string, dstDir string, fileDeleteFlag bool) *Converter {
	return &Converter{
		srcExt:         srcExt,
		dstExt:         dstExt,
		srcDir:         srcDir,
		dstDir:         dstDir,
		fileDeleteFlag: fileDeleteFlag,
	}
}

// Run converter
func (conv *Converter) Run() error {
	return filepath.Walk(conv.srcDir, func(srcFilePath string, info os.FileInfo, err error) error {
		// Skip the process, if there is an error when getting os.FileInfo inside filepath.Walk ()
		if err != nil {
			return err
		}
		// Skip the process, if it is a directory
		if info.IsDir() {
			return nil
		}
		// Skip the process if an acquired extension and a source extension are different
		if ext := fileutil.GetFormattedFileExt(srcFilePath); ext != conv.srcExt {
			return nil
		}
		// Get an absolute path of a destination directory
		dstDir, err := conv.getAbsPathDstDir(srcFilePath)
		if err != nil {
			return err
		}
		// Make a destination directory if it does not exist
		// Determine the existence of a destination directory in os.MkdirAll()
		if err = os.MkdirAll(dstDir, 0777); err != nil {
			return err
		}
		// Get a destination file path
		dstFilePath := conv.getDstFilePath(srcFilePath, dstDir)
		// Convert a image
		if err := conv.convert(srcFilePath, dstFilePath); err != nil {
			return err
		}

		if conv.fileDeleteFlag {
			// Delete a source file
			if err := fileutil.DeleteFile(srcFilePath); err != nil {
				return err
			}
		}
		return nil
	})
}

// Get an absolute path of a destination directory
func (conv *Converter) getAbsPathDstDir(srcFilePath string) (string, error) {
	// Get a relative path of a source directory
	relPathSrcDir, err := conv.getRelPathSrcDir(srcFilePath)
	if err != nil {
		return "", err
	}
	// Get an absolute path of a destination directory
	path := filepath.Join(conv.dstDir, relPathSrcDir)
	absPathDstDir, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return absPathDstDir, nil
}

// Get a relative path of a source directory
func (conv *Converter) getRelPathSrcDir(srcFilePath string) (string, error) {
	dir := fileutil.GetDirName(srcFilePath)
	return fileutil.GetRelPath(conv.srcDir, dir)
}

// Get a destination file path
func (conv *Converter) getDstFilePath(srcFilePath string, dstDir string) string {
	dstFileName := fileutil.GetFileStem(srcFilePath) + conv.dstExt
	return filepath.Join(dstDir, dstFileName)
}

// Convert an image file
func (conv *Converter) convert(srcFilePath string, dstFilePath string) error {
	srcImage, err := getSrcImage(srcFilePath)
	if err != nil {
		return err
	}
	dstImage, err := getDstImage(dstFilePath)
	if err != nil {
		return err
	}
	if err := conv.encode(dstImage, srcImage); err != nil {
		return err
	}
	return nil
}

// Encode an image
func (conv *Converter) encode(dstImage io.Writer, srcImage image.Image) error {
	switch conv.dstExt {
	case extPng:
		return png.Encode(dstImage, srcImage)
	case extJpg, extJpeg:
		return jpeg.Encode(dstImage, srcImage, nil)
	case extGif:
		return gif.Encode(dstImage, srcImage, nil)
	default:
		return fmt.Errorf("failed to encode due to unsupported extension: %v", conv.dstExt)
	}
}

// Get a source image
func getSrcImage(srcFilePath string) (image.Image, error) {
	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		return nil, err
	}
	// TODO defer 使用時 かつ Close() 失敗時の error の返却方法をレビュー時に確認
	defer srcFile.Close()
	srcImage, _, err := image.Decode(srcFile)
	if err != nil {
		return nil, err
	}
	return srcImage, nil
}

// Get a destination image
func getDstImage(dstFilePath string) (*os.File, error) {
	return os.Create(dstFilePath)
}

// Check if a slice contains a target string
func containsStringInSlice(s []string, target string) bool {
	for _, v := range s {
		if v == target {
			return true
		}
	}
	return false
}

// Get a type
func getType(i interface{}) string {
	return reflect.TypeOf(i).String()
}
