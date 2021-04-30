# 分割ダウンローダ

## 仕様

- 分割ダウンロードを行う
  - Range アクセスを用いる
  - いくつかのゴルーチンでダウンロードしてマージする
  - エラー処理を工夫する
    - golang.org/x/sync/errgourp パッケージなどを使ってみる
  - キャンセルが発生した場合の実装を行う

## オプション

| オプション | 内容                                   | デフォルト |
| ---------- | -------------------------------------- | ---------- |
| -d         | ファイルをダウンロードするディレクトリ | $PWD       |
| -t         | タイムアウトするまでの時間（秒）       | 10         |

## 利用方法

### setup

```shell
$ make  # テスト & ビルド
```

### ダウンロードコマンドの例

```shell
$ ./pdownload https://blog.golang.org/gopher/header.jpg
$ # ディレクトリとタイムアウトまでの時間の指定
$ ./pdownload -d testdata/ -t 30 https://blog.golang.org/gopher/header.jpg
```

## ディレクトリ構造

```
.
├── Makefile
├── README.md
├── go.mod
├── go.sum
├── main.go
├── pdownload
├── pdownload.go
├── pdownload_test.go
├── requests.go
├── requests_test.go
├── termination
│   ├── termination.go
│   └── termination_test.go
├── testdata
│   ├── header.jpg
│   └── test_download
│       └── header.jpg
└── util.go
```

## 参考

[Code-Hex/pget](https://github.com/Code-Hex/pget)と[gopherdojo/dojo3#50](https://github.com/gopherdojo/dojo3/pull/50)を参考にさせていただきました。
