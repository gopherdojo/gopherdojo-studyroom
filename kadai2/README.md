# io.Reader と io.Writer について

## とは？

io.Reader: io パッケージで定義されているインターフェイス。以下のような Read 関数を実装している型を抽象化して扱うことができる。

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

io.Writer も同様に以下のような Write 関数を実装している型を抽象化して扱うことができる。

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

## 標準パッケージでどう使われている??

### 例. os.MultiReader

複数の Reader を一つの io.Reader にまとめることができる。

```
func MultiReader(readers ...Reader) Reader
```

### 何が嬉しい？

io.Reader を実装している型ならなんでもまとめられる。

- os.Stdin
- os.File
- bytes.Buffer

なんでも OK。読み取り先を気にしないで扱える。

## 具体例 mycat

オプションなしのシンプルな cat コマンド。

- コマンドライン引数がない場合：標準入力から読み取る
- コマンドライン引数がある場合：引数のパスのファイルの内容を読み取る。複数 OK。
- io.Readerを使い抽象化することで、読み取り→表示の実装では読み取り先を意識しなくてよくなる

ソース（import、エラー処理等は省略。全体は mycat/mycat.go 参照）

```go
func main() {
	flag.Parse()

	// 読み取り元をio.Readerに代入する
	var reader io.Reader
	if len(flag.Args()) == 0 {
		reader = os.Stdin // 引数がない場合は標準入力から読み取る
	} else {
		for i := 0; i < len(flag.Args()); i++ {
			fs, err := os.Open(flag.Args()[i]) // 引数がある場合はファイルから読み取る
			defer fs.Close()
			if i == 0 {
				reader = fs
			} else {
				reader = io.MultiReader(reader, fs) // fsをio.Readerに抽象化することで、一つのreaderにまとめることができる
			}
		}
	}

	// io.Readerで抽象化しているのでここから先は読み取り元を意識しなくてもよい
	buf := make([]byte, 128)
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}
```

# 画像変換コマンド
[Gopher道場](https://gopherdojo.org/)の課題1です。

# 機能
コマンドライン引数に指定したディレクトリ以下の画像ファイルを変換します。

## 対応ファイル形式

- jpg -> png
- png -> jpg

※ その他の形式のファイル/ディレクトリは無視されます。

# 使い方
## コンパイル

```
go build convert.go
```

## 実行
jpgからpng

```
./convert directoryname -i=jpg -o=png
```

pngからjpg

```
./convert directoryname -i=png -o=jpg
```

## オプション

`-i`: 入力画像形式（デフォルト: png）
`-o`: 出力画像形式（デフォルト: jpg）
