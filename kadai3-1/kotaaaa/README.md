# Typing Game
- Receive one line from standard input
- Display English words on the standard output (you can choose what to display) - Receive one line from the standard input
- Display English words on standard output (any output) - Receive one line from standard input - Display how many questions were solved within the time limit
- (Original Function) Display word's meaning of target English word.

# How to use
```
$ go build -o game
$ ./game 
```

# How to test
```
$ go test ./... --count=1 -cover
?       github.com/kotaaaa/gopherdojo-studyroom/kadai3-1/kotaaaa        [no test files]
ok      github.com/kotaaaa/gopherdojo-studyroom/kadai3-1/kotaaaa/questions      0.010s  coverage: 100.0% of statements
ok      github.com/kotaaaa/gopherdojo-studyroom/kadai3-1/kotaaaa/starter        0.005s  coverage: 69.2% of statements
```

# Notes 
- Word source: ielts-4000-academic-word
  - https://tuxdoc.com/download/ielts-4000-academic-word-listpdf-4_pdf
  - Vocabulary file format  
    - "word:meaning"
