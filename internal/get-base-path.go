package internal

import (
	"path/filepath"
	"runtime"
)

// GetBasePath returns the path to the project root
func GetBasePath() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "../")
}
