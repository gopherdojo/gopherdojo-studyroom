## io.Reader / io.Writerとは
io.Readerとio.Writerとは、バイト配列の読み書きを行うReadおよびWriteメソッドをもつ構造体に適用されるインターフェース型である。

## 標準パッケージでの使用例
標準パッケージ内での、io.Readerの使用例として、net/httpパッケージの、ClientタイプのPostメソッドが挙げられる。POSTメソッドは、以下のように定義されている。

```go
func (c *Client) Post(url, contentType string, body io.Reader) (resp *Response, err error)
```

ここでは、リクエストボディを指定する第三引数の型が```io.Reader```となっている。
第三引数が```io.Reader```となっているため、```Read```メソッドを持つあらゆる構造体を、リクエストボディに指定することができる。

例えば、```os.Open```の一つ目の返り値である、```File```タイプは```Read```メソッドを持つ。

```go
var r io.Reader
var err error
r, err = os.Open("./body.json")
```

そのため、上のように記述することで変数```r```を、そのまま```Post```メソッドの第三引数に指定することができる。

もし、文字列をリクエストボディとしたい場合、以下のようになる。

```go
var r io.Reader
r = strings.NewReader("My request")
```

このように、```io.Reader```を使って引数を定義していることで、必要に応じて様々な型を引数にあたえることが可能となっている。

## 利点
io.Readerとio.Wirterの利点は、入出力にまつわる型の振る舞いを、一括して取り扱うことができる点にある。