# 課題 1: imgconv

### Usage
```
imgconv [-s <src image type>] [-d <dest image type>] <target directory path>
```
オプションの画像形式で指定できるものは以下の通りです。

- `jpg`: jpeg ファイル
- `png`: png ファイル
- `gif`: gif ファイル


### make コマンド

#### ビルド
`make build` でビルドでき、 `bin/imgconv` が出力されます。

#### 実行 
また、`make run` で `testdata/` にサンプルデータを生成し、それをターゲットとして実行することができます。 
