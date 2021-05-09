# おみくじ API

## 仕様

- JSON 形式でおみくじの結果を返す
- 正月（1/1-1/3）だけ大吉にする
- ハンドラのテストを書いてみる

## 利用方法

### setup

```shell
$ make  # テスト & ビルド
```

### サーバーの起動

```shell
$ ./omikuji-server
```

### おみくじを引く

```shell
$ curl "http://127.0.0.1:8080/"
> {"result":"小吉"}
$ curl "http://127.0.0.1:8080/?date=2021-01-01"
> {"result":"大吉"}
```

## ディレクトリ構造

```shell
.
├── Makefile
├── README.md
├── go.mod
├── main.go
├── omikuji-server
├── omikuji.go
└── omikuji_test.go
```
