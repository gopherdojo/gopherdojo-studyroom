package conv

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func getFiles(path string) {

	fmt.Println(path)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name(), f.IsDir(), filepath.Ext(f.Name()))
	}
}
