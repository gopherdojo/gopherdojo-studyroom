# Kadai1 (Kadai2 のテスト込み)

## 使い方

```
go run main.go [-from=ext] [-to=ext] /path/to/dorectory
# ext = jpg | jpeg | png | gif
```

### 例
* jpg => png [default]

  ```
  $ go run main.go /path/to/directory
  ```

* png => gif

  ```
  $ go run main.go -from=png -to=gif /path/to/directory
  ```

