# 課題1 画像変換コマンド

## 仕様

- ディレクトリを指定する
- 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
- ディレクトリ以下は再帰的に処理する
- 変換前と変換後の画像形式を指定できる（オプション）


## 制限/条件

- mainパッケージと分離する
- 自作パッケージと標準パッケージと準標準パッケージのみ使う
  - 準標準パッケージ：golang.org/x以下のパッケージ
- ユーザ定義型を作る
- GoDocを生成
- Go Modulesを使う

## 使い方

- オプションなし
  - jpgファイル→pngファイル

```
go run kadai1/wada/main.go directory_path
```

- オプションあり
  - -before 変換前の拡張子
  - -after  変換後の拡張子

```
go run kadai1/wada/main.go -before ? -after ? directory_path 
```

## 対応している画像フォーマット

- jpg, jpeg
- png
- gif