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

func (ic *imgConverter) decode(r io.Reader, imgFmt string) (err error) {
	switch imgFmt {
	case "jpg":
		ic.Image, err = jpeg.Decode(r)
		break
	case "png":
		ic.Image, err = png.Decode(r)
		break
	case "gif":
		ic.Image, err = gif.Decode(r)
	default:
		err = errors.New("decode format is incorrect")
	}
	return
}

func (ic *imgConverter) encode(w io.Writer, imgFmt string) (err error) {
	// 変換
	switch imgFmt {
	case "png":
		return png.Encode(w, ic.Image)
	case "gif":
		return gif.Encode(w, ic.Image, nil)
	case "jpg":
		return jpeg.Encode(w, ic.Image, nil)
	default:
		return errors.New("encode format is incorrect")
	}
	return
}
