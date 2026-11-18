package store

import "os"

// Save writes data to path.
func Save(path string, data []byte) error {
	return os.WriteFile(path, data, 0o644)
}
