## io.Readerとio.Writer
### [io.Reader](https://pkg.go.dev/io#Reader)
- バイト列を読み出すためのインターフェース
- 与えられたバイトスライス p []byte を先頭から埋めていく
- 埋まったバイト数 n と、埋める過程で発生したエラー err を返す
- n < len(p) の場合、err != nil である可能性がある
- 組 (0, nil) を返すことは非推奨
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

### [io.Writer](https://pkg.go.dev/io#Writer)
```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

### 標準パッケージでどのように使われているか
- *os.File は io.Reader インターフェースを持っている
- http.Response の Body メンバが io.ReadCloser というインターフェースを持っていますがこれは io.Reader と io.Closer の複合的なインターフェース

### io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる


### 参考資料
- https://pkg.go.dev/io#Reader
- https://pkg.go.dev/io#Writer
- https://qiita.com/ktnyt/items/8ede94469ba8b1399b12

## テスト
- [x] テストのしやすさを考えてリファクタリングしてみる
- [x] テストのカバレッジを取ってみる
- [x] テーブル駆動テストを行う
- [x] テストヘルパーを作ってみる
