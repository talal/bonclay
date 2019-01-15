package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// Link creates a symbolic link between the source and target.
// If cleanParentDir = true then broken symlinks in the source's parent
// directory are removed.
//
// Non-returned errors are sent to a channel.
func Link(src, target string, overwrite bool, cleanParentDir bool, ch chan string) error {
	srcAbsPath, err := fullPath(src)
	if err != nil {
		return err
	}
	targetAbsPath, err := fullPath(target)
	if err != nil {
		return err
	}

	// validate target
	_, err = os.Lstat(targetAbsPath)
	if err != nil {
		if os.IsNotExist(err) {
			return makeError(targetNotExists, targetAbsPath)
		}

		return makeError(targetProblem, err)
	}

	srcParentDir := filepath.Dir(srcAbsPath)
	// validate source
	srcFi, err := os.Lstat(srcAbsPath)
	if err != nil {
		if os.IsNotExist(err) {
			// check if source's parent directory exists
			// create one, if it doesn't
			_, err := os.Stat(srcParentDir)
			if err != nil {
				if os.IsNotExist(err) {
					err = os.MkdirAll(srcParentDir, os.ModePerm)
					if err != nil {
						return makeError(syncFailure, err)
					}
					// no need to clean a newly created directory
					cleanParentDir = false
				}

				return makeError(srcParentProblem, err)
			}
		}

		return makeError(srcProblem, targetAbsPath)
	}

	if !overwrite {
		if srcFi.Mode().IsDir() {
			return makeError(dirSkip, srcAbsPath)
		}
		if srcFi.Mode().IsRegular() {
			return makeError(fileSkip, srcAbsPath)
		}
	}

	err = os.RemoveAll(srcAbsPath)
	if err != nil {
		return makeError(syncFailure, err)
	}

	if cleanParentDir {
		removeBrokenSymlinks(srcParentDir, ch)
	}

	err = os.Symlink(targetAbsPath, srcAbsPath)
	if err != nil {
		return makeError(linkCreateFail, err)
	}

	return nil
}

// removeBrokenSymlinks is a helper function for sync(), it removes broken
// symbolic links inside a directory.
//
// An error at the directory level will result in function exit.
func removeBrokenSymlinks(path string, ch chan string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		ch <- makeError(dirCleanFail, err).Error()
		return
	}

	for _, file := range files {
		if file.Mode()&os.ModeSymlink != 0 {
			// check if link is broken
			linkPath := filepath.Join(path, file.Name())
			linkDestination, err := os.Readlink(linkPath)
			if err != nil {
				ch <- makeError(linkRemoveFail, err).Error()
				continue
			}

			_, err = os.Stat(linkDestination)
			if os.IsNotExist(err) {
				err := os.Remove(linkPath)
				if err != nil {
					ch <- makeError(linkRemoveFail, err).Error()
				}
			}
		}
	}
}
