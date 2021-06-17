## 課題3-1 タイピングゲームを作ろう
- 標準出力に英単語を出す（出すものは自由）
- 標準入力から1行受け取る
- 制限時間内に何問解けたか表示する


### コマンドラインオプション

 | オプション | 説明 | デフォルト |
 | --- | --- | --- |
 | -limit | ゲーム全体の制限時間を秒数で設定する | 20 |


### インストール方法
```bash
go get github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai3-1/Mizushima
```

### 使い方
1. バイナリビルド（実行ファイル作成）
```bash
$ make
```
2. ディレクトリを指定して実行
```bash
$ ./bin/typing [option]
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

### ディレクトリ構成
```
.
├─ gamedata
│   └─ words.csv
├─ typing
│   └─ typing.go
│   └─ typing_test.go
├─ .gitignore
├─ go.mod
├─ main.go
├─ Makefile
└─ README.md
```