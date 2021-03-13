package imgcov

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

// A converter changes iamge foramt
type Converter struct {
	SrcFormat  string
	DestFormat string
}

// Convert decodes image from file, then encodes to a different format
func (c *Converter) Convert(src, dest string) (bool, error) {
	in, err := os.Open(src)
	if err != nil {
		return false, err
	}
	defer in.Close()

	// encode image
	img, format, err := image.Decode(in)
	if err != nil {
		return false, nil
	}
	if format != c.SrcFormat {
		return false, nil
	}

	out, err := os.Create(dest)
	if err != nil {
		return false, err
	}
	defer out.Close()

	// decode image
	switch c.DestFormat {
	case "jpeg":
		err = jpeg.Encode(out, img, nil)
	case "png":
		err = png.Encode(out, img)
	default:
		err = errors.New("invalid format: " + c.DestFormat)
	}

	if err != nil {
		return false, err
	}

	return true, nil
}
