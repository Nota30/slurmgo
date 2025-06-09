package v0040

import (
	"net/http"
	"os"
)

type UnixSocketConfig struct {
	*Configuration
	socketPath string
}

func NewUnixSocketConfig(socketPath string) *UnixSocketConfig {
	if _, err := os.Stat(socketPath); err != nil {
		panic("Slurm socket not found: " + socketPath)
	}

	return &UnixSocketConfig{
		Configuration: NewConfiguration(),
		socketPath:    socketPath,
	}
}

func (c *UnixSocketConfig) HTTPClient() *http.Client {
	return NewUnixSocketHTTPClient(c.socketPath).client
}

func (c *UnixSocketConfig) Host() string {
	return "http:%2f%2funix" // Required for proper URL resolution
}
