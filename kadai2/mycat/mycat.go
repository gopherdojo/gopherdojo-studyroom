package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	flag.Parse()

	// 読み取り元をio.Readerに代入する
	var reader io.Reader
	if len(flag.Args()) == 0 {
		reader = os.Stdin // 引数がない場合は標準入力から読み取る
	} else {
		for i := 0; i < len(flag.Args()); i++ {
			fs, err := os.Open(flag.Args()[i]) // 引数がある場合はファイルから読み取る
			if err != nil {
				fmt.Fprintln(os.Stderr, "failed to read")
				continue
			}
			defer fs.Close()
			if i == 0 {
				reader = fs
			} else {
				reader = io.MultiReader(reader, fs) // fsをio.Readerに抽象化することで、一つのreaderにまとめることができる
			}

		}
	}

	// io.Readerで抽象化しているのでここから先は読み取り元を意識しなくてもよい
	buf := make([]byte, 128)
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintln(os.Stderr, "failed to read")
			os.Exit(1)
		}
		os.Stdout.Write(buf[:n])
	}
}
