package word

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

//単語の一覧取得
func GetWordList(nofq int) ([][]string, error) {
	//ファイルオープン
	fl, err := os.Open("./word/NAWL_1.0_with_en_definitions.csv")
	//deferのエラーハンドリング
	defer func() {
		if err := fl.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err != nil {
		return nil, err
	}

	//ファイル読み込み
	rd := csv.NewReader(fl)
	rd.Comma = ','
	var wordList [][]string
	//１行ずつ内容を読み込み
	for {
		record, err := rd.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			//読み込みエラー
			return nil, err
		}
		wordList = append(wordList, record)
		if nofq == len(wordList) {
			break
		}
	}
	return wordList, nil
}
