package v0040

import (
	"context"
	"net"
	"net/http"
)

type UnixSocketHTTPClient struct {
	socketPath string
	client     *http.Client
}

func NewUnixSocketHTTPClient(socketPath string) *UnixSocketHTTPClient {
	return &UnixSocketHTTPClient{
		socketPath: socketPath,
		client: &http.Client{
			Transport: &http.Transport{
				DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
					return net.Dial("unix", socketPath)
				},
			},
		},
	}
}

func (c *UnixSocketHTTPClient) Do(req *http.Request) (*http.Response, error) {
	// Rewrite URL to use http+unix scheme
	req.URL.Scheme = "http"
	req.URL.Host = "unix"
	return c.client.Do(req)
}
