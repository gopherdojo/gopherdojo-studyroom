# 課題1 画像変換コマンドの作成

## 仕様

1. ディレクトリを指定する
    * input, output 両方指定可能
2. 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
    * flag のデフォルトに指定
3. ディレクトリ以下は再帰的に処理する
    * filepath.Walkを使用
4. 変換前と変換後の画像形式を指定できる
    * jpg, jpeg, gif, png に対応（ただし gif への変換は画像がバグる。原因調査中）
5. mainパッケージと分離する
    * main と imgconv に分離
6. 自作パッケージと標準パッケージと準標準パッケージのみ使う
    * 指定以外のパッケージは未使用
7. ユーザ定義型を作ってみる
    * target, ext を作成
8. GoDocを生成してみる
    * 起動を確認
9. Go Modulesを使ってみる
    * go mod でプロジェクトを生成

## 使い方
```
$ make build
$ ./bin/imgconv -id <input dir>, -od <output dir>, -ie <input ext>, -oe <output ext>
```
