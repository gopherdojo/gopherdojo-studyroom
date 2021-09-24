# Image Converter

## Options

|option|default value|description|
|---|---|---|
|`-src-ext`|`.jpg`|Source image extension. `.jpg`, `.jpeg`, `.png`, `.gif`, `.tif`, `.tiff`, `.bmp`|
|`-dst-ext`|`.png`|Destination image extension. `.jpg`, `.jpeg`, `.png`, `.gif`, `.tif`, `.tiff`, `.bmp`|
|`-src-dir`|`./testdata/src`|Source directory|
|`-dst-dir`|`./testdata/dst`|Destination directory|
|`-delete`|`false`|Whether to delete source images after converting|

## Usage

1. Put your image files in `./testdata/src`

2. Convert image files
```shell
$ go run main.go -delete=true -src-ext=.png -dst-ext=.gif
```
or
```shell
$ go build main.go
$ ./main -delete=true -src-ext=.png -dst-ext=.gif
```

3. Confirm result
```shell
$ ls -al ./testdata/dst
```
