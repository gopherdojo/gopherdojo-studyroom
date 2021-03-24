# 画像変換コマンド
[Gopher道場](https://gopherdojo.org/)の課題1です。

# 機能
コマンドライン引数に指定したディレクトリ以下の画像ファイルを変換します。

## 対応ファイル形式

- jpg -> png
- png -> jpg

※ その他の形式のファイル/ディレクトリは無視されます。

# 使い方
## コンパイル

```
go build convert.go
```

## 実行
jpgからpng

```
./convert directoryname -i=jpg -o=png
```

pngからjpg

```
./convert directoryname -i=png -o=jpg
```

## オプション

`-i`: 入力画像形式（デフォルトpng）
`-o`: 出力画像形式（デフォルトjpg）
