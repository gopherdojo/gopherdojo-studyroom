package conversion

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
	"youenn/fileutil"
)

type PicType int

type ConvertInfo struct {
	path     string
	src, dst PicType
}

//check if a type name
func IsSupported(tp string) bool {
	if _, ok := SupportedFormat[tp]; ok {
		return true
	}
	return false
}

//check if two type names are actually same type
func IsSameType(tp1 string, tp2 string) bool {
	if !IsSupported(tp1) || !IsSupported(tp2) {
		return false
	}
	return SupportedFormat[tp1] == SupportedFormat[tp2]
}

//Convert a picture from one type to another type
func ConvertPic(info ConvertInfo) bool {
	var srcImage image.Image
	path := info.path
	newFileName := strings.TrimSuffix(path, filepath.Ext(path)) + "." + PicType2String[info.dst]

	//Open source file
	srcFile, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot open "+path)
		return false
	}
	defer srcFile.Close()

	//decode source file
	switch PicType2String[info.src] {
	case "jpeg", "jpg":
		srcImage, err = jpeg.Decode(srcFile)
	case "png":
		srcImage, err = png.Decode(srcFile)
	default:
		fmt.Fprintln(os.Stderr, "Parameter error.")
		return false
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, path+" decode error")
		return false
	}

	//Open Destination file
	dstFile, err := os.Create(newFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot create "+newFileName)
		return false
	}
	defer dstFile.Close()

	//encode image to new file
	switch PicType2String[info.dst] {
	case "jpeg":
		err = jpeg.Encode(dstFile, srcImage, nil)
	case "png":
		err = png.Encode(dstFile, srcImage)
	default:
		fmt.Fprintln(os.Stderr, "Parameter error.")
		return false
	}

	return true
}

//recursively walk through a directory, convert picture from src type to dst type
//return how many files have been changed
func WalkConvert(targetPath string, src string, dst string) int {
	var cnt int
	err := filepath.Walk(targetPath,
		func(path string, info os.FileInfo, err error) error {
			ext := fileutil.GetPureExt(path)
			if IsSameType(ext, src) {
				success := ConvertPic(ConvertInfo{path: path, src: SupportedFormat[src], dst: SupportedFormat[dst]})
				if success {
					cnt++
					if errDelete := os.Remove(path); errDelete != nil {
						fmt.Fprintln(os.Stderr, "fail to delete "+path)
					}
				}
			}
			return nil
		})

	if err != nil {
		log.Fatal("Conversion failed.")
	}
	return cnt
}
