# 課題 1: imgconv

指定されたディレクトリ以下の画像ファイルを異なる形式に変換するコンソールアプリケーションです。

### Usage
```
imgconv [-s <src image type>] [-d <dest image type>] <target directory path>
```

`-s` で指定した形式の画像ファイルを、`-d` で指定した形式に変換します。

例えば、jpeg ファイルを png ファイルに変換すると、
指定されたディレクトリ以下の jpeg ファイル、`hoge.jpeg` に対し、
同じディレクトリにそれを　png ファイルに変換した `hoge.png` が生成されます。

オプションの画像形式で指定できるものは以下の通りです。

- `jpg`: jpeg ファイル
- `png`: png ファイル
- `gif`: gif ファイル

オプションを指定しない場合、jpeg ファイルを png ファイルに変換します。

### make コマンド

#### ビルド
`make build` でビルドでき、 `bin/imgconv` が出力されます。

#### 実行 
また、`make run` で `testdata/` にサンプルデータを生成し、それをターゲットとして実行することができます。 
