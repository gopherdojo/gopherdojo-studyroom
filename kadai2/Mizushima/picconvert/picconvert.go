package picconvert

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

// PicConverter is user-defined type for converting
// a picture file format has root path, pre-conversion format,
// and after-conversion format.
type PicConverter struct {
	Path        string
	PreFormat   []string
	AfterFormat string
}

// Conv converts the picture file format.
func (p *PicConverter) Conv() error {
	files, err := Glob(p.Path, p.PreFormat)

	if err != nil {
		return err
	}

	if files == nil {
		return fmt.Errorf("there's no %s file", p.PreFormat)
	}

	// prosessing for each file.
	for _, file := range files {
		// fmt.Println("from:", file)
		f, err := os.Open(file)
		if err != nil {
			return err
		}
		defer f.Close()

		// reading the image file.
		img, _, err := image.Decode(f)
		if err != nil {
			return err
		}

		// creating filepath for output.
		output, err := os.Create(fmt.Sprintf(
			"%s_converted.%s",
			baseName(file), p.AfterFormat))

		if err != nil {
			return err
		}

		// converting the file.
		switch p.AfterFormat {
		case "png":
			err = png.Encode(output, img)
		case "jpg":
		case "jpeg":
			err = jpeg.Encode(output, img, nil)
		case "gif":
			err = gif.Encode(output, img, nil)
		default:
			err = fmt.Errorf("%s is not supported", p.AfterFormat)
		}

		if err != nil {
			return err
		}

		// fmt.Printf("converted %s\n", output.Name())
	}

	return nil
}

// NewPicConverter is the constructor for a PicConverter.
func NewPicConverter(Path string, PreFormat string, AfterFormat string) *PicConverter {
	res := new(PicConverter)
	res.Path = Path

	if PreFormat == "jpeg" || PreFormat == "jpg" {
		res.PreFormat = []string{"jpeg", "jpg"}
	} else {
		res.PreFormat = []string{PreFormat}
	}

	res.AfterFormat = AfterFormat
	return res
}

// Glob returns a slice of the file paths that meets the "format".
func Glob(path string, format []string) ([]string, error) {
	var res []string

	var err error
	for _, f := range format {
		err = filepath.Walk(path,
			func(path string, info os.FileInfo, err error) error {
				if filepath.Ext(path) == "."+f && !info.IsDir() {
					res = append(res, path)
				}
				return nil
			})
	}

	return res, err
}

// baseName returns the filepath without a extension.
func baseName(filePath string) string {
	ext := filepath.Ext(filePath)
	return filePath[:len(filePath)-len(ext)]
}
