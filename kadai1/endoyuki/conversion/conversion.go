package conversion

import (
	"image"
	"image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
)

func Conversion(diraName string, outDirectory string, beforeExt *string, afterExt *string) {
	files, err := filepath.Glob(diraName + "*." + *beforeExt)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fileName := getFileNameWithoutExt(file)

		img, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer img.Close()

		config, _, err := image.Decode(img)
		if err != nil {
			log.Fatal(err)
		}

		out, err := os.Create(outDirectory + fileName + "." + *afterExt)
		if err != nil {
			log.Fatal()
		}
		defer out.Close()

		jpeg.Encode(out, config, nil)
	}
}

func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}
