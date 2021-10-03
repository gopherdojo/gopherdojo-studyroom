package imgconv

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

type Decoder interface {
	Decode(r io.Reader) (image.Image, error)
}

type Encoder interface {
	Encode(w io.Writer, m image.Image) error
}

type Converter interface {
	Decoder
	Encoder
}

type Config struct {
	InputType  string
	OutputType string
	Directory  string
}

type ImageDecoder struct {
}

func (d *ImageDecoder) Decode(r io.Reader) (image.Image, error) {
	img, _, err := image.Decode(r)
	return img, err
}

type JPGEncoder struct {
}

func (e *JPGEncoder) Encode(w io.Writer, m image.Image) error {
	return jpeg.Encode(w, m, nil)
}

type PNGEncoder struct {
}

func (e *PNGEncoder) Encode(w io.Writer, m image.Image) error {
	return png.Encode(w, m)
}

type GIFEncoder struct {
}

func (e *GIFEncoder) Encode(w io.Writer, m image.Image) error {
	return gif.Encode(w, m, nil)
}

type ImageConverter struct {
	Decoder
	Encoder
}

type ImgConv struct {
	OutStream io.Writer
}

func NewConverter(config *Config) Converter {
	var encorder Encoder
	var decorder Decoder = &ImageDecoder{}

	switch config.InputType {
	case "jpg":
		encorder = &JPGEncoder{}
	case "png":
		encorder = &PNGEncoder{}
	case "gif":
		encorder = &GIFEncoder{}
	}

	return &ImageConverter{
		decorder,
		encorder,
	}
}

func (c *ImgConv) Run(converter Converter, directory string) error {
	return nil
}
