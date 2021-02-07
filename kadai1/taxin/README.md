## 課題1
### 課題の要件(Requirements)
- 次の仕様を満たすコマンドを作成する
    - [X] ディレクトリを指定する
    - [X] 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
        - [X] ディレクトリ以下は再帰的に処理する
    - [X] 変換前と変換後の画像形式を指定できる（オプション）

- 下記の要件を満たすように作成する
    - [X] mainパッケージと分離する
    - [X] 自作パッケージと標準パッケージと準標準パッケージ(golang.org/x以下のパッケージ)のみ使う
    - [X] ユーザ定義型を作ってみる
    - [X] GoDocを生成してみる
    - [X] Go Modulesを使ってみる

### CLIツールの利用方法(Usage)
```
$ ls -la ./testdata/testdir
total 12352
drwxr-xr-x  4 nishikawatakushi  staff      128 Feb  7 12:55 .
drwxr-xr-x  5 nishikawatakushi  staff      160 Feb  7 10:41 ..
-rw-r--r--@ 1 nishikawatakushi  staff  5669621 Jan 30 23:44 raphael-renter-csae9W8JAsw-unsplash.jpg

$ make build
go build -o bin/imgconverter ./main.go

$ ./bin/imgconverter -h                                              
Usage of ./bin/imgconverter:
  -d string
        画像が配置されているディレクトリのパス (default ".")
  -f string
        画像ファイルの変更前のフォーマット(jpg/jpeg/png/gif) (default "jpg")
  -o string
        画像ファイルの変更後のフォーマット(jpg/jpeg/png/gif) (default "png")


$ ./bin/imgconverter testdata -f jpg -o png
Start to covert image file...
Finished to covert image file
```
