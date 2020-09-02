# [TRY] io.Reader と io.Writer

## ・ io.Reader と io.Writer について調べてみよう

・標準パッケージでどのように使われているか？
・io.Reader と io.Writer があることで
どういう利点があるのか具体例を挙げて考えてみる

io.Writer,io.Reader はそれぞれ

```
type Writer interface {
    Write(p []byte) (n int, err error)
}

type Reader interface {
    Read(p []byte) (n int, err error)
}
```

と定義されている。
「バイト列 b を書き（読み）込み、書き（読み）込んだバイト数 n と、エラーが起きた場合はそのエラー error を返す」というデータの入出力を抽象化する interface である。

これにより、異なるデータ型も、あらかじめ宣言されてる interface([]byte 型のデータを引数として受け取り、それぞれ write,read メソッドをもっている)を満たしていれば、同じように扱える

標準パッケージでは encoding や ioutil などで使用されている。
