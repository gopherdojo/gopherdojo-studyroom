package convImages

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

type imgConverter struct {
	name  string
	Image image.Image
}

func (c *imgConverter) decode(buf io.Reader, imgFmt string) error {
	var err error
	switch imgFmt {
	case "jpg":
		c.Image, err = jpeg.Decode(buf)
		break
	case "png":
		c.Image, err = png.Decode(buf)
		break
	case "gif":
		c.Image, err = gif.Decode(buf)
	default:
		err = errors.New("decode format is incorrect")
	}
	return err
}

func (c *imgConverter) encode(buf io.Writer, imgFmt string) error {
	// 変換
	switch imgFmt {
	case "png":
		return png.Encode(buf, c.Image)
	case "gif":
		return gif.Encode(buf, c.Image, nil)
	case "jpg":
		return jpeg.Encode(buf, c.Image, nil)
	default:
		return errors.New("encode format is incorrect")
	}
	return nil
}
