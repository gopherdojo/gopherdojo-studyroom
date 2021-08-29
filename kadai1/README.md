# 課題 1:画像変換コマンドを作ろう

画像を指定の形式に変換する
変換後の画像ファイルは元の画像と同じディレクトリに配置される

## コマンドラインオプション

| オプション | 値の説明                               | デフォルト値 |
| ---------- | -------------------------------------- | ------------ |
| -from       | 変換前の画像形式(jpg,jpeg,pngから選択) | jpg          |
| -to      | 変換後の画像形式(jpg,jpeg,pngから選択) | png          |
| -dir | ディレクトリの指定 | ./ |

## 使い方

- ビルド

```bash
go build -o image-convert
```

- ディレクトリを指定して実行

```bash
./image-convert [...options] [directory]
```
```example
./image-convert -from png -to jpg -dir d1
```