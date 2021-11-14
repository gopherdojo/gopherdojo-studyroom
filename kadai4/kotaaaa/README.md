# Let's make a fortune API.
- Returning omikuji results in JSON format
- Only New Year's Day (Jan. 1 - Jan. 3) is set to "Daikichi".
- Writing tests for the handler

# How to use
```
$ go build -o fortune-api
$ ./fortune-api 
$ curl "http://localhost:8080/draw?p=2021-11-14"
{"status":"Success","result":"末吉"}
$ curl "http://localhost:8080/draw"
{"status":"Success","result":"小吉"}
```

# How to Test
```
$ go test ./... --count=1 -cover
?       github.com/kotaaaa/gopherdojo-studyroom/kadai4/kotaaaa  [no test files]
ok      github.com/kotaaaa/gopherdojo-studyroom/kadai4/kotaaaa/fortune  0.007s  coverage: 46.2% of statements
ok      github.com/kotaaaa/gopherdojo-studyroom/kadai4/kotaaaa/handler  0.013s  coverage: 73.1% of statements
$ go test ./...  -v
```
