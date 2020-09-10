## 仕様

- ディレクトリを指定する
  
- 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
  
- ディレクトリ以下は再帰的に処理する
  
- 変換前と変換後の画像形式を指定できる（オプション）
  

## オプション

| オプション | 内容  | デフォルト |
| --- | --- | --- |
| -path | 探索を始めるディレクトリ | $PWD |
| -b  | 変換前の拡張子 | jpeg |
| -a  | 変換後の拡張子 | png |
| -d  | 変換前のファイルを削除する | false |

## 使い方

##### 事前準備

```shell
$ make # テストデータ作成＆ビルド
$ make create # テストデータ作成
$ make clean # テストデータ削除
```

##### 変換コマンドの一例

```shell
$ # jpg to png （再帰的に処理）
$ ./imgconv


$ # gif to jpg（ファイルを削除）
$ ./imgconv -b=gif -a=jpg -d
```

## ディレクトリ構造

```tree
shinnosuke-K
├── conv
│   ├── convert.go
│   └── convert_test.go
├── file
│   ├── files.go
│   └── files_test.go
├── Makefile
├── README.md
├── go.mod
├── init.sh
├── main.go
└── standard.png
```

### 懸念点

変換する処理が指定した拡張子へエンコードする処理以外に共通部分が多い。

戻り値やCloseの処理を考えると切り分ける方法が思いつきませんでした。

懸念点の処理を行っている関数

- [convertToPNG](https://github.com/shinnosuke-K/gopherdojo-studyroom/blob/kadai1-shinnosuke-K/kadai1/shinnosuke-K/conv/convert.go#L91)
  
- [convertToJPG](https://github.com/shinnosuke-K/gopherdojo-studyroom/blob/kadai1-shinnosuke-K/kadai1/shinnosuke-K/conv/convert.go#L111)
  
- [convertToGIF](https://github.com/shinnosuke-K/gopherdojo-studyroom/blob/kadai1-shinnosuke-K/kadai1/shinnosuke-K/conv/convert.go#L131)