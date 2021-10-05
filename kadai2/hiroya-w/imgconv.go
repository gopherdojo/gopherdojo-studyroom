/*
	Package imgconv provides image converter functions.
	JPG, PNG, and GIF are supported.
*/

package imgconv

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"path/filepath"
)

// Decoder is an interface for image decoding.
type Decoder interface {
	Decode(r io.Reader) (image.Image, error)
	GetExt() string
}

// Encoder is an interface for image encoding.
type Encoder interface {
	Encode(w io.Writer, m image.Image) error
	GetExt() string
}

// Config is the configuration for arguments and flags.
type Config struct {
	InputType  string
	OutputType string
	Directory  string
}

// Extention is a struct for holding the file extension.
type Extention struct {
	Ext string
}

// GetExtt returns the file extension.
func (e *Extention) GetExt() string {
	return e.Ext
}

// ImageDecoder is an image decoder for jpg, png, and gif.
type ImageDecoder struct {
	*Extention
}

func (d *ImageDecoder) Decode(r io.Reader) (image.Image, error) {
	img, _, err := image.Decode(r)
	return img, err
}

// JPGEncoder is an image encoder for jpg.
type JPGEncoder struct {
	*Extention
}

func (e *JPGEncoder) Encode(w io.Writer, m image.Image) error {
	return jpeg.Encode(w, m, nil)
}

// PNGEncoder is an image encoder for png.
type PNGEncoder struct {
	*Extention
}

func (e *PNGEncoder) Encode(w io.Writer, m image.Image) error {
	return png.Encode(w, m)
}

// GIFEncoder is an image encoder for gif.
type GIFEncoder struct {
	*Extention
}

func (e *GIFEncoder) Encode(w io.Writer, m image.Image) error {
	return gif.Encode(w, m, nil)
}

// ImgConv is the main struct for the image converter.
type ImgConv struct {
	Decoder   Decoder
	Encoder   Encoder
	TargetDir string
}

// NewDecoder returns a new image decoder.
func NewDecoder(inputType string) (Decoder, error) {
	switch inputType {
	case "jpg", "png", "gif":
		return &ImageDecoder{&Extention{inputType}}, nil
	default:
		return nil, fmt.Errorf("%s is not a supported image type", inputType)
	}
}

// NewEncoder returns a new image encoder for the given outputType.
func NewEncoder(outputType string) (Encoder, error) {
	switch outputType {
	case "jpg":
		return &JPGEncoder{&Extention{outputType}}, nil
	case "png":
		return &PNGEncoder{&Extention{outputType}}, nil
	case "gif":
		return &GIFEncoder{&Extention{outputType}}, nil
	default:
		return nil, fmt.Errorf("unsupported output type: %s", outputType)
	}
}

// renameExt renames the file extension of the file at filePath to newExt.
func renameExt(filePath, newExt string) string {
	return filePath[:len(filePath)-len(filepath.Ext(filePath))] + "." + newExt
}

// GetFiles returns a slice of file paths for specific extension in the target directory.
// Decoder.GetExt() is used to determine the extension.
func (c *ImgConv) GetFiles() ([]string, error) {
	var imgPaths []string

	if f, err := os.Stat(c.TargetDir); err != nil {
		return nil, err
	} else if !f.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", c.TargetDir)
	}

	err := filepath.Walk(c.TargetDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == "."+c.Decoder.GetExt() {
			imgPaths = append(imgPaths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return imgPaths, nil
}

// Convert converts an image file at filePath to the outputType.
func (c *ImgConv) Convert(dec Decoder, enc Encoder, filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Printf("Error closing file: %s\n", err)
		}
	}()

	img, _, err := image.Decode(f)
	if err != nil {
		return "", err
	}

	outputPath := renameExt(filePath, enc.GetExt())
	output, err := os.Create(outputPath)
	if err != nil {
		return "", err
	}
	defer func() {
		if err := output.Close(); err != nil {
			log.Printf("Error closing file: %s\n", err)
		}
	}()

	if err := enc.Encode(output, img); err != nil {
		return "", err
	}
	return outputPath, nil
}

// Run converts all images in the target directory to the outputType.
func (c *ImgConv) Run() ([]string, error) {
	var convertedFiles []string
	imgPaths, err := c.GetFiles()
	if err != nil {
		return nil, err
	}

	for _, path := range imgPaths {
		outputPath, err := c.Convert(c.Decoder, c.Encoder, path)
		if err != nil {
			return convertedFiles, err
		}
		convertedFiles = append(convertedFiles, outputPath)
	}

	return convertedFiles, nil
}
