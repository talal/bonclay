/**********************************************************************************
*
* Copyright 2018 Muhammad Talal Anwar <talalanwar@protonmail.com>
*
* This program is free software: you can redistribute it and/or modify it under
* the terms of the GNU General Public License as published by the Free Software
* Foundation, either version 3 of the License, or (at your option) any later
* version.
*
* This program is distributed in the hope that it will be useful, but WITHOUT ANY
* WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR
* A PARTICULAR PURPOSE. See the GNU General Public License for more details.
*
* You should have received a copy of the GNU General Public License along with
* this program. If not, see <http://www.gnu.org/licenses/>.
*
**********************************************************************************/

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
