package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// Link creates a symbolic link between the source and target.
//
// If cleanParentDir = true then broken symlinks in the source's parent
// directory are removed.
func Link(src, target string, overwrite bool, cleanParentDir bool) (errors []error) {
	srcAbsPath, err := fullPath(src)
	if err != nil {
		return append(errors, err)
	}
	targetAbsPath, err := fullPath(target)
	if err != nil {
		return append(errors, err)
	}

	// validate target
	_, err = os.Lstat(targetAbsPath)
	if err != nil {
		if os.IsNotExist(err) {
			return append(errors, &taskError{targetNotExists, targetAbsPath})
		}

		return append(errors, &taskError{targetProblem, err})
	}

	srcParentDir := filepath.Dir(srcAbsPath)

	// validate source
	srcFi, err := os.Lstat(srcAbsPath)
	switch {
	case err == nil:
		// no need to continue further if source exists and overwrite is false
		// existing symlinks are deleted regardless of the overwrite value
		if !overwrite {
			if srcFi.Mode().IsDir() {
				return append(errors, &taskError{dirSkip, srcAbsPath})
			}
			if srcFi.Mode().IsRegular() {
				return append(errors, &taskError{fileSkip, srcAbsPath})
			}
		}

		// delete any existing source file/directory
		err = os.RemoveAll(srcAbsPath)
		if err != nil {
			return append(errors, &taskError{syncFailure, err})
		}
	case os.IsNotExist(err):
		// check if source's parent directory exists create one, if it doesn't
		_, err = os.Lstat(srcParentDir)
		if err != nil {
			if os.IsNotExist(err) {
				err = os.MkdirAll(srcParentDir, os.ModePerm)
				if err != nil {
					return append(errors, &taskError{dirCreateFail, err})
				}
				// no need to clean a newly created directory
				cleanParentDir = false
			} else {
				return append(errors, &taskError{srcParentProblem, err})
			}
		}
	default:
		return append(errors, &taskError{srcProblem, err})
	}

	if cleanParentDir {
		removeBrokenSymlinks(srcParentDir, &errors)
	}

	err = os.Symlink(targetAbsPath, srcAbsPath)
	if err != nil {
		return append(errors, &taskError{linkCreateFail, err})
	}

	return nil
}

// removeBrokenSymlinks is a helper function for sync(), it removes broken
// symbolic links inside a directory.
//
// An error at the directory level will result in function exit.
func removeBrokenSymlinks(path string, errors *[]error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		*errors = append(*errors, &taskError{dirCleanFail, err})
		return
	}

	for _, file := range files {
		if file.Mode()&os.ModeSymlink != 0 {
			// check if link is broken
			linkPath := filepath.Join(path, file.Name())
			linkTarget, err := os.Readlink(linkPath)
			if err != nil {
				*errors = append(*errors, &taskError{linkRemoveFail, err})
				continue
			}

			_, err = os.Stat(linkTarget)
			if os.IsNotExist(err) {
				err = os.Remove(linkPath)
				if err != nil {
					*errors = append(*errors, &taskError{linkRemoveFail, err})
				}
			}
		}
	}
}
