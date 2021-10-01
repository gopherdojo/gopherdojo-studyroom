package imgconv

import (
	"io"
)

type CLI struct {
	OutStream, ErrStream io.Writer
}

func (cli *CLI) Run() int {
	return 0
}
