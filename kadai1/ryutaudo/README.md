# 課題1 
## 次の仕様を満たすコマンドを作って下さい
- [x] ディレクトリを指定する
- [x] 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
- [x] ディレクトリ以下は再帰的に処理する
- [x] 変換前と変換後の画像形式を指定できる（オプション）

## 以下を満たすように開発してください
- [x] mainパッケージと分離する
- [x] 自作パッケージと標準パッケージと準標準パッケージのみ使う
- [x] ユーザ定義型を作ってみる
- [x] GoDocを生成してみる (https://pkg.go.dev/github.com/ryutaudo/imgcvt)
- [x] Go Modulesを使ってみる 

# 使い方
### 1. バイナリを作成
```zsh
go build -o main
```

### 2. コマンドを使用
```zsh
./main -from={extension} -to={extension} {directory}
```

### オプション

| flag | Options |
|:--|:--|
|`-from`|`png`, `jpeg`, `jpg`, `gif`|
|`-to`|`png`, `jpeg`, `jpg`, `gif`|

### GoDoc
https://pkg.go.dev/github.com/ryutaudo/imgcvt
