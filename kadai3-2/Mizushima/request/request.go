package request

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

// Request returns a response from url and a error.
func Request(method string, url string, setH string, setV string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	if len(setH) == 0 {
		req.Header.Set(setH, setV)
	}
	
	dump, _ := httputil.DumpRequestOut(req, false)
	fmt.Printf("request:\n%s\n", dump)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}


