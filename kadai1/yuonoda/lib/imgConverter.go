package convImages

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
)

type ImgConverter struct {
	Name  string
	Image image.Image
}

func (ic *ImgConverter) Decode(r io.Reader, imgFmt string) (err error) {
	log.Println("(ic *imgConverter) decode(r io.Reader, imgFmt string) (err error) ")
	ic.Image, _, err = image.Decode(r)
	return
}

func (ic *ImgConverter) Encode(w io.Writer, imgFmt string) (err error) {
	log.Println("(ic *imgConverter) encode(w io.Writer, imgFmt string) (err error) ")
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
}
