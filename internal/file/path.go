package file

import (
	"os"
	"path/filepath"
	"strings"
)

// fullPath returns the absolute path of a file and substitutes any '~' in the
// given path string with the user's home directory location.
func fullPath(path string) (string, error) {
	path = strings.Replace(path, "~", os.Getenv("HOME"), 1)

	return filepath.Abs(path)
}
