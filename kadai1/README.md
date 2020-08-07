# Eimg - image encoder(kadai1)


## Description
Package eimg encodes image files with set directory recursively.
Both before and after extension can be specified.
Default setting is `.jpg/jpeg` -> `.png`

Convert the images under the `./img` from `gif` to `jpg`.
```
$ eimg ./img -f=gif -t=jpg
```

## Options
```
Options:
-f    file extension before executing
-t    file extension after executing
```

## Author
[task4233](https://github.com/task4233)

## 画像変換コマンドのコマンド作成(eimg: Encode IMaGe)
 - ディレクトリを指定可能
 - 指定したディレクトリ以下のjpgファイルをpngに変換（デフォルト）
 - ディレクトリ以下は再帰的に処理する
 - 変換前と変換後の画像形式を指定可能（オプショナル）

## 要件
 - mainパッケージと分離する
   - eimgパッケージを作りました
 - 自作パッケージと標準パッケージと準標準パッケージのみ使う
   - golang.org/x以下のパッケージのこと
 - ユーザ定義型を作る
   - [Error型](https://github.com/task4233/gopherdojo-studyroom/blob/kadai1-task4233/kadai1/eimg/errors.go#L31-L41)
 - GoDocを作成する
   - [ここ](https://task4233.github.io/gopherdojo-studyroom/)に自動生成されるようにCD組みました
 - Go Modulesを使ってみる
   - 使いました
 
