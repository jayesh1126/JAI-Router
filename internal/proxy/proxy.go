package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func NewProxy(baseURL, apiKey string) (http.Handler, error) {
	target, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			rest := strings.TrimPrefix(req.URL.Path, "/v1")

			req.URL.Scheme = target.Scheme
			req.URL.Host = target.Host
			req.URL.Path = target.Path + rest
			req.Host = target.Host
			req.Header.Set("Authorization", "Bearer "+apiKey)
			// Remove hop-by-hop headers that shouldn't be forwarded
			req.Header.Del("Accept-Encoding")
			req.Header.Set("stream", "true")
		},
		FlushInterval: -1,
	}

	return proxy, nil
}
