package options

import (
	"bytes"
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"
	errors "github.com/pkg/errors"
)

type Options struct {
	Help   bool   `short:"h" long:"help"`
	Procs  int    `short:"p" long:"procs"`
	Output string `short:"o" long:"output" default:"./"`
	Tm     int    `short:"t" long:"time" default:"3"`
}

func (opts *Options) parse(argv []string) ([]string, error) {
	p := flags.NewParser(opts, flags.PrintErrors)
	args, err := p.ParseArgs(argv)

	if err != nil {
		os.Stderr.Write(opts.usage())
		return nil, errors.Wrap(err, "invalid command line options")
	}

	return args, nil
}

func (opts Options) usage() []byte {
	buf := bytes.Buffer{}

	fmt.Fprintln(&buf,
		`Usage: pd [options] URL

	Options:
	-h,   -help               print usage and exit
	-p,   -procs <num>        the number of split to download
	-o,   -output <filename>  path and file name of the file downloaded
	-t,   -time <num>         Time limit minuts until the download is completed
	`,
	)

	return buf.Bytes()
}

func ParseOptions(argv []string) (*Options, []string, error) {
	var opts Options
	if len(argv) == 0 {
		os.Stdout.Write(opts.usage())
		return nil, nil, errors.New("no options")
	}

	o, err := opts.parse(argv)
	if err != nil {
		return nil, nil, err
	}

	if opts.Help {
		os.Stdout.Write(opts.usage())
		return nil, nil, errors.New("print usage")
	}

	return &opts, o, nil
}
