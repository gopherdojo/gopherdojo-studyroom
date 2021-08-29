package convert

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetSelectedExtensionPath(fileType string, directory string) [][]string {
	var retval [][]string
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		slice := strings.Split(path, ".")
		if slice[len(slice)-1] == fileType {
			retval = append(retval, slice)
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return retval
}
