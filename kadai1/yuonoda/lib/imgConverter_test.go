package convImages_test

import (
	"bytes"
	"errors"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"testing"

	"github.com/yuonoda/gopherdojo-studyroom/kadai1/yuonoda/lib"
)

// Decodeメソッドのテスト
func TestDecode(t *testing.T) {
	t.Logf("TestDecode")

	// テストケース
	cases := []struct {
		name string
		fmt  string
		err  error
	}{
		{"decodePng", "png", nil},
		{"decodeJpg", "jpg", nil},
		{"decodeGif", "gif", nil},
		{"decodeDocx", "docx", errors.New("encode format is incorrect")},
	}
	var ic convImages.ImgConverter

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {

			// テスト画像ファイルを作成
			testImage := testGenerateImage(t)
			ic.Image = testImage
			encodedBuf1 := bytes.NewBuffer([]byte{})
			encodedBuf2 := bytes.NewBuffer([]byte{})

			// テスト画像のエンコード
			err := ic.Encode(encodedBuf1, c.fmt)
			if err != nil {
				if err.Error() != c.err.Error() {
					t.Errorf(err.Error())
				}
				return
			}
			err = ic.Encode(encodedBuf2, c.fmt)
			if err != nil && err.Error() != c.err.Error() {
				if err.Error() != c.err.Error() {
					t.Errorf(err.Error())
				}
				return
			}

			// imgConnverterでデコード
			ic = convImages.ImgConverter{}
			err = ic.Decode(encodedBuf1, c.fmt)
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
		})

	}
	return

}

// エンコードのテスト
func TestEncode(t *testing.T) {
	// 生画像を生成
	image := testGenerateImage(t)

	// サンプルデータでループ
	cases := []struct {
		name          string
		fmt           string
		isErrExpected bool
	}{
		{"encodePng", "png", false},
		{"encodeJpg", "jpg", false},
		{"encodeGif", "gif", false},
		{"encodeDocx", "docx", true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {

			// エンコード
			ic := convImages.ImgConverter{"test", image}
			buff1 := bytes.NewBuffer([]byte{})
			err := ic.Encode(buff1, c.fmt)
			if err != nil && !c.isErrExpected {
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
			if err != nil && !c.isErrExpected {
				t.Errorf(err.Error())
			}

			// 双方のバイト列が一致するか検証
			if bytes.Compare(buff1.Bytes(), buff2.Bytes()) != 0 {
				t.Errorf("encode result mismatch")
			}
		})

	}

	return
}

// インメモリイメージを作成する
func testGenerateImage(t *testing.T) image.Image {
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
