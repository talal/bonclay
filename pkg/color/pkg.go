package color

import "fmt"

type color string

// ANSI color escape sequences
// https://en.wikipedia.org/wiki/ANSI_escape_code#Colors
const (
	Red   color = "0;31"
	Green color = "0;32"
	Blue  color = "0;34"
	Cyan  color = "0;36"
)

// Sprintf is like fmt.Sprintf but with color.
func Sprintf(c color, str string, args ...interface{}) string {
	return fmt.Sprintf("\x1B[%sm%s\x1B[0m", c, fmt.Sprintf(str, args...))
}

// Printf is like fmt.Printf but with color.
func Printf(c color, str string, args ...interface{}) {
	fmt.Printf("\x1B[%sm%s\x1B[0m", c, fmt.Sprintf(str, args...))
}

// Println is like fmt.Println but with color.
func Println(c color, str string) {
	fmt.Println("\x1B[" + string(c) + "m" + str + "\x1B[0m")
}
