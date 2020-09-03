package imgconv

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// Define supported extentions
const (
	JPG  string = "jpg"
	JPEG string = "jpeg"
	PNG  string = "png"
	GIF  string = "gif"
)

// Converter define structure of converting informations.
type Converter struct {
	SrcFile string
	SrcExt  string
	DstExt  string
}

// NewConverter construct and return Converter.
func NewConverter(sf, se, de string) *Converter {
	return &Converter{SrcFile: sf, SrcExt: se, DstExt: de}
}

// Validate validates the extension.
func (c *Converter) Validate() error {
	seOk := validateExt(c.SrcExt)
	deOk := validateExt(c.DstExt)

	if !seOk && !deOk {
		return fmt.Errorf("an unsupported extension is specified: se=%s, de=%s", c.SrcExt, c.DstExt)
	} else if !seOk {
		return fmt.Errorf("an unsupported extension is specified: se=%s", c.SrcExt)
	} else if !deOk {
		return fmt.Errorf("an unsupported extension is specified: de=%s", c.DstExt)
	}

	return nil
}

func validateExt(ext string) bool {
	var supportExt []string
	supportExt = append(supportExt, JPG, JPEG, PNG, GIF)
	e := strings.ToLower(ext)

	for _, supext := range supportExt {
		if supext == e {
			return true
		}
	}
	return false
}

// Decode returns ImageData which compose file path and decoded image data.
func (c *Converter) Decode() (img image.Image, err error) {
	sf, err := os.Open(filepath.Clean(c.SrcFile))
	if err != nil {
		return nil, fmt.Errorf("the image file could not be opened: file=%s, err=%v", c.SrcFile, err)
	}
	defer func() {
		if rerr := sf.Close(); rerr != nil {
			err = fmt.Errorf("the image file could not be closed: %v, the original error: %v", rerr, err)
		}
	}()

	switch c.SrcExt {
	case JPG, JPEG:
		img, err = jpeg.Decode(sf)
	case PNG:
		img, err = png.Decode(sf)
	case GIF:
		img, err = gif.Decode(sf)
	default:
		return nil, fmt.Errorf("decode could not be performed because the extension is not supported: se=%s", c.SrcExt)
	}
	if err != nil {
		fmt.Println(c.SrcExt)
		return nil, fmt.Errorf("decode of image file failed: file=%s, err=%v", c.SrcFile, err)
	}

	return img, nil
}

// Encode encodes image file.
func (c *Converter) Encode(img image.Image) (err error) {
	dst := strings.TrimSuffix(c.SrcFile, filepath.Ext(c.SrcFile)) + "." + c.DstExt
	df, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("the image file could not be encoded: file=%s, err=%v", c.SrcFile, err)
	}
	defer func() {
		if rerr := df.Close(); rerr != nil {
			err = fmt.Errorf("the image file could not be closed. %v, the original error: %v", rerr, err)
		}
	}()

	switch c.DstExt {
	case JPG, JPEG:
		err = jpeg.Encode(df, img, nil)
	case PNG:
		err = png.Encode(df, img)
	case GIF:
		err = gif.Encode(df, img, nil)
	default:
		return fmt.Errorf("could not encode because of an unsupported extension: de=%s", c.DstExt)
	}
	if err != nil {
		return fmt.Errorf("failed to encode the image file: file=%s, err=%v", c.SrcFile, err)
	}

	return nil
}
