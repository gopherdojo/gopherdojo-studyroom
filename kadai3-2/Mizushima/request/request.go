package request

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func Request(method string, url string, setH string, setV string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set(setH, setV)

	dump, _ := httputil.DumpRequestOut(req, false)
	fmt.Printf("request:\n%s\n", dump)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}


