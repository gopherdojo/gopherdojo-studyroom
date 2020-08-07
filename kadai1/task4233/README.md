# Eimg - image encoder(kadai1)
## 画像変換コマンドのコマンド作成(eimg: Encode IMaGe)
### ディレクトリを指定可能 + 指定したディレクトリ以下のjpgファイルをpngに変換（デフォルト）
`img`ディレクトリを指定してjpg->pngの変換
```
$ mkdir -p img && convert -size 128x128 xc:blue blue.jpeg && mv blue.jpeg ./img
$ ./eimg img/
$ file img/blue.png 
img/blue.png: PNG image data, 128 x 128, 16-bit/color RGB, non-interlaced
```

### ディレクトリ以下は再帰的に処理する
`img`ディレクトリを指定して再帰的に処理
```
$ mkdir -p img/img/img/img/img/img && convert -size 128x128 xc:blue blue.jpeg && mv blue.jpeg ./img/img/img/img/img
$ ./eimg img/
$ find .
.
./eimg
./img
./img/img
./img/img/img
./img/img/img/img
./img/img/img/img/img
./img/img/img/img/img/img
./img/img/img/img/img/blue.png
$ file ./img/img/img/img/img/blue.png
./img/img/img/img/img/blue.png: PNG image data, 128 x 128, 16-bit/color RGB, non-interlaced
```

### 変換前と変換後の画像形式を指定可能（オプショナル）
変換前と変換後の画像形式を指定して処理
```
$ ./eimg -f=png -t=gif .
$ find .
.
./eimg
./img
./img/img
./img/img/img
./img/img/img/img
./img/img/img/img/img
./img/img/img/img/img/img
./img/img/img/img/img/blue.gif
$ file ./img/img/img/img/img/blue.gif
./img/img/img/img/img/blue.gif: GIF image data, version 89a, 128 x 128
```

## 要件
 - mainパッケージと分離する
   - eimgパッケージを作りました
 - 自作パッケージと標準パッケージと準標準パッケージのみ使う
   - golang.org/x以下のパッケージのこと
 - ユーザ定義型を作る
   - [Error型](https://github.com/task4233/gopherdojo-studyroom/blob/kadai1-task4233/kadai1/task4233/eimg/errors.go#L31-L41)
 - GoDocを作成する
   - [ここ](https://task4233.github.io/gopherdojo-studyroom/)に自動生成されるようにCD組みました
 - Go Modulesを使ってみる
   - 使いました
 


## Description
Package eimg encodes image files with set directory recursively.  
Both before and after extension can be specified.  
Default setting is `.jpg/jpeg` -> `.png`

Convert the images under the `./img` from `gif` to `jpg`.
```
$ ./eimg ./img -f=gif -t=jpg
```

## Build
```
$ git clone https://github.com/task4233/gopherdojo-studyroom.git
$ git switch kadai1-task4233
$ cd kadai1/task4233/eimg/cmd
$ go build -o ./eimg
```

## Options
```
Usage of ./eimg:
  -f string
    	file extension before executing (default "jpeg")
  -t string
    	file extension after executing (default "png")
```

## Executable(pre-release)
Download for Linux from [here](https://github.com/task4233/gopherdojo-studyroom/releases/tag/0.0.1).

## Author
[task4233](https://github.com/task4233)

