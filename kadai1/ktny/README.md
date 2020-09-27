# Gopher道場 課題1
画像変換コマンドを作ろう

## 仕様
- ディレクトリを指定する
- 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
- ディレクトリ以下は再帰的に処理する
- 変換前と変換後の画像形式を指定できる（オプション）

## 制限/条件
- mainパッケージと分離する
- 自作パッケージと標準パッケージと準標準パッケージのみ使う
- 準標準パッケージ：golang.org/x以下のパッケージ
- ユーザ定義型を作ってみる
- GoDocを生成してみる
- Go Modulesを使ってみる

## ビルド
```
go build -o convimg cmd/convimg/main.go
```

## 使い方
```
./convimg [options] dir

dir
    target directory path

options
    -from string
        from ext. support jpg, jpeg, png, gif. (default "jpg")
    -to string
        to ext. support jpg, jpeg, png, gif. (default "png")
```

## 例
testdata配下のjpg画像をpngに変換する
`./convimg testdata`

testdata/child配下のpng画像をgifに変換する
`./convimg --from=png --to=gif testdata/child`
