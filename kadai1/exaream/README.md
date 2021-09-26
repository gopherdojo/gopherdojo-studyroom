# Image Converter

## Overview
Image converter using Golang

## Options

|option|default value|description|
|---|---|---|
|`-src-ext`|`.jpg`|Source image extension `.jpg` `.jpeg` `.png` `.gif` `.tif` `.tiff` `.bmp`|
|`-dst-ext`|`.png`|Destination image extension `.jpg` `.jpeg` `.png` `.gif` `.tif` `.tiff` `.bmp`|
|`-src-dir`|`./testdata/src`|Source directory|
|`-dst-dir`|`./testdata/dst`|Destination directory|
|`-delete`|`false`|Whether to delete source images after converting|

### Notes
If you delete sample images using `-delete`, unit tests will not work properly.

## Usage

1. Clone this repository
```shell
$ git clone -b kadai1-exaream https://github.com/exaream/gopherdojo-studyroom.git
```

2. Put your image files in `./testdata/src`

3. Convert image files
Change the following options to match the extension of your image file.
```shell
$ go run main.go -src-ext=.png -dst-ext=.jpg
```
or
```shell
$ go build main.go
$ ./main -src-ext=.png -dst-ext=.jpg
```

4. Confirm the result
```shell
$ ls -al ./testdata/dst
```

## Directory structure

```
gopherdojo-studyroom/kadai1/exaream
├── .gitignore
├── README.md
├── fileutil
│   ├── fileutil.go
│   └── fileutil_test.go
├── go.mod
├── go.sum
├── imgconv
│   ├── converter.go
│   ├── converter_test.go
│   └── export_test.go
├── main.go
└── testdata
    ├── .gitkeep
    ├── dst
    │   └── .gitkeep
    └── src
        ├── .gitkeep
        ├── sample1.jpg
        ├── sample2.jpeg
        ├── sample3.png
        ├── sample4.gif
        ├── sample5.tif
        ├── sample6.tiff
        └── sample7.bmp
```

## TODO
* Add unit tests for Assignment 2.
* Find out a solution to an error that occurs when using `t.Parallel()` in a unit test with package `flag`.
