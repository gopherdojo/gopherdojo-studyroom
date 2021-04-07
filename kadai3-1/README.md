# 課題3-1 タイピングゲームを作ろう

## 機能
- 標準出力にGoの標準パッケージ名がランダムで10問でます。
- 標準入力から1行受け取り、正しいかどうか確認します。
- 制限時間（10秒）で何問解けたか表示します。

## 使い方
`go run typingGame.go` でゲームスタートです！

```
$ go run typingGame.go
os
> os
OK :)
bytes
> bytes
OK :)
crypto
> crypto
OK :)
encoding
> endocing
KO :(
unsafe
> unsafe
OK :)
expvar
> exp
Times up!!
Score: 4/10: KO :(
```
