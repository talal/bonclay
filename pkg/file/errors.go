package file

import (
	"fmt"
)

const (
	dstProblem       = "problem with destination"
	srcProblem       = "problem with source"
	srcParentProblem = "problem with source's parent directory"
	targetProblem    = "problem with target"

	existDeleteFail = "could not delete existing file/directory"
	fileCopyFail    = "could not copy file"
	dirCreateFail   = "could not create directory"
	dirCleanFail    = "could not clean directory"
	linkCreateFail  = "could not create symlink"
	linkRemoveFail  = "could not remove broken symlink"
	syncFailure     = "could not sync"

	dstExists       = "destination already exists"
	srcNotExists    = "source does not exist"
	targetNotExists = "target does not exist"

	dirSkip  = "directory skipped"
	fileSkip = "file skipped"
	linkSkip = "symlink skipped"
)

// makeError is a helper function that for a given string and some value,
// returns a new error in the form of "string: value".
func makeError(str string, val interface{}) error {
	return fmt.Errorf("%s: %v", str, val)
}
