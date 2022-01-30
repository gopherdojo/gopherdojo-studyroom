package imgcvt

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)


func Do() error {
	flag.Parse()

	err := filepath.Walk(
		flag.Args()[0],
		func(path string, info os.FileInfo, err error) error {
			fmt.Println(path)
			return nil
	})
	if err != nil {
		return err
	}
	return err
}