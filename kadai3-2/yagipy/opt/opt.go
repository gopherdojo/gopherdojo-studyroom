package opt

import (
	"errors"
	"flag"
	"net/url"
	"path"
	"time"
)

var errExist = errors.New("file already exists")

// Options has the options required for parallel-download.
type Options struct {
	Parallelism int
	Output      string
	URL         *url.URL
	Timeout     time.Duration
}

// Parse parses args and returns Options.
func Parse(args ...string) (*Options, error) {
	flg := flag.NewFlagSet("parallel-download", flag.ExitOnError)

	parallelism := flg.Int("p", 8, "Download files in parallel according to the specified number.")
	output := flg.String("o", "", "Save the downloaded file in the specified path. (Overwrite if duplicates.)")
	timeout := flg.Duration("t", 30*time.Second, "Terminate when the specified value has elapsed since download started.")

	flg.Parse(args)

	u, err := url.ParseRequestURI(flg.Arg(0))
	if err != nil {
		return nil, err
	}

	if *output == "" {
		_, filename := path.Split(u.Path)

		// Inspired by the --default-page option of wget
		if filename == "" {
			filename = "index.html"
		}

		*output = filename
	}

	return &Options{
		Parallelism: *parallelism,
		Output:      *output,
		URL:         u,
		Timeout:     *timeout,
	}, nil
}