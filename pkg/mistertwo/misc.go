package mistertwo

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type color string

const (
	red     color = "0;31"
	green   color = "0;32"
	blue    color = "0;34"
	magenta color = "0;35"
	cyan    color = "0;36"
)

func fullPath(str string) (string, error) {
	str = strings.Replace(str, "~", os.Getenv("HOME"), 1)

	path, err := filepath.Abs(str)
	if err != nil {
		return "", err
	}

	return path, nil
}

func withColorf(c color, str string, args ...interface{}) string {
	return withColor(c, fmt.Sprintf(str, args...))
}

func withColor(c color, str string) string {
	return fmt.Sprintf("\x1B[%sm%s\x1B[0m", c, str)
}

func printTaskResponse(success bool, src, dst string) {
	arrow := "-->"
	c := green
	if !success {
		c = red
	}

	fmt.Printf("%s %s %s\n", withColor(blue, src), withColor(c, arrow), withColor(blue, dst))
}
