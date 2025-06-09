package v0040

import (
	"os"
	"path/filepath"
)

// NewUnixSocketConfiguration creates API config for Unix sockets
func NewUnixSocketConfiguration(socketPath string) *Configuration {
	// Resolve absolute path to socket
	absPath, err := filepath.Abs(socketPath)
	if err != nil {
		panic("invalid socket path: " + socketPath)
	}
	
	// Verify socket exists
	if _, err := os.Stat(absPath); err != nil {
		panic("Slurm socket not found: " + absPath)
	}

	cfg := NewConfiguration()
	cfg.Host = "unix-socket" // Dummy hostname
	cfg.Scheme = "http"
	cfg.HTTPClient = NewUnixSocketHTTPClient(absPath)
	
	return cfg
}
