package utils

import (
	"path"
	"path/filepath"
	"runtime"
	"time"
)

// Fetch root directory path
func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func TimeToSQLDateConverter(input time.Time) string {
	return input.Format("2006-01-02")
}
