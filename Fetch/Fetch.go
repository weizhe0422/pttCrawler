package Fetch

import (
	"net/http"
	"net/http/httputil"
	"time"
)

var rateLimiter = time.Tick(1000 * time.Millisecond)

func Fetch(URL string) ([]byte, error) {
	var (
		respBody *http.Response
		err      error
	)
	<- rateLimiter
	if respBody, err = http.Get(URL); err != nil {
		return nil, err
	}

	defer respBody.Body.Close()

	return httputil.DumpResponse(respBody, true)
}
