// Package appname returns the current executable's name. If the name ends with a .exe suffix, it is elided.
package appname

import (
	"os"
	"path/filepath"
	"strings"
)

func Get() string {
	return get(os.Args[0])
}

func get(s string) string {
	s = filepath.Base(s)
	ext := filepath.Ext(s)
	if ext == ".exe" {
		s = strings.TrimSuffix(s, ext)
	}
	return s
}
