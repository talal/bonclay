package write

import (
	"fmt"
	"strings"

	"github.com/talal/go-bits/color"
)

const arrow = "-->"

// TaskHeader writes the task header message followed by a new line.
func TaskHeader(taskName string) {
	color.Printf(color.Cyan, "bonclay: %s task\n\n", taskName)
}

// TaskFooter writes the task footer message followed by a new line.
func TaskFooter(taskName string, wasSuccessful bool) {
	if wasSuccessful {
		color.Printf(color.Green, "\n===> %s Successful\n\n", strings.Title(taskName))
	} else {
		color.Printf(color.Red, "\nSome errors occurred during %s:\n", taskName)
	}
}

// TaskSuccess writes a success response for a src/dst pair, where src/dst is
// either a file or a directory.
func TaskSuccess(src, dst string) {
	taskResponse(src, dst, true)
}

// TaskFailure writes a failure response for a src/dst pair, where src/dst is
// either a file or a directory.
func TaskFailure(src, dst string) {
	taskResponse(src, dst, false)
}

func taskResponse(src, dst string, wasSuccessful bool) {
	c := color.Green
	if !wasSuccessful {
		c = color.Red
	}

	fmt.Printf("%s %s %s\n", color.Sprintf(color.Blue, src),
		color.Sprintf(c, arrow), color.Sprintf(color.Blue, dst))
}

// TaskErrors writes the errors, if any occurred.
//
// Duplicates are removed.
func TaskErrors(errors []string) {
	if len(errors) == 0 {
		return
	}

	var uniqueErrors []string
	var errExists = make(map[string]bool)
	for _, v := range errors {
		if _, exists := errExists[v]; !exists {
			uniqueErrors = append(uniqueErrors, v)
			errExists[v] = true
		}
	}

	for i, v := range uniqueErrors {
		if i == (len(uniqueErrors) - 1) {
			fmt.Printf("\t%s\n\n", v)
		} else {
			fmt.Printf("\t%s\n", v)
		}
	}
}
