package download

import (
	"flag"
	"net/url"
	"runtime"
)

// Options struct
type Options struct {
	Parallel int
	Timeout  int
	Filename string
	URL      string
}

// Parse method parses options
func (opts *Options) Parse(args ...string) (*Options, error) {
	flg := flag.NewFlagSet("parallelDownload", flag.ExitOnError)
	parallel := flg.Int("p", runtime.NumCPU(), "separate Content-Length with this argument")
	timeout := flg.Int("t", 30, "timeout for this second")
	filename := flg.String("f", "paralleldownload", "save the file as this name")
	if err := flg.Parse(args); err != nil {
		return nil, err
	}
	u, err := url.Parse(flg.Arg(0))
	if err != nil {
		return nil, err
	}

	return &Options{
		Parallel: *parallel,
		Timeout:  *timeout,
		Filename: *filename,
		URL:      u.String(),
	}, nil
}
