package constant

import "strings"

// Extension は拡張子を表します
type Extension string

// Extensions は拡張子のスライスを表します
type Extensions []Extension

// S は拡張子文字列を返却します。
func (e Extension) S() string {
	return string(e)
}

// Contains は指定された拡張子が含まれるか検証します。
// なお、大文字小文字の区別は行いません。
func (es Extensions) Contains(ext string) bool {
	ext = strings.ToLower(ext)
	for _, e := range es {
		if e == Extension(ext) {
			return true
		}
	}
	return false
}

const (
	ExtensionJpeg Extension = "jpeg" // jpeg を表します
	ExtensionPng  Extension = "png"  // png を表します
	ExtensionGif  Extension = "gif"  // gif を表します
	ExtensionBmp  Extension = "bmp"  // bmp を表します
	ExtensionTiff Extension = "tiff" // tiff を表します
)

// AllExtension はサポートしている全ての拡張子を表します
var AllExtension = Extensions{
	ExtensionJpeg,
	ExtensionPng,
	ExtensionGif,
	ExtensionBmp,
	ExtensionTiff,
}
