package Fetcher

import (
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

var rateLimiter = time.Tick(1000 * time.Millisecond)

func Fetch(URL string) ([]byte, bool) {
	var (
		respBody *http.Response
		err      error
		contents []byte
	)
	<-rateLimiter
	if respBody, err = http.Get(URL); err != nil {
		panic(err)
	}
	defer respBody.Body.Close()

	if contents, err = httputil.DumpResponse(respBody, true); err != nil {
		log.Printf("failed to dump response: %v", err)
		return nil, false
	}
	return contents, true
}
