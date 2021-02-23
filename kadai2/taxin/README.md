## 課題2
### 課題の要件(Requirements)
- 次の仕様を満たすコマンドを作成する
    - [X] 課題1のテストコードを作成する

- 下記の要件を満たすように作成する
    - [X] テストのしやすさを考えてリファクタリングする
        - 複数の処理のまとまりではなく、個別の処理ごとにテストできるように関数に分割する
        - 参照透過性を持った関数として実装する
    - [X] テストのカバレッジを取る 
        - 各パッケージごとのテストカバレッジ(C0)は下記の通り
            - main package: 66.7%
            - converter poackage: 82.7%
    ```bash
    ~/g/s/g/t/g/k/taxin ❯❯❯ make test
    go test ./... -v -cover
    === RUN   TestValidateArgs
    === RUN   TestValidateArgs/case1
    ...
        --- PASS: TestValidateArgs/case5 (0.00s)
    PASS
    coverage: 66.7% of statements
    ok      github.com/taxintt/gopherdojo-studyroom/kadai2/taxin    0.082s  coverage: 66.7% of statements

    === RUN   TestFilePathConvert
    === RUN   TestFilePathConvert/case1
    ...
        --- PASS: TestConvertOtherKindsOfFiles/case1 (0.00s)
        --- PASS: TestConvertOtherKindsOfFiles/case2 (0.00s)
    PASS
    coverage: 82.7% of statements
    ok      github.com/taxintt/gopherdojo-studyroom/kadai2/taxin/converter  0.666s  coverage: 82.7% of statements
    ```
    - [X] テーブル駆動テストを行う
    - [X] テストヘルパーを実装する

