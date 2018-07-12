package fs

import (
	"os"
)

func EnsureDir(path string) error {
	return os.MkdirAll(path, 0755)
}

func FileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}
