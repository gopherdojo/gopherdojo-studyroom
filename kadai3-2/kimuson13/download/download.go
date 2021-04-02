package download

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func Get(url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stdout, "body:%s", b)
	return nil
}
