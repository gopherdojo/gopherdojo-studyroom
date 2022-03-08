
# io.Readerとio.Writerについて調べてみよう

## 標準パッケージでどのように使われているか
各パッケージは標準ライブラリ内で以下のように定義されている
```go
// io.Reader
type Reader interface {
    Read(p []byte) (n int, err error)
}
```
→ パッケージを利用する対象をデータとして格納する際に用いられる

```go
// io.Writer 
type Writer interface {
	Write(p []byte) (n int, err error)
}
```
→ パッケージを利用することで、その対象にデータを追加することができる

上記の理由により、ファイルの読み込みなどI/Oが関係するパッケージの一部に基本組みこまれているものとなっている

## io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる

例として、ファイルの読み書きをする場合を考えると、それぞれ以下の通りとなる。

1. ファイル（text.txt）を読み込む場合
```go
f, err := os.Open("text.txt")
data := make([]byte, 1024)
count, err := f.Read(data)
```
→ ファイルの内容を変数dataに格納する

2. ファイル（text.txt）にデータを書き込む場合
```go
f, err := os.Create("text.txt")
data := []byte("データの中身")
count, err := f.Write(data)
```
→ 「データの中身」という文字列をファイルに書き込む

