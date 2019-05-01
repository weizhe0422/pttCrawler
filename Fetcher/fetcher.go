package Fetcher

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func Fetch(Url string) ([]byte, error) {
	var (
		resp *http.Response
		err  error
	)
	if resp, err = http.Get(Url); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	contents, _ := httputil.DumpResponse(resp, true)
	log.Print(string(contents))
	return httputil.DumpResponse(resp, true)
}
