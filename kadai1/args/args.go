package args

import "flag"

type CmdArgs struct {
	From string
	To   string
	Dir  string
}

func ParseArgs() *CmdArgs {
	var cmd CmdArgs
	flag.StringVar(&cmd.From, "from", "jpg", "from")
	flag.StringVar(&cmd.To, "to", "png", "to")
	flag.StringVar(&cmd.Dir, "dir", "./", "directory")
	flag.Parse()
	return &cmd
}
