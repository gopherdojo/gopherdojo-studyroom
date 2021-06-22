package getheader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

// get all headers
func Headers(w io.Writer, r *http.Response) {
	h := r.Header
	fmt.Fprintln(w, h)
}

// get the specified header
func ResHeader(w io.Writer, r *http.Response, header string) ([]string, error) {
	h, is := r.Header[header]
	fmt.Println(h)
	if !is {
		return nil, fmt.Errorf("cannot find %s header", header)
	}
	fmt.Fprintf(w, "Header[%s] = %s\n", header, h)
	return h, nil
}

// get the specified header by commas
func ResHeaderComma(w io.Writer, r *http.Response, header string) (string, error) {
	h := r.Header.Get(header)
	// if !is {
	// 	return "error", fmt.Errorf("cannot find %s header", header)
	// }
	fmt.Fprintf(w, "Header[%s] = %s\n", header, h)
	return h, nil
}


func GetSize(resp *http.Response) (uint, error) {
	contLen, err := ResHeader(os.Stdout, resp, "Content-Length")
	if err != nil {
		return 0, err
	}
	ret, err := strconv.ParseUint(contLen[0], 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(ret), nil
}

