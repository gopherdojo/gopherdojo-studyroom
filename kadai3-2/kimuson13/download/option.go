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
func (opts *Options) Parse(args []string) (*Options, error) {
	parallel := flag.Int("p", runtime.NumCPU(), "download files with parallel")
	timeout := flag.Int("t", 30, "timeout for this second")
	filename := flag.String("f", "paralleldownload", "save the fiel as this name")
	flag.Parse()
	u, err := url.Parse(flag.Arg(0))
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
