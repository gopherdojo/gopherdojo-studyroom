package imgconv

import (
	"testing"
)

func TestDo(t *testing.T) {

}
func TestToString(t *testing.T) {
	var gif ImageExt = "gif"
	if gif.toString() != "gif" {
		t.Errorf("toString = %s; want gif", gif)
	}
}

func TestToImageExt(t *testing.T) {
}
func TestNewImageConverter(t *testing.T) {
}
func TestReadImage(t *testing.T) {
}
func TestSaveImage(t *testing.T) {
}
func TestGetFileNameWithoutExt(t *testing.T) {
}
func TestConvert(t *testing.T) {
}
func TestConvertAll(t *testing.T) {
}
