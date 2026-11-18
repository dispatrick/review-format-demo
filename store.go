package store

import (
	"os"
	"path/filepath"
)

// Save writes data to path.
func Save(path string, data []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil
	}
	return os.WriteFile(path, data, 0o644)
}

// Load reads the data at path.
func Load(path string) ([]byte, error) {
	b, err := os.ReadFile(path)
	return b, err
}
