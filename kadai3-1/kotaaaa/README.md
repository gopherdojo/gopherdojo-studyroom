# Typing Game
- Receive one line from standard input
- Output English words to standard output (you can choose what to output) 
- Receive one line from standard input
- (Original Function) Display word's meaning of target English word.

# How to use
```
$ go build -o game
$ ./game 
( or $ go run .)
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


# How to play
```
$ go run .
Target:  generation
generation
Meaning:   all offspring at same stage from common ancestor; interval of time between the birth of parents and their offspring
Correct! 1pt 
Target:  centigrade
centigrade
Meaning:   measure of temperature, used widely in Europe
Correct! 2pt 
Target:  shield
shield
Meaning:   protective covering or structure; protect; guard
Correct! 3pt 
Target:  ultimate
ultimate
Meaning:   final; being the last or concluding; fundamental; elemental; extreme
Correct! 4pt 
Target:  numerous
aaaaa
Meaning:   many; various; amounting to a large indefinite number
Miss! 4pt Target:  network
rrrr
Meaning:   any system of lines or channels crossing like the fabric of a net; complex, interconnected group or system
Miss! 4pt Target:  admit

======================
Times up! point:  4 pt
======================

```
