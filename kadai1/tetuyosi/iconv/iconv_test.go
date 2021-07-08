package iconv_test

import (
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"reflect"
	"testing"

	iconv "github.com/tetuyosi/gopherdojo-studyroom/kadai1/tetuyosi/iconv"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

// 正常ケース
func Test_New_01(t *testing.T) {
	actual, _ := iconv.New("test.jpg", "jpg", "png")
	expected := iconv.Conv{
		FilePath: "test.jpg",
		SrcExt:   "jpg",
		DestExt:  "png",
	}
	if !reflect.DeepEqual(actual, &expected) {
		t.Errorf("got: %v\nwant: %v", actual, &expected)
	}
}

// エラーケース
func Test_New_02(t *testing.T) {
	_, actual := iconv.New("test.txt", "jpg", "png")
	expected := "変換ファイル拡張子が不正です。"
	if actual.Error() != expected {
		t.Errorf("got: %v\nwant: %v", actual.Error(), expected)
	}
}

func Test_New_03(t *testing.T) {
	_, actual := iconv.New("test.jpg", "txt", "png")
	expected := "変換元フォーマット指定の誤りです。"
	if actual.Error() != expected {
		t.Errorf("got: %v\nwant: %v", actual.Error(), expected)
	}
}

func Test_New_04(t *testing.T) {
	_, actual := iconv.New("test.jpg", "jpg", "txt")
	expected := "変換先フォーマット指定の誤りです。"
	if actual.Error() != expected {
		t.Errorf("got: %v\nwant: %v", actual.Error(), expected)
	}
}

func Test_New_05(t *testing.T) {
	actual, _ := iconv.New("test.JPEG", "jpg", "png")
	expected := iconv.Conv{
		FilePath: "test.JPEG",
		SrcExt:   "jpg",
		DestExt:  "png",
	}
	if !reflect.DeepEqual(actual, &expected) {
		t.Errorf("got: %v\nwant: %v", actual, &expected)
	}
}

// test jpg -> png
func Test_Convert_01(t *testing.T) {
	makeImage("jpg")
	c, _ := iconv.New("test.jpg", "jpg", "png")
	c.Imaging()
	c.Convert()
	_, actual := os.Stat("test.png")
	if actual != nil {
		t.Errorf("convert file not created")
	}
	defer func() {
		os.Remove("test.jpg")
		os.Remove("test.png")
	}()
}

// test png -> jpg
func Test_Convert_02(t *testing.T) {
	makeImage("png")
	c, _ := iconv.New("test.png", "png", "jpg")
	c.Imaging()
	c.Convert()
	_, actual := os.Stat("test.jpg")
	if actual != nil {
		t.Errorf("convert file not created")
	}
	defer func() {
		os.Remove("test.jpg")
		os.Remove("test.png")
	}()
}

// test gif -> jpg
func Test_Convert_03(t *testing.T) {
	makeImage("gif")
	c, _ := iconv.New("test.gif", "gif", "jpg")
	c.Imaging()
	c.Convert()
	_, actual := os.Stat("test.jpg")
	if actual != nil {
		t.Errorf("convert file not created")
	}
	defer func() {
		os.Remove("test.jpg")
		os.Remove("test.gif")
	}()
}

func makeImage(t string) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 50))
	img.Set(2, 3, color.RGBA{255, 0, 0, 255})
	f, _ := os.OpenFile("test."+t, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	switch t {
	case "jpg":
		jpeg.Encode(f, img, &jpeg.Options{})
	case "png":
		png.Encode(f, img)
	case "gif":
		gif.Encode(f, img, &gif.Options{})
	}
}
