package color

import (
	"fmt"
	"io"
)

// ANSICode type defines some color in the form of ANSI color escape sequence.
// https://en.wikipedia.org/wiki/ANSI_escape_code#Colors
type ANSICode string

const (
	Black   ANSICode = "0;30"
	Red     ANSICode = "0;31"
	Green   ANSICode = "0;32"
	Yellow  ANSICode = "0;33"
	Blue    ANSICode = "0;34"
	Magenta ANSICode = "0;35"
	Cyan    ANSICode = "0;36"
	White   ANSICode = "0;37"
)

const (
	BrightBlack   ANSICode = "0;90"
	BrightRed     ANSICode = "0;91"
	BrightGreen   ANSICode = "0;92"
	BrightYellow  ANSICode = "0;93"
	BrightBlue    ANSICode = "0;94"
	BrightMagenta ANSICode = "0;95"
	BrightCyan    ANSICode = "0;96"
	BrightWhite   ANSICode = "0;97"
)

// Fprint is like fmt.Fprint but with color.
func Fprint(w io.Writer, c ANSICode, a ...interface{}) (n int, err error) {
	format := makeFormat(a)
	return fmt.Fprintf(w, "\x1B[%sm%s\x1B[0m", c, fmt.Sprintf(format, a...))
}

// Fprintf is like fmt.Fprintf but with color.
func Fprintf(w io.Writer, c ANSICode, format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(w, "\x1B[%sm%s\x1B[0m", c, fmt.Sprintf(format, a...))
}

// Fprintln is like fmt.Fprintln but with color.
func Fprintln(w io.Writer, c ANSICode, a ...interface{}) (n int, err error) {
	format := makeFormat(a)
	return fmt.Fprintf(w, "\x1B[%sm%s\x1B[0m\n", c, fmt.Sprintf(format, a...))
}

// Sprintf is like fmt.Sprintf but with color.
func Sprintf(c ANSICode, format string, a ...interface{}) string {
	return fmt.Sprintf("\x1B[%sm%s\x1B[0m", c, fmt.Sprintf(format, a...))
}

// Printf is like fmt.Printf but with color.
func Printf(c ANSICode, format string, a ...interface{}) (n int, err error) {
	return fmt.Printf("\x1B[%sm%s\x1B[0m", c, fmt.Sprintf(format, a...))
}

// Println is like fmt.Println but with color.
func Println(c ANSICode, a ...interface{}) (n int, err error) {
	format := makeFormat(a)
	return fmt.Printf("\x1B[%sm%s\x1B[0m\n", c, fmt.Sprintf(format, a...))
}

func makeFormat(a ...interface{}) (format string) {
	for i := 0; i < len(a); i++ {
		if i == 0 {
			format += "%v"
		} else {
			format += " %v"
		}
	}
	return format
}
