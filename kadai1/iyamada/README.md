# 課題1 画像変換コマンドを作ろう

## Discription
指定したディレクトリ以下にある画像ファイルを変換するコマンドラインツールです。
ディレクトリがサブディレクトリを含む場合、サブディレクトリ以下の画像ファイルも変換します。
オプションで変換対象の画像形式を指定できます。対応している形式は、jpg、png、gifです。

## Build
makeすると`convert`という実行ファイルが作成されます。

``` sh
$ make
```

## Usage
オプションなしで実行すると、jpgファイルをpngファイルに変換します。
``` sh
$ ./convert [-io] [dir ...]
```

### Option
``` sh
$ ./convert -h
Usage of ./convert:
  -i string
        input file extension (default "jpg")
  -o string
        output file extension (default "png")
```

## Test
自作した画像変換パッケージをテストするには、Makefileがあるディレクトリで`make test`を実行してください。カバレッジは`imgconv/coverage.html`で確認できます。
``` sh
$ make test
$ ls imgconv/coverage.html
imgconv/coverage.html
```
