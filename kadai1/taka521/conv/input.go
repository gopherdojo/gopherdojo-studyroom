package conv

import (
	"fmt"
	"os"

	"github.com/taka521/gopherdojo-studyroom/kadai1/taka521/conv/constant"
)

type HandlerInput struct {
	Dir  string // 対象ディレクトリのパス
	From string // 変換元拡張子
	To   string // 変換先拡張子
}

// Validate はバリデーションチェックを行い、違反している場合は error を返却します。
func (i HandlerInput) Validate() error {
	if f, err := os.Stat(i.Dir); os.IsNotExist(err) || !f.IsDir() {
		return fmt.Errorf("指定されたディレクトリは存在しません。[directory = %q]", i.Dir)
	}

	if err := validateExtension(i.From); err != nil {
		return err
	}

	if err := validateExtension(i.To); err != nil {
		return err
	}

	return nil
}

func validateExtension(ext string) error {
	if !constant.AllExtension.Contains(ext) {
		return fmt.Errorf("拡張子 %s はサポートされていません。指定可能な拡張子は %v です。", ext, constant.AllExtension)
	}
	return nil
}
