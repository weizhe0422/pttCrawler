package Fetcher

import (
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

var rateLimiter = time.Tick(100 * time.Millisecond)

func Fetch(URL string) ([]byte, bool) {
	var (
		resp    *http.Response
		err     error
		content []byte
	)

	<-rateLimiter
	if resp, err = http.Get(URL); err != nil {
		log.Printf("failed to GET: %s: %v", URL, err)
		return nil, false
	}

	if content, err = httputil.DumpResponse(resp, true); err != nil {
		log.Printf("failed to dump response: %v", err)
		return nil, false
	}

	return content, true
}
