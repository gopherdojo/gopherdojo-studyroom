# 仕様
- ディレクトリを指定する
- 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
- ディレクトリ以下は再帰的に処理する
- 変換前と変換後の画像形式を指定できる（オプション）


# 開発条件
- mainパッケージと分離する
  - 変換ロジックをconvimgパッケージに分離
- 自作パッケージと標準パッケージと準標準パッケージのみ使う
- ユーザ定義型を作ってみる
  - convimgパッケージにて、`Ext`型を使用
- GoDocを生成してみる
- Go Modulesを使ってみる


# オプション
|オプション|説明|デフォルト|対応
|:---:|:---:|:---:|:---:|
|-from|変換前の拡張子を指定|jpg|jpg・png・gif|
|-to|変換後の拡張子を指定|png|jpg・png・gif|
|-r|変換元のファイルを削除|false|-|


# 使い方
```shell
$ cd {Path_To_Repository}/kadai1/shinji
$ ./testdata.zsh    # テストデータ生成&初期化
$ go build main.go  # ビルド
```

```shell
$ # jpg -> png（デフォルト）
$ ./main ./testdata

$ # png -> jpg（変換前後の拡張子指定）
$ ./main -from=.png -to=.jpg ./testdata

$ # jpg -> png（変換元のファイルを削除）
$ ./main -r ./testdata

$ # png -> gif（変換前後の拡張子指定、変換元のファイルを削除）
$ ./main -r -from=.png -to=.gif ./testdata
```

# テストデータのディレクトリ構造
```
testdata
├── azarashi.jpg
├── tanuki.jpg
├── osaru.png
└── img
    ├── azarashi.jpg
    └── tanuki.jpg
```