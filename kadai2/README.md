# 課題2

## 1. io.Readerとio.Writerについて調べてみよう

### 標準パッケージでどのように使われているか
【io.Reader】
* io.Readerは、入力処理の基本インターフェース
* `Read(p []byte) (n int, err error)` をもつインターフェース
* バイト列を読み出すインターフェース

【io.Reader】
* io.Writerは、出力処理の基本インターフェース
* `Write(p []byte) (n int, err error)` をもつインターフェース
* バイト列を書き出すインターフェース

### io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる
* 様々な場面で出てくる入出力にio.Readerインターフェースに準拠しているため、汎用性に優れている。
* Read関数を実装している型を抽象化して扱うことができる。ファイル、文字列、メモリのデータ、ネットワーク接続情報等。
* 入力の値を明確に意識しなくても、シンプルに実装が可能。

[参考サイト](https://qiita.com/ktnyt/items/8ede94469ba8b1399b12)

## 2. テストを書いてみよう

### 1回目の課題のテストを作ってみて下さい
#### 課題要件
* テストのしやすさを考えてリファクタリングしてみる 
* テストのカバレッジを取ってみる 
* テーブル駆動テストを行う
* テストヘルパーを作ってみる

#### Usage
モジュール配下に移動
`cd convert`

テスト実行
`go test -run ''`

カバー内容を吐き出す
`go test -cover ./... -coverprofile=cover.out`

go toolを用いてcover.htmlを作成する
`go tool cover -html=cover.out -o cover.html`

cover.htmlを開く
`open cover.html`

#### カバレッジ
current : coverage: 75.0% of statements

