## 課題2

### 【TRY】io.Readerとio.Writer

- io.Readerとio.Writerについて調べてみよう
  - 標準パッケージでどのように使われているか
  - io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる

#### 回答

- 標準パッケージでどのように使われているか
  - [bufio.Reader](https://golang.org/pkg/bufio/#Reader.Read)や[os.File](https://golang.org/pkg/os/#File.Write)で使われている
  - それぞれのインターフェースを満たすために、上記の標準パッケージの構造体でメソッドが実装されている
- io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる
  - [net.Conn](https://golang.org/pkg/net/#Conn)のように、ハイブリッドなインターフェースを作成することができ、書き込み先がなんであろうと`Write()`メソッドを使って書き出すことができる
    - サーバーにリクエストを送信することも可能

#### 参考

[Goならわかるシステムプログラミング](https://www.lambdanote.com/products/go)

### 【TRY】テストを書いてみよう

[前回のPull Request](https://github.com/gopherdojo/gopherdojo-studyroom/pull/34)にテストもまとめて書きましたので、そちらを参照頂けますと幸いです。
