# 課題1 画像変換コマンドの作成

## 仕様

- ディレクトリの指定
- 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
- ディレクトリ以下は再帰的に処理
- 変換前と変換後の画像形式を指定できる（オプション）
- mainパッケージと分離する\
- 自作パッケージと標準パッケージと準標準パッケージのみ使う
- ユーザ定義型の定義
- GoDocを生成
- Go Modulesを使う

## ディレクトリの指定・再帰処理、JPGファイル->PNG変換（デフォルト）

```zsh
$ cd kadai1/yossy0806
$ tree testdata
testdata
└── a
    ├── b
    │   └── yoshineko.jpg
    ├── sample.txt
    └── yoshineko.jpg
$ make build
$ bin/imgconv -dir testdata
the image conversion was successful.
$ tree testdata
testdata
└── a
    ├── b
    │   ├── yoshineko.jpg
    │   └── yoshineko.png
    ├── sample.txt
    ├── yoshineko.jpg
    └── yoshineko.png
```

## 画像形式の指定

```zsh
$ bin/imgconv --help
Usage of bin/imgconv:
  -de string
      変更後の画像ファイルの指定(jpg|jpeg|png|gif) (default "png")
  -dir string
      変換対象のディレクトリの指定
  -se string
      変更前の画像ファイルの指定(jpg|jpeg|png|gif) (default "jpg")
```
