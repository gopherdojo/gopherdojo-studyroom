# Image Converter

## Overview
Image converter using Golang

## Options

|option|default value|description|
|---|---|---|
|`-src-ext`|`jpg`|Source image extension `jpg` `jpeg` `png` `gif` `tif` `tiff` `bmp`|
|`-dst-ext`|`png`|Destination image extension `jpg` `jpeg` `png` `gif` `tif` `tiff` `bmp`|
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

2. Move to the project directory
```shell
$ cd gopherdojo-studyroom/kadai1/exaream
```

3. Put your image files in `./testdata/src`

4. Convert image files  
Change the following options to match the extension of your image file.
```shell
$ cd cmd/imgconv
$ go run main.go -src-ext=png -dst-ext=jpg
```
or
```shell
$ cd cmd/imgconv
$ go build main.go
$ ./main -src-ext=png -dst-ext=jpg
```

Confirm the result
```shell
$ ls -al ../../testdata/dst
```

## Directory structure

```
gopherdojo-studyroom/kadai1/exaream
├── cmd
│   └── imgconv
│       └── main.go
├── files
│   ├── files.go
│   └── files_test.go
├── imgconv
│   ├── converter.go
│   ├── converter_test.go
│   └── export_test.go
├── slices
│   ├── slices.go
│   └── slices_test.go
├── testdata
├── .gitignore
├── README.md
├── go.mod
└── go.sum
```
