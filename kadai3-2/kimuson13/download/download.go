package download

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/asaskevich/govalidator"
)

type Download struct {
	parallel int
	URLs     []string
	args     []string
	timeout  int
}

func New() *Download {
	return &Download{
		parallel: runtime.NumCPU(),
		timeout:  10,
	}
}

func (d *Download) Run() error {
	return nil
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
