package convImages

import (
	"testing"
)

func TestImgEncode(t *testing.T) {
	t.Logf("TestImgEncode")
	//// 画像をデコード
	//nameNoExt := "test"
	//img := img{nameNoExt, nil}
	//buffer := bytes.NewReader(imageBytes)
	//err = img.decode(buffer, *fromFmt)
	//if err != nil {
	//	log.Fatal(err)
	//}
	expect := "a"
	actual := "a"
	if expect != actual {
		t.Errorf("expect=%s, actual=%s", expect, actual)
	}

}
