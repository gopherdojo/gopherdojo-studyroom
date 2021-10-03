package imgconv

import (
	"fmt"
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

type ImgConv struct {
	OutStream io.Writer
}

func NewDecoder() Decoder {
	return &ImageDecoder{}
}

func NewEncoder(outputType string) (Encoder, error) {
	switch outputType {
	case "jpg":
		return &JPGEncoder{}, nil
	case "png":
		return &PNGEncoder{}, nil
	case "gif":
		return &GIFEncoder{}, nil
	default:
		return nil, fmt.Errorf("unsupported output type: %s", outputType)
	}
}

func (c *ImgConv) Run(dec Decoder, enc Encoder, directory string) error {
	return nil
}
