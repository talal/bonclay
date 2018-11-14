package mistertwo

import (
	"fmt"

	"github.com/alecthomas/kingpin"
	"github.com/talal/bonclay/pkg/color"
)

const (
	dstProblem     = "problem with destination"
	srcProblem     = "problem with source"
	srcPDirProblem = "problem with source's parent directory"
	targetProblem  = "problem with target"

	existDeleteFailure = "could not delete existing file/directory"
	fCopyFailure       = "could not copy file"
	dCreateFailure     = "could not create directory"
	dCleanFailure      = "could not clean directory"
	lCreateFailure     = "could not create symlink"
	lRemoveFailure     = "could not remove broken symlink"
	syncFailure        = "could not sync"

	dstExists       = "destination already exists"
	srcNotExists    = "source does not exist"
	targetNotExists = "target does not exist"

	dirSkip  = "directory skipped"
	fileSkip = "file skipped"
	linkSkip = "symlink skipped"
)

// newError is a helper function that for a given string and value returns a
// new error in the form of "string: value"
func newError(str string, val interface{}) error {
	return fmt.Errorf("%s: %v", str, val)
}

// printTaskErrors removes duplicates in the errors slice and prints
// the errors along with the task name.
func printTaskErrors(taskName string, errors []string) {
	var uniqueErrors []string
	var errExists = make(map[string]bool)
	for _, v := range errors {
		if _, exists := errExists[v]; !exists {
			uniqueErrors = append(uniqueErrors, v)
			errExists[v] = true
		}
	}

	color.Printf(color.Red, "\nSome errors occured during %s:\n", taskName)
	for _, v := range uniqueErrors {
		fmt.Printf("\t%s\n", v)
	}
}

// fatalIfError is a convenient wrapper for kingpin.FatalIfError().
func fatalIfError(err error, str string) {
	kingpin.FatalIfError(err, str)
}
