# 画像変換コマンドを作る

- [x] ディレクトリを指定する
- [x] 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
- [x] ディレクトリ以下は再帰的に処理する
- [x] 変換前と変換後の画像形式を指定できる

- [x] mainパッケージと分離する
- [x] 自作パッケージと標準パッケージと準標準パッケージのみ使う
- [x] ユーザ定義型を作ってみる
- [x] GoDocを生成してみる
- [x] Go Modulesを使ってみる

## ファイル構成
image_trans以下が画像変換プログラム

- conversion/conversion.go
  
画像変換の本体

- findFile/findFile.go

対象ファイル検索

## オプション

|option|Contents|default|
|----|----|----|
|`-i_fmt` |入力ファイルの拡張子|jpeg|
|`-o_fmt`|出力ファイルの拡張子|png|

## 使用方法

```
#拡張子指定なし
./image_trans ./testdir
# 拡張子指定
./image_trans -i_fmt jpeg -o_fmt png ./testdir
```
