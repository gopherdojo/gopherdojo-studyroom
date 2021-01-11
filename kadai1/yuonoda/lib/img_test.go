package convImages

import (
	"bytes"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"testing"
)

func TestEncode(t *testing.T) {
	t.Logf("TestEncode")

	// 生画像を生成
	image := testGenerateImage(t)

	// サンプルデータでループ
	fmts := []string{"png", "gif", "jpg"}
	for _, fmt := range fmts {

		// エンコード
		ic := imgConverter{"test", image}
		buff1 := bytes.NewBuffer([]byte{})
		err := ic.encode(buff1, fmt)
		if err != nil {
			t.Errorf("failed to encode err:%v", err)
		}

		// imgConverterを使わずにエンコード
		buff2 := bytes.NewBuffer([]byte{})
		switch fmt {
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
		if err != nil {
			t.Errorf("failed to test encode err:%v ", err)
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
