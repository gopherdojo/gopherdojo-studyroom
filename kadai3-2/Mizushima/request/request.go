package request

import (
	"context"
	"fmt"
	"net/http"
)

// Request throws a request and returns a response from url and a error.
func Request(ctx context.Context, method string, url string, setH string, setV string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return nil, err
	}

	if len(setH) != 0 {
		req.Header.Set(setH, setV)
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request.Request err: %s", err)
	}
	return resp, nil
}
