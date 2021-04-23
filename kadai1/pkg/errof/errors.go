package errof

type UserErr string

func (e UserErr) Error() (msg string) {
	var ok bool
	if msg, ok = ErrCodeNames[e]; !ok {
		return string(e)
	}
	return msg
}

var ErrCodeNames = map[UserErr]string{
	ErrWalkingSrcPath:  "指定されたディレクトリのトラバースに失敗しました",
	ErrCreateDirectory: "指定されたディクトリの作成に失敗しました",
	ErrGetSrcFileInfo:  "指定されたファイル情報の取得に失敗しました",
	ErrGetDirInfo:      "指定されたディレクトリ情報の取得に失敗しました",
	ErrOpenSrcFile:     "指定されたファイルの展開に失敗しました",
	ErrCloseSrcFile:    "指定されたファイルを閉じるのに失敗しました",
	ErrRemoveSrcFile:   "指定されたファイルの削除に失敗しました",
	ErrDecodeImage:     "指定されたイメージのデコードに失敗しました",
	ErrCreateDstFile:   "指定されたファイルの作成に失敗しました",
	ErrEncodePNGImg:    "pngファイルのエンコードに失敗しました",
	ErrEncodeJPGImg:    "jpgファイルのエンコードに失敗しました",
	ErrEncodeGIFImg:    "gifファイルのエンコードに失敗しました",
}

var (
	ErrWalkingSrcPath  UserErr = "ErrWalkingSrcPath"
	ErrCreateDirectory UserErr = "ErrCreateDirectory"
	ErrGetSrcFileInfo  UserErr = "ErrGetSrcFileInfo"
	ErrGetDirInfo      UserErr = "ErrGetDirInfo"
	ErrOpenSrcFile     UserErr = "ErrOpenFile"
	ErrCloseSrcFile    UserErr = "ErrCloseSrcFile"
	ErrRemoveSrcFile   UserErr = "ErrRemoveSrcFile"
	ErrDecodeImage     UserErr = "ErrDecodeImage"
	ErrCreateDstFile   UserErr = "ErrCreateDstFile"
	ErrEncodePNGImg    UserErr = "ErrEncodePNGImg"
	ErrEncodeJPGImg    UserErr = "ErrEncodeJPGImg"
	ErrEncodeGIFImg    UserErr = "ErrEncodeGIFImg"
)
