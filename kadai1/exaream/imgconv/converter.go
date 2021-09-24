// Package imgconv provides functions to convert images.
package imgconv

import (
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

	"assignment/fileutil"

	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
)

const (
	// Extensions
	extJpg        = ".jpg"
	extJpeg       = ".jpeg"
	extPng        = ".png"
	extGif        = ".gif"
	extTif        = ".tif"
	extTiff       = ".tiff"
	extBmp        = ".bmp"
	defaultSrcExt = extJpg
	defaultDstExt = extPng

	// Arguments
	argSrcExt         = "src-ext"
	argDstExt         = "dst-ext"
	argSrcDir         = "src-dir"
	argDstDir         = "dst-dir"
	argFileDeleteFlag = "delete"

	// Directories
	defaultSrcDir = "./testdata/src"
	defaultDstDir = "./testdata/dst"

	// Flag
	defaultFileDeleteFlag = false
)

// Image converter
type Converter struct {
	srcExt         string
	dstExt         string
	srcDir         string
	dstDir         string
	fileDeleteFlag bool
}

var (
	// Supported extensions
	extList    []string = []string{extJpg, extJpeg, extPng, extGif, extTif, extTiff, extBmp}
	extListStr string   = strings.Join(extList, " ")

	// Arguments
	SrcExt         *string = flag.String(argSrcExt, defaultSrcExt, "Source extension (choices "+extListStr+")")
	DstExt         *string = flag.String(argDstExt, defaultDstExt, "Destination extension (choices "+extListStr+")")
	SrcDir         *string = flag.String(argSrcDir, defaultSrcDir, "Source directory")
	DstDir         *string = flag.String(argDstDir, defaultDstDir, "Destination directory")
	FileDeleteFlag *bool   = flag.Bool(argFileDeleteFlag, defaultFileDeleteFlag, "File delete flag")
)

// Validate arguments
func ValidateArgs() error {
	if !containsStringInSlice(extList, *SrcExt) {
		return fmt.Errorf("the %v must be selected from: %v", argSrcExt, extListStr)
	}

	if !containsStringInSlice(extList, *DstExt) {
		return fmt.Errorf("the %v must be selected from: %v", argDstExt, extListStr)
	}

	srcDir, err := fileutil.OpenFile(*SrcDir)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("the %v does not exist: %w", argSrcDir, err)
		} else {
			return fmt.Errorf("faild to open the %v: %w", argSrcDir, err)
		}
	}

	srcDirInfo, err := srcDir.Stat()
	if err != nil {
		return fmt.Errorf("failed to get the %v's info: %w", argSrcDir, err)
	}
	if !srcDirInfo.IsDir() {
		return fmt.Errorf("the %v must be a directory: %v", argSrcDir, *SrcDir)
	}
	if err := srcDir.Close(); err != nil {
		return fmt.Errorf("failed to close the %v: %w", argSrcDir, err)
	}

	if getType(*FileDeleteFlag) != "bool" {
		return fmt.Errorf("the %v must be true or false", argFileDeleteFlag)
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
		if err = os.MkdirAll(dstDir, 0750); err != nil {
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
			if err := os.Remove(srcFilePath); err != nil {
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
	tmpDstDir := filepath.Join(conv.dstDir, relPathSrcDir)
	absPathDstDir, err := filepath.Abs(tmpDstDir)
	if err != nil {
		return "", err
	}
	return absPathDstDir, nil
}

// Get a relative path of a source directory
func (conv *Converter) getRelPathSrcDir(srcFilePath string) (string, error) {
	srcDir := filepath.Dir(srcFilePath)
	return filepath.Rel(conv.srcDir, srcDir)
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
	case extJpg, extJpeg:
		return jpeg.Encode(dstImage, srcImage, nil)
	case extPng:
		return png.Encode(dstImage, srcImage)
	case extGif:
		return gif.Encode(dstImage, srcImage, nil)
	case extTif, extTiff:
		return tiff.Encode(dstImage, srcImage, nil)
	case extBmp:
		return bmp.Encode(dstImage, srcImage)
	default:
		return fmt.Errorf("the %v must be selected from: %v", argDstExt, extListStr)
	}
}

// Get a source image
// (Use named return values to return an error from the inside of defer)
func getSrcImage(srcFilePath string) (srcImage image.Image, err error) {
	srcFile, err := fileutil.OpenFile(srcFilePath)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = srcFile.Close()
	}()
	srcImage, _, err = image.Decode(srcFile)
	return srcImage, err
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
