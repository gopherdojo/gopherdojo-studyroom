package convImages

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"testing"
)

func TestEncode(t *testing.T) {
	t.Logf("TestImgEncode")

	// 生画像を生成
	image := generateImage(t)

	// エンコード
	ic := imgConverter{"test", image}
	buff1 := bytes.NewBuffer([]byte{})
	err := ic.encode(buff1, "png")
	if err != nil {
		t.Errorf("failed to encode err:%v", err)
	}

	// imgConverterを使わずにエンコード
	buff2 := bytes.NewBuffer([]byte{})
	err = png.Encode(buff2, image)
	if err != nil {
		t.Errorf("failed to test encode err:%v ", err)
	}

	// 双方のエンコード結果が一致するか検証
	if bytes.Compare(buff1.Bytes(), buff2.Bytes()) != 0 {
		t.Errorf("encode result mismatch")
	}

	//// 目視用にファイル出力
	//f, _ := os.Create("test1.png")
	//f.Write(buff1.Bytes())
	//
	//
	//// 目視用にファイル出力
	//f, _ = os.Create("test2.png")
	//f.Write(buff2.Bytes())
	return
}

// インメモリイメージを作成する
func generateImage(t *testing.T) image.Image {
	t.Log("generateImage")
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
