### TODOリスト
順不同

- go routineによる分割ダウンロード処理
  - fileの大きさの事前取得
  - range access の range取得
- Makefile作成
  - テスト作成
  - ダミーサーバー

### 処理の流れ
- オプションからURL取得（複数可）一つ一つのURLにつき、
  - ~~Access-Range=bytesかどうか~~
    - ~~否なら、通常のダウンロード~~
  - ~~完成版ファイルを作成~~
  - ~~一時フォルダを作る~~
    - ~~並行してダウンロードし連番の一時ファイルに保存~~
  - ~~完成版フォルダに一時ファイルを順番にコピー~~
  - ~~一時フォルダ毎削除~~
  - 時間制限を設ける

### 参考にしたもの
[pget](https://qiita.com/codehex/items/d0a500ac387d39a34401)  (goroutineを使ったダウンロード処理、コマンドラインオプションの処理等々)  
https://github.com/gopherdojo/dojo3/pull/50  (ctrl+cを押したときのキャンセル処理など)