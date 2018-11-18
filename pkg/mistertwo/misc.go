package mistertwo

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/talal/go-bits/color"
)

func fullPath(str string) (string, error) {
	str = strings.Replace(str, "~", os.Getenv("HOME"), 1)

	path, err := filepath.Abs(str)
	if err != nil {
		return "", err
	}

	return path, nil
}

func printTaskResponse(success bool, src, dst string) {
	arrow := "-->"
	c := color.Green
	if !success {
		c = color.Red
	}

	fmt.Printf("%s %s %s\n", color.Sprintf(color.Blue, src),
		color.Sprintf(c, arrow), color.Sprintf(color.Blue, dst))
}
