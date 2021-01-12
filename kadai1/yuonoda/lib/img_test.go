package convImages

import (
	"bytes"
	"errors"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"testing"
)

func TestDecode(t *testing.T) {
	t.Logf("TestDecode")
	cases := []struct {
		fmt string
		err error
	}{
		{"png", nil},
		{"jpg", nil},
		{"gif", nil},
		{"docx", errors.New("encode format is incorrect")},
	}
	var ic imgConverter

	for _, c := range cases {

		// テスト画像を生成しエンコード
		testImage := testGenerateImage(t)
		ic.Image = testImage
		encodedBuf1 := bytes.NewBuffer([]byte{})
		encodedBuf2 := bytes.NewBuffer([]byte{})
		err := ic.encode(encodedBuf1, c.fmt)
		if err != nil {
			if err.Error() != c.err.Error() {
				t.Errorf(err.Error())
			}
			continue
		}
		err = ic.encode(encodedBuf2, c.fmt)
		if err != nil && err.Error() != c.err.Error() {
			if err.Error() != c.err.Error() {
				t.Errorf(err.Error())
			}
			continue
		}

		// imgConnverterでデコード
		ic = imgConverter{}
		err = ic.decode(encodedBuf1, c.fmt)
		if err != nil {
			t.Errorf("image decode error err: %v", err)
		}
		image1 := ic.Image

		// imgConverterを使わずにデコード
		image2, _, err := image.Decode(encodedBuf2)
		if err != nil {
			t.Errorf("image test decode error err:%v", err)
		}

		// サイズの比較
		if image1.Bounds() != image2.Bounds() {
			t.Errorf("image size doesn't match")
		}

		// 色の比較
		for x := 0; x < image1.Bounds().Max.X; x++ {
			for y := 0; y < image1.Bounds().Max.Y; y++ {
				color1 := image1.At(x, y)
				color2 := image2.At(x, y)
				if color1 != color2 {
					t.Errorf("color doesn't match")
				}
			}
		}
	}
	return

}

func TestEncode(t *testing.T) {
	t.Logf("TestEncode")

	// 生画像を生成
	image := testGenerateImage(t)

	// サンプルデータでループ
	cases := []struct {
		fmt   string
		error error
	}{
		{"png", nil},
		{"jpg", nil},
		{"gif", nil},
		{"docx", errors.New("encode format is incorrect")},
	}
	for _, c := range cases {

		// エンコード
		ic := imgConverter{"test", image}
		buff1 := bytes.NewBuffer([]byte{})
		err := ic.encode(buff1, c.fmt)
		if err != nil && err.Error() != c.error.Error() {
			t.Errorf(err.Error())
		}

		// imgConverterを使わずにエンコード
		buff2 := bytes.NewBuffer([]byte{})
		switch c.fmt {
		case "png":
			err = png.Encode(buff2, image)
			break
		case "jpg":
			err = jpeg.Encode(buff2, image, nil)
			break
		case "gif":
			err = gif.Encode(buff2, image, nil)
			break
		}
		if err != nil && err.Error() != c.error.Error() {
			t.Errorf(err.Error())
		}

		// 双方のエンコード結果が一致するか検証
		if bytes.Compare(buff1.Bytes(), buff2.Bytes()) != 0 {
			t.Errorf("encode result mismatch")
		}
	}

	return
}

// インメモリイメージを作成する
func testGenerateImage(t *testing.T) image.Image {
	t.Log("generateImage")
	t.Helper()
	// 画像を作成
	w := 200
	h := 100
	upLeft := image.Point{0, 0}
	lowRight := image.Point{w, h}
	img := image.NewRGBA64(image.Rectangle{upLeft, lowRight})
	cyan := color.RGBA{100, 200, 200, 0xff}

	// 画像を塗りつぶし
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			img.Set(x, y, cyan)
		}
	}

	return img
}
