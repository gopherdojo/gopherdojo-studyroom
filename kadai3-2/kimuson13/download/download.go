package download

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/kimuson13/gopherdojo-studyroom/kimuson13/options"
)

type Download struct {
	parallel int
	URLs     []string
	args     []string
	timeout  time.Duration
	filename string
}

func New() *Download {
	return &Download{
		parallel: runtime.NumCPU(),
		timeout:  10,
	}
}

func (d *Download) Run() error {
	if err := d.Ready(); err != nil {
		return err
	}
	if err := d.Download(); err != nil {
		return err
	}
	if err := d.MergeFiles(); err != nil {
		return err
	}
	return nil
}

func (d *Download) Ready() error {
	opts, err := d.parseOptions(os.Args[1:])
	if err != nil {
		return errors.New("failed to parse command line args")
	}
	if opts.Parallel > 2 {
		d.parallel = opts.Parallel
	}
	if opts.Timeout > 0 {
		d.timeout = opts.Timeout
	}
	if opts.Output != "" {
		d.filename = opts.Output
	}
	if err := d.parseURLs(); err != nil {
		return errors.New("failed to parse of url")
	}
	return nil
}

func (d *Download) Download() error {
	return nil
}

func (d *Download) MergeFiles() error {
	return nil
}

func (d *Download) parseOptions(argv []string) (*options.Options, error) {
	var opt options.Options
	o, err := opt.Parse(argv)
	if err != nil {
		return nil, errors.New("failed to parse command line options")
	}

	if opt.Help {
		os.Stdout.Write(opt.Usage())
		return nil, errors.New("this is usage")
	}
	d.args = o
	return &opt, nil
}

func (d *Download) parseURLs() error {
	for _, arg := range d.args {
		if govalidator.IsURL(arg) {
			d.URLs = append(d.URLs, arg)
		}
	}
	if len(d.URLs) > 1 {
		fmt.Fprintf(os.Stdout, "please input url\n")
		fmt.Fprintf(os.Stdout, "start download...")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			scan := scanner.Text()
			urls := strings.Split(scan, " ")
			for _, url := range urls {
				if govalidator.IsURL(url) {
					d.URLs = append(d.URLs, url)
				}
			}
		}
		if err := scanner.Err(); err != nil {
			return err
		}
		if len(d.URLs) < 1 {
			return errors.New("urls not found in the argument")
		}
	}
	return nil
}
