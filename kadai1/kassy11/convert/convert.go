package convert

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

func Do(dir string, outdir string, inputFormat string, outputFormat string) {
	err := os.Mkdir(outdir, 0777)
	logError(err, "cannot create directory")
	error := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == "."+inputFormat {
			file, err := os.Open(path)
			logError(err, "Cannot open file")
			defer file.Close()
			img, _, err := image.Decode(file)
			logError(err, "Failed to convert file to image")
			if outputFormat == format.Png {
				dstPath := filepath.Join(outdir, getFileNameWithoutExt(path)+".png")
				out, err := os.Create(dstPath)
				logError(err, "Failed to create destination path")
				defer out.Close()
				png.Encode(out, img)
			} else if outputFormat == format.Jpg || outputFormat == format.Jpeg {
				dstPath := filepath.Join(outdir, getFileNameWithoutExt(path)+".jpg")
				out, err := os.Create(dstPath)
				logError(err, "Failed to create destination path")
				defer out.Close()
				jpeg.Encode(out, img, nil)
			}
		}
		if err != nil {
			return err
		}
		return nil
	})
	logError(error, "Error on filepath.Walk")
	fmt.Println("Succuessfully convert image files")
	fmt.Printf("Check %s directory\n", outdir)
}
