# 課題1 画像変換コマンドの作成

## 仕様

- ディレクトリの指定 → flagで指定できるように
- 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト） → ○
- ディレクトリ以下は再帰的に処理 → filepath.Walkを使用
- 変換前と変換後の画像形式を指定できる（オプション） → flagで指定できるように
- mainパッケージと分離する → imgconvを作成
- 自作パッケージと標準パッケージと準標準パッケージのみ使う → ○
- ユーザ定義型の定義 → package imgconvでConverterを定義
- GoDocを生成 → 作成されることを確認
- Go Modulesを使う → 使用

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

- 以下の拡張子に対応（対応していない拡張子が指定されていればerrを返す）
  - jpg
  - jpeg
  - png
  - gif

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
