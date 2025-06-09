package v0040

import (
	"context"
	"net"
	"net/http"
)

// UnixSocketRoundTripper handles HTTP transport over Unix sockets
type UnixSocketRoundTripper struct {
	socketPath string
	transport  *http.Transport
}

func NewUnixSocketRoundTripper(socketPath string) *UnixSocketRoundTripper {
	return &UnixSocketRoundTripper{
		socketPath: socketPath,
		transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", socketPath)
			},
		},
	}
}

func (u *UnixSocketRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// Clean up URL formatting for Unix sockets
	req.URL.Scheme = "http"
	req.URL.Host = "unix-socket"
	return u.transport.RoundTrip(req)
}

// NewUnixSocketHTTPClient creates HTTP client for Unix sockets
func NewUnixSocketHTTPClient(socketPath string) *http.Client {
	return &http.Client{
		Transport: NewUnixSocketRoundTripper(socketPath),
	}
}
