# 仕様

- ディレクトリを指定する
- 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
- ディレクトリ以下は再帰的に処理する
- 変換前と変換後の画像形式を指定できる（オプション）

# 条件

- mainパッケージと分離する
- 自作パッケージと標準パッケージと準標準パッケージのみ使う
  - 準標準パッケージ：golang.org/x以下のパッケージ
- ユーザ定義型を作ってみる
- GoDocを生成してみる
- Go Modulesを使ってみる

# Usage

```shell script
$ cd kadai1

$ chmod +x ./testdata.sh
$ ./testdata.sh

$ go build main.go

$ ./main [-from=<ext>] [-to=<ext>] testdata

-from, -to =<gif, jpeg, jpg, png>

## for example
$ ./main testdata
変換が完了しました

$ ./main -from=png -to=jpg
変換が完了しました
```