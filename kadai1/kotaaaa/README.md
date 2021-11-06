# Specification
- 次の仕様を満たすコマンドを作って下さい
  - ディレクトリを指定する
  - 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
  - ディレクトリ以下は再帰的に処理する
  - 変換前と変換後の画像形式を指定できる（オプション）
- 以下を満たすように開発してください
  - mainパッケージと分離する
  - 自作パッケージと標準パッケージと準標準パッケージのみ使う
  - 準標準パッケージ：golang.org/x以下のパッケージ
  - ユーザ定義型を作ってみる
  - GoDocを生成してみる
  - Go Modulesを使ってみる

# How to use
```
$ pwd
(YOUR_PATH)/gopherdojo-studyroom/kadai1/kotaaaa
$ go build -o converter
$ ./converter -path="./testdata/" -srcExt=".jpg" -dstExt=".png"
```

# How to test
```
$ go test ./... --count=1 -cover
?       github.com/kotaaaa/gopherdojo-studyroom/kadai1/kotaaaa  [no test files]
ok      github.com/kotaaaa/gopherdojo-studyroom/kadai1/kotaaaa/convert  8.231s  coverage: 70.4% of statements
ok      github.com/kotaaaa/gopherdojo-studyroom/kadai1/kotaaaa/search   0.005s  coverage: 91.7% of statements
ok      github.com/kotaaaa/gopherdojo-studyroom/kadai1/kotaaaa/validator        0.006s  coverage: 100.0% of statements
```

# Notes 
### Extensions that you can convert to.
- .gif, .jpeg, .jpg, .png

## Help
```
$./converter --help
Usage of /var/folders/nx/xqljz2y954qbyppfwn4w0tcr0000gn/T/go-build027276676/b001/exe/main:
  -dstExt string
        変換後の拡張子 (default ".png")
  -path string
        ファイルパス
  -srcExt string
        変換前の拡張子 (default ".jpg")
exit status 2
```

