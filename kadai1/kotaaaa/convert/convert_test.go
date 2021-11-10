package convert

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestNewFileInfo(t *testing.T) {
	srcFilename := "sample.jpg"
	dstExt := ".png"
	basePath := "./dir"
	result := NewFileInfo(srcFilename, dstExt, basePath)

	if result.srcFilename != srcFilename {
		t.Errorf("Result: %v, Expected: %v", result.srcFilename, srcFilename)
	}
	if result.dstExt != dstExt {
		t.Errorf("Result: %v, Expected: %v", result.dstExt, dstExt)
	}
	if result.basePath != basePath {
		t.Errorf("Result: %v, Expected: %v", result.basePath, basePath)
	}
}

func TestRemoveFile(t *testing.T) {
	fileName := "../testdata/sample.log"
	removeFile(fileName)
	if _, err := os.Stat(fileName); err == nil {
		t.Errorf("File still exists: %v", fileName)
	}
	fp, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()
	fp.WriteString("hello")
}

func TestGetFilePathFromBaseSuccess(t *testing.T) {
	path := "./dir/sample.jpg"
	expected := "./dir/sample"
	result := getFilePathFromBase(path)
	if result != expected {
		t.Errorf("Result: %v, Expected: %v", result, expected)
	}
}

func TestGetFilePathFromBaseFail(t *testing.T) {
	path := "./dir/sample.jpg"
	expected := "./dir/sample.jpg"
	result := getFilePathFromBase(path)
	if result == expected {
		t.Errorf("Result: %v, Expected: %v", result, expected)
	}
}

func TestConvert(t *testing.T) {

	srcFilename := "sample.jpg"
	dstFilename := "sample.png"
	dstExt := ".png"
	basePath := "../testdata/"

	copyFile(basePath+"owl.jpg", basePath+srcFilename)
	fileInfo := NewFileInfo(srcFilename, dstExt, basePath)
	if !isExist(basePath + srcFilename) {
		t.Errorf("File for test is not created.")
	}
	fileInfo.Convert()
	if !isExist(basePath + dstFilename) {
		t.Errorf("Expected file is not created: %v", dstFilename)
	}
	os.Remove(basePath + dstFilename)
}

func copyFile(srcName string, dstName string) {
	src, err := os.Open(srcName)
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		panic(err)
	}
}

func isExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil
}
