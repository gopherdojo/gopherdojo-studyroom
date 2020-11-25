# 課題
## 次の仕様を満たすコマンドを作って下さい
- [x] ディレクトリを指定する
- [x] 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
- [x] ディレクトリ以下は再帰的に処理する
- [x] 変換前と変換後の画像形式を指定できる（オプション）

## 以下を満たすように開発してください
- [x] mainパッケージと分離する
- [x] 自作パッケージと標準パッケージと準標準パッケージのみ使う
- [x] 準標準パッケージ：`golang.org/x`以下のパッケージ
- [x] ユーザ定義型を作ってみる
- [x] GoDocを生成してみる
- [x] Go Modulesを使ってみる

# 使い方
ツールは`go build -o bin/imgconv imgconv/main.go`でビルドすると、`bin/`配下に生成される
```
[options]
  -debug
        debug message flag. default value is false.
  -from .jpg
        extension before convert. default is .jpg (default ".jpg")
  -out ./output
        output directory. default is ./output (default "./output")
  -to .png
        extension after converted. default is .png (default ".png")
```