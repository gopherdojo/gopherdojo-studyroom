# 仕様
- ディレクトリを指定する
- 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
- ディレクトリ以下は再帰的に処理する
- 変換前と変換後の画像形式を指定できる（オプション）


# 開発条件
- mainパッケージと分離する
  - 変換ロジックをimgconvパッケージに分離
- 自作パッケージと標準パッケージと準標準パッケージのみ使う
- ユーザ定義型を作ってみる
  - imgconvパッケージにて、`Converter`型を使用
- GoDocを生成してみる
- Go Modulesを使ってみる


# オプション
|オプション|説明|デフォルト|対応
|:---:|:---:|:---:|:---:|
|-s|変換前の拡張子を指定|jpg|jpg・png|
|-d|変換後の拡張子を指定|png|jpg・png|


# 使い方
```shell
$ cd {Path_To_Repository}/kadai1
$ ./testdata.zsh    # テストデータ生成&初期化
$ go build -o exec  # ビルド
```

```shell
$ # 単体ファイル(jpg->png)
$ ./exec ./testdata/azarashi.jpg
$ # 単体ファイル(png->jpg)
$ ./exec -s png -d jpg ./testdata/osaru.png
$ # ディレクトリ(jpg->png)
$ ./exec ./testdata/ 
$ # ディレクトリ(png->jpg)
$ ./exec -s png -d jpg ./testdata
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
