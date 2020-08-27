# 仕様  
・ディレクトリを指定する
・指定したディレクトリ以下の.jpgファイルを.pngに変換（デフォルト）
・ディレクトリ以下は再帰的に処理する
・変換前と変換後の画像形式を指定できる(オプション)    
・変換終了後、元のファイルを消去するかy/nでユーザーが任意で選択可

# オプション  
| オプション | 説明　| デフォルト | 対応　|
| :-------: | :----: | :-------: | :----: |
|   -f    |変換前の拡張子を指定 | .jpg | .jpg・.png・.gif |
|   -t    |変換後の拡張子を指定| .png | .jpg・.png・.gif |  

# 使い方
```
$ # .jpg -> .png (デフォルト)
$./converter ./img

$ # .png -> .jpg (拡張子指定)
$./converter -f .png -t .jpg ./img  

変換後
$ Would you want to delete the original fail? y/n
y ----->　削除
n ----->  そのまま

それ以外
$ Please enter again
$ Would you want to delete the original fail? y/n
.....繰り返し
```

# ファイル
・cli.go 処理全般  
・main.go コマンドライン引数の取得,cli.goの関数呼び出し

# ディレクトリ構成  
```
└── converter
    ├── cli.go
    ├── converter
    ├── go.mod
    ├── img
    │   ├── cat.png
    │   ├── sub
    │   │   └── dog.png
    │   └── tiger.png
    ├── img2
    │   └── monkey.png
    └── main.go

4 directories, 8 files
```
