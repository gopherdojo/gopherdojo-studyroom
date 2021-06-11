## 課題2 テストを書いてみよう


### 1回目の課題のテストを作る
- テストのしやすさを考えてリファクタリングしてみる
- テストのカバレッジを取ってみる
- テーブル駆動テストを行う
- テストヘルパーを作ってみる


### コマンドラインオプション

 | オプション | 説明 | デフォルト |
 | --- | --- | --- |
 | -path | 変換したい画像があるディレクトリ| カレントディレクトリ |
 | -pre | 変換前フォーマット | jpeg |
 | -post | 変換後フォーマット | png |


### 対応している画像フォーマット
- png
- jpeg(jpg)
- gif


### 使い方
1. バイナリビルド（実行ファイル作成）
```bash
$ make
```
2. ディレクトリを指定して実行
```bash
$ ./bin/converter [options...]
```


### テストの方法
- バイナリビルド & テスト
```bash
$ make test
```
- テスト後の処理(掃除)
```bash
$ make clean
```

### テスト結果
```bash
=== RUN   TestIsSupportedFormat
=== RUN   TestIsSupportedFormat/jpg
=== RUN   TestIsSupportedFormat/jpeg
=== RUN   TestIsSupportedFormat/png
=== RUN   TestIsSupportedFormat/gif
=== RUN   TestIsSupportedFormat/xls
--- PASS: TestIsSupportedFormat (0.00s)
    --- PASS: TestIsSupportedFormat/jpg (0.00s)
    --- PASS: TestIsSupportedFormat/jpeg (0.00s)
    --- PASS: TestIsSupportedFormat/png (0.00s)
    --- PASS: TestIsSupportedFormat/gif (0.00s)
    --- PASS: TestIsSupportedFormat/xls (0.00s)
=== RUN   TestValidate
=== RUN   TestValidate/jpg_and_png
=== RUN   TestValidate/jpeg_and_png
=== RUN   TestValidate/jpg_and_gif
=== RUN   TestValidate/jpeg_and_gif
=== RUN   TestValidate/png_and_jpg
=== RUN   TestValidate/png_and_jpeg
=== RUN   TestValidate/png_and_gif
=== RUN   TestValidate/gif_and_jpeg
=== RUN   TestValidate/gif_and_jpg
=== RUN   TestValidate/gif_and_png
=== RUN   TestValidate/xls_and_gif
=== RUN   TestValidate/gif_and_xls
=== RUN   TestValidate/jpg_and_jpeg
=== RUN   TestValidate/jpeg_and_jpg
=== RUN   TestValidate/png_and_png
=== RUN   TestValidate/gif_and_gif
--- PASS: TestValidate (0.00s)
    --- PASS: TestValidate/jpg_and_png (0.00s)
    --- PASS: TestValidate/jpeg_and_png (0.00s)
    --- PASS: TestValidate/jpg_and_gif (0.00s)
    --- PASS: TestValidate/jpeg_and_gif (0.00s)
    --- PASS: TestValidate/png_and_jpg (0.00s)
    --- PASS: TestValidate/png_and_jpeg (0.00s)
    --- PASS: TestValidate/png_and_gif (0.00s)
    --- PASS: TestValidate/gif_and_jpeg (0.00s)
    --- PASS: TestValidate/gif_and_jpg (0.00s)
    --- PASS: TestValidate/gif_and_png (0.00s)
    --- PASS: TestValidate/xls_and_gif (0.00s)
    --- PASS: TestValidate/gif_and_xls (0.00s)
    --- PASS: TestValidate/jpg_and_jpeg (0.00s)
    --- PASS: TestValidate/jpeg_and_jpg (0.00s)
    --- PASS: TestValidate/png_and_png (0.00s)
    --- PASS: TestValidate/gif_and_gif (0.00s)
PASS
coverage: 73.1% of statements
ok  	github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai2/Mizushima	0.003s

=== RUN   TestGlob
=== RUN   TestGlob/jpeg_files
=== RUN   TestGlob/png_files
=== RUN   TestGlob/gif_files
=== RUN   TestGlob/xls_files
--- PASS: TestGlob (0.00s)
    --- PASS: TestGlob/jpeg_files (0.00s)
    --- PASS: TestGlob/png_files (0.00s)
    --- PASS: TestGlob/gif_files (0.00s)
    --- PASS: TestGlob/xls_files (0.00s)
=== RUN   TestPicConverter_Conv
=== RUN   TestPicConverter_Conv/pre:_jpeg,_after:_png
=== RUN   TestPicConverter_Conv/pre:_jpeg,_after:_gif
=== RUN   TestPicConverter_Conv/pre:_jpg,_after:_png
=== RUN   TestPicConverter_Conv/pre:_jpg,_after:_gif
=== RUN   TestPicConverter_Conv/pre:_png,_after:_jpg
=== RUN   TestPicConverter_Conv/pre:_png,_after:_gif
=== RUN   TestPicConverter_Conv/pre:_gif,_after:_jpg
=== RUN   TestPicConverter_Conv/pre:_gif,_after:_png
=== RUN   TestPicConverter_Conv/pre:_jpg,_after:_xls
=== RUN   TestPicConverter_Conv/no_file
--- PASS: TestPicConverter_Conv (0.08s)
    --- PASS: TestPicConverter_Conv/pre:_jpeg,_after:_png (0.00s)
    --- PASS: TestPicConverter_Conv/pre:_jpeg,_after:_gif (0.01s)
    --- PASS: TestPicConverter_Conv/pre:_jpg,_after:_png (0.00s)
    --- PASS: TestPicConverter_Conv/pre:_jpg,_after:_gif (0.01s)
    --- PASS: TestPicConverter_Conv/pre:_png,_after:_jpg (0.00s)
    --- PASS: TestPicConverter_Conv/pre:_png,_after:_gif (0.01s)
    --- PASS: TestPicConverter_Conv/pre:_gif,_after:_jpg (0.00s)
    --- PASS: TestPicConverter_Conv/pre:_gif,_after:_png (0.05s)
    --- PASS: TestPicConverter_Conv/pre:_jpg,_after:_xls (0.00s)
    --- PASS: TestPicConverter_Conv/no_file (0.00s)
=== RUN   TestNewPicConverter
=== RUN   TestNewPicConverter/pre:_jpg,_post:_png
=== RUN   TestNewPicConverter/pre:_jpeg,_post:_png
=== RUN   TestNewPicConverter/pre:_jpg,_post:_gif
=== RUN   TestNewPicConverter/pre:_png,_post:_jpeg
=== RUN   TestNewPicConverter/pre:_png,_post:_jpg
=== RUN   TestNewPicConverter/pre:_png,_post:_gif
=== RUN   TestNewPicConverter/pre:_gif,_post:_jpeg
=== RUN   TestNewPicConverter/pre:_gif,_post:_jpg
=== RUN   TestNewPicConverter/pre:_gif,_post:_png
--- PASS: TestNewPicConverter (0.00s)
    --- PASS: TestNewPicConverter/pre:_jpg,_post:_png (0.00s)
    --- PASS: TestNewPicConverter/pre:_jpeg,_post:_png (0.00s)
    --- PASS: TestNewPicConverter/pre:_jpg,_post:_gif (0.00s)
    --- PASS: TestNewPicConverter/pre:_png,_post:_jpeg (0.00s)
    --- PASS: TestNewPicConverter/pre:_png,_post:_jpg (0.00s)
    --- PASS: TestNewPicConverter/pre:_png,_post:_gif (0.00s)
    --- PASS: TestNewPicConverter/pre:_gif,_post:_jpeg (0.00s)
    --- PASS: TestNewPicConverter/pre:_gif,_post:_jpg (0.00s)
    --- PASS: TestNewPicConverter/pre:_gif,_post:_png (0.00s)
PASS
coverage: 87.8% of statements
ok  	github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai2/Mizushima/picconvert	0.087s

```
