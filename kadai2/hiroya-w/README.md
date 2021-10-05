# 課題2

## テストを書いてみよう

- [x] テストのしやすさを考えてリファクタリングしてみる
- [x] テストのカバレッジを取ってみる
- [x] テーブル駆動テストを行う
- [ ] テストヘルパーを作ってみる

今回のテストで、どの部分にテストヘルパーを利用出来るのかがわからなかった。他の方のPRを参考に眺めてみようと思う。
Goでオブジェクト指向をしようとしてハマるやつをやりかけてしまっているように感じたので、もう少しGoらしい書き方を勉強してもいいなと思った。

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

ビルドに必要なパッケージを取得します。

```
make devel-deps
```

ビルドすると `bin` フォルダに `imgconv` のバイナリが生成されます。

```
make build
```

### test

`testdata` にテスト用の画像を生成します。
その後、 `make test` でテストを実行します。

```
make test-deps
make test
```

### coverage

テストの実行後、カバレッジを表示します。

```
make cover
```

### document

ドキュメントを表示します。

```
make doc
```
