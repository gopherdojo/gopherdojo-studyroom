# 課題1

## 画像変換コマンドを作成する

- 次の仕様を満たすコマンドを作って下さい
    - ディレクトリを指定する
    - 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
    - ディレクトリ以下は再帰的に処理する
    - 変換前と変換後の画像形式を指定できる（オプション）
- 以下を満たすように開発してください
    - mainパッケージと分離する
    - 自作パッケージと標準パッケージと準標準パッケージのみ使う
    - 準標準パッケージ：golang.org/x 以下のパッケージ
    - ユーザ定義型を作ってみる
    - GoDocを生成してみる
    - Go Modulesを使ってみる

## usage

```
.bin/imgconv -h
Usage of .bin/imgconv:
  -input-type string
        input type[jpg|jpeg|png|gif] (default "jpg")
  -output-type string
        output type[jpg|jpeg|png|gif] (default "png")
```

基本的なコマンドは `Makefile` で利用できます。

### build

ビルドすると `bin` フォルダに `imgconv` のバイナリが生成されます。


```
make build
```

### run

`testdata` ディレクトリ内にあるPNG画像ファイルをJPG画像ファイルに変換します。

```
make run
```

### test

`testdata` にテスト用の画像を生成してから、 `make run` を実行します。
画像データを用いた実行を行うだけで、テストを行っているわけではありません。

```
make test
```
