package download

import (
	"errors"
	"net/http"
	"strconv"
)

func GetContentLength(url string) (size int, err error) {
	res, err := http.Head(url)
	if err != nil {
		return 0, err
	}
	cl, ok := res.Header["Content-Length"]
	if !ok {
		return 0, errors.New("Content-Length does not be found\n")
	}
	size, err = strconv.Atoi(cl[0])
	return
}
