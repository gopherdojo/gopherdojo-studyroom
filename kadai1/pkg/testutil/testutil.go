package testutil

import (
	"log"
	"os"
)

// PrepareTestSrc:
func PrepareTestSrc(dirPath string, srcPaths []string) {
	var err error
	if err := os.MkdirAll(dirPath, 0777); err != nil {
		log.Fatal(err)
	}
	for _, srcPath := range srcPaths {
		if _, err = os.Create(srcPath); err != nil {
			log.Fatal(err)
		}
	}
}

// RemoveAllTestSrc:
func RemoveAllTestSrc(srcPath string) {
	if err := os.RemoveAll(srcPath); err != nil {
		log.Fatal(err)
	}
}

// RemoveTestFile:
func RemoveTestFile(srcPath string) {
	if err := os.Remove(srcPath); err != nil {
		log.Fatal(err)
	}
}
