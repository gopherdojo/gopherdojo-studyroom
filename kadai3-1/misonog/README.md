# 課題 3-1 【TRY】タイピングゲームを作ろう

## 仕様

- 標準出力に英単語を出す（出すものは自由）
- 標準入力から 1 行受け取る
- 制限時間内に何問解けたか表示する

## 遊び方

```shell
$ make
$ ./typinggame
go
>go
+1

rust
>hoge
Wrong input

ruby
>
-----
Finish!
Result: 1 point.
```

## ディレクトリ構造

```
.
├── Makefile
├── README.md
├── go.mod
├── main.go
├── typegame.go
├── typegame_test.go
├── typing.go
├── typing_test.go
├── typinggame
└── word_list.txt
```

## 感想

`run()`と`input()`のよいテスト方法が思い浮かばなかった
