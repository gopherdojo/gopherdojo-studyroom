package imgconv

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

type ext string

const (
	// GIF extension
	GIF ext = "gif"
	// PNG extension
	PNG ext = "png"
	// JPG extension
	JPG ext = "jpg"
	// JPEG extension
	JPEG ext = "jpeg"
)

var extensions = []ext{GIF, PNG, JPG, JPEG}

type target struct {
	inputPath  string
	inputExt   ext
	fileName   string
	outputPath string
	outputExt  ext
}

func (t *target) getInputFile() string {
	return t.inputPath + t.fileName + "." + string(t.inputExt)
}

func (t *target) getOutputFile() string {
	return t.outputPath + t.fileName + "." + string(t.outputExt)
}

func getTargets(id *string, od *string, ie *string, oe *string) ([]target, error) {
	var targets []target
	err := filepath.Walk(*id, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == "."+*ie {
			t := target{
				inputPath:  filepath.Dir(path) + "/",
				inputExt:   ext(*ie),
				fileName:   strings.TrimSuffix(filepath.Base(path), "."+*ie),
				outputPath: strings.TrimRight(*od, "/") + "/",
				outputExt:  ext(*oe),
			}
			targets = append(targets, t)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return targets, nil
}

func (t *target) decode() (image.Image, error) {
	inputFile := t.getInputFile()
	input, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer input.Close()
	img, _, err := image.Decode(input)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func (t *target) encode(img image.Image) error {
	outputFile := t.getOutputFile()
	output, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer output.Close()
	switch t.outputExt {
	case JPG, JPEG:
		err = jpeg.Encode(output, img, nil)
	case PNG:
		err = png.Encode(output, img)
	case GIF:
		err = gif.Encode(output, img, nil)
	default:
		err = errors.New("invalid extension")
	}
	if err != nil {
		return err
	}
	return nil
}

func convert(targets []target) error {
	for _, t := range targets {
		img, err := t.decode()
		if err != nil {
			return err
		}
		if err := t.encode(img); err != nil {
			return err
		}
	}
	return nil
}

// Do performs the conversion and returns an error if it fails.
func Do(id *string, od *string, ie *string, oe *string) error {
	if err := validation(id, od, ie, oe); err != nil {
		return err
	}
	targets, err := getTargets(id, od, ie, oe)
	if err != nil {
		return err
	}
	if err := convert(targets); err != nil {
		return err
	}
	return nil
}
