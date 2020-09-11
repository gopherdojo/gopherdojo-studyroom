# 画像変換コマンドを作ろう

## 仕様

- ディレクトリを指定する
- 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
- ディレクトリ以下は再帰的に処理する
- 変換前と変換後の画像形式を指定できる（オプション）

## 開発条件

- mainパッケージと分離する
- 自作パッケージと標準パッケージと準標準パッケージのみ使う
  - 準標準パッケージ：golang.org/x以下のパッケージ
- ユーザ定義型を作ってみる
- GoDocを生成してみる
- Go Modulesを使ってみる

---

# 使用方法

```
使用方法:
  image_conversion [-from=<ext>] [-to=<ext>] target_directory
引数:
  -from=<ext> 変換前の拡張子 [gif jpeg jpg png bmp] (default: jpg)
  -to=<ext>   変換後の拡張子 [gif jpeg jpg png bmp] (default: png)
```

例）
```
❯ go build image_conversion.go

❯ ./image_conversion testdata
変換が完了しました

❯ ./image_conversion -from=png -to=bmp testdata
変換が完了しました

❯ ./image_conversion testdata testdata    // 引数が多い
使用方法:
  image_conversion [-from=<ext>] [-to=<ext>] target_directory
引数:
  -from=<ext> 変換前の拡張子 [gif jpeg jpg png bmp] (default: jpg)
  -to=<ext>   変換後の拡張子 [gif jpeg jpg png bmp] (default: png)
```

---

## 分からなかったこと

答えてもらえると嬉しいです。

1. go.mod に書かれているも module に github.com/gopherdojo-studyroom/kadai1/uesho と現在書いている。
  実際は何を書くのが正解なのか？ fork した場合には自分のリポジトリのURLを書いた方がいいのか？

2. GoDoc の使い方がいまいち分からなかった。godocコマンドで生成するって言っている記事もあれば、自動生成されると言っているものもある。
  実際に[go.dev](https://go.dev/about)に掲載するにはどうすればいいのか？

3. テストを使用と思ったが、同じパッケージなのに参照できなかった。
何故なのか？
また、プログラム引数を取る場合、どうすやってテストすれば良いのか。