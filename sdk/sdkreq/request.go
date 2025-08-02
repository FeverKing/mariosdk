package sdkreq

import (
	"io"
	"net/http"
	"time"
)

type Requester interface {
	Do(r *http.Request) (*http.Response, error)
}

type HttpRequester struct {
	hClient *http.Client
}

func NewHttpRequester() *HttpRequester {
	return &HttpRequester{
		hClient: &http.Client{
			Timeout: 10 * time.Minute,
		},
	}
}

func (r *HttpRequester) Do(req *http.Request) (*http.Response, error) {
	return r.hClient.Do(req)
}

func NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequest(method, url, body)
}
