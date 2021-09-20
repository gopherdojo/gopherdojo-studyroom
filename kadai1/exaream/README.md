# Image Converter

## Options

|option|default value|description|
|---|---|---|
|`--src-ext`|`jpg`|変換元の画像の拡張子|
|`--dst-ext`|`png`|変換後の画像の拡張子|
|`--src-dir`|`./testdata/src`|変換元の対象ディレクトリ|
|`--dst-dir`|`./testdata/dst`|変換後の出力先ディレクトリ|
|`--delete`|`false`|変換元の画像の削除の有無|

### Valid extensions of options
* `jpg`, `jpeg`
* `png`
* `gif`

## Usage

1. Put your image files in `./testdata/src`

2. Convert image files
```shell
$ go run main.go --src-ext .png --dst-ext .gif
```
or
```shell
$ go build main.go
$ ./main --src-ext .png --dst-ext .gif
```

3. Confirm result
```shell
$ ls -al ./testdata/dst
```
