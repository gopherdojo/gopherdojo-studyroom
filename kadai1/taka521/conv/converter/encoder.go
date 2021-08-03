package converter

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/taka521/gopherdojo-studyroom/kadai1/taka521/conv/constant"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
)

// encoder は画像のエンコード関数です。
type encoder func(out io.Writer, img image.Image) error

var encoders = map[constant.Extension]encoder{
	constant.ExtensionJpeg: jpegEncoder,
	constant.ExtensionPng:  pngEncoder,
	constant.ExtensionGif:  gifEncoder,
	constant.ExtensionBmp:  bmpEncoder,
	constant.ExtensionTiff: tiffEncoder,
}

var (
	jpegEncoder encoder = func(out io.Writer, img image.Image) error {
		return jpeg.Encode(out, img, nil)
	}

	pngEncoder encoder = func(out io.Writer, img image.Image) error {
		return png.Encode(out, img)
	}

	gifEncoder encoder = func(out io.Writer, img image.Image) error {
		return gif.Encode(out, img, nil)
	}

	bmpEncoder encoder = func(out io.Writer, img image.Image) error {
		return bmp.Encode(out, img)
	}

	tiffEncoder encoder = func(out io.Writer, img image.Image) error {
		return tiff.Encode(out, img, nil)
	}
)
