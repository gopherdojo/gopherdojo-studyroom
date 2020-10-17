/*
Conversion は 画像の拡張子の変更を行うためのパッケージです。

*/
package conversion

const (
	JPEG = "jpeg"
	JPG  = "jpg"
	GIF  = "gif"
	PNG  = "png"
)

func ExtensionCheck(ext string) bool {
	switch ext {
	case JPEG, JPG, GIF, PNG:
		return true
	default:
		return false
	}
}
