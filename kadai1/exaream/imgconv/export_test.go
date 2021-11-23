package imgconv

var ExportGetAbsPathDstDir = (*Converter).getAbsPathDstDir
var ExportGetRelPathSrcDir = (*Converter).getRelPathSrcDir
var ExportGetDstFilePath = (*Converter).getDstFilePath
var ExportConvert = (*Converter).convert
var ExportEncode = (*Converter).encode
var ExportGetSrcImage = getSrcImage
var ExportGetDstImage = getDstImage
var ExportGetType = getType

var ExportSrcExt = (*Converter).getSrcExt
var ExportDstExt = (*Converter).getDstExt
var ExportSrcDir = (*Converter).getSrcDir
var ExportDstDir = (*Converter).getDstDir
var ExportFileDeleteFlag = (*Converter).getFileDeleteFlag

func (conv *Converter) getSrcExt() string {
	return conv.srcExt
}
func (conv *Converter) getDstExt() string {
	return conv.dstExt
}
func (conv *Converter) getSrcDir() string {
	return conv.srcDir
}
func (conv *Converter) getDstDir() string {
	return conv.dstDir
}
func (conv *Converter) getFileDeleteFlag() bool {
	return conv.fileDeleteFlag
}
