package file

import (
	"fmt"
)

const (
	dstParentProblem = "problem with destination's parent directory"
	dstProblem       = "problem with destination"
	srcParentProblem = "problem with source's parent directory"
	srcProblem       = "problem with source"
	targetProblem    = "problem with target"

	dirCleanFail    = "could not clean directory"
	dirCreateFail   = "could not create directory"
	existDeleteFail = "could not delete existing file/directory"
	fileCopyFail    = "could not copy file"
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

type taskError struct {
	Message string
	Value   interface{}
}

func (e *taskError) Error() string {
	return fmt.Sprintf("%s: %v", e.Message, e.Value)
}
