/*
	ExtensionChange = 拡張子の指定
*/
package conversion

func ExtensionChange(ext string) string {
	var result string

	if ext == "jpeg" {
		result = "jpegなんだけど"
	} else {
		result = "チャーーーーリーーーー！！"
	}

	return result
}
