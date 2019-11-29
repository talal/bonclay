package file

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Copy copies a src to dst, where src is either a file or a directory.
//
// Directories are copied recursively. if dst already exists then it is
// overwritten as per the value of overcore.
func Copy(src, dst string, overwrite bool) error {
	srcAbsPath, err := fullPath(src)
	if err != nil {
		return err
	}
	dstAbsPath, err := fullPath(dst)
	if err != nil {
		return err
	}

	return filepath.Walk(srcAbsPath, makeWalkFunc(srcAbsPath, dstAbsPath, overwrite))
}

func makeWalkFunc(src, dst string, overwrite bool) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		// validate source
		if err != nil {
			if os.IsNotExist(err) {
				return &taskError{srcNotExists, path}
			}

			return &taskError{srcProblem, err}
		}

		// source should not be a symlink
		if info.Mode()&os.ModeSymlink != 0 {
			return &taskError{linkSkip, path}
		}

		srcIsDir := info.Mode().IsDir()

		// validate destination
		dstPath := filepath.Join(dst, strings.TrimPrefix(path, src))
		dstFi, err := os.Lstat(dstPath)
		switch {
		case err == nil:
			// no need to continue further if destination exists and overwrite
			// is false
			if !overwrite {
				return &taskError{dstExists, dstPath}
			}

			// if destination is a symlink then delete it and create
			// a new directory if required
			if dstFi.Mode()&os.ModeSymlink != 0 {
				err = os.Remove(dstPath)
				if err != nil {
					return &taskError{existDeleteFail, err}
				}

				if srcIsDir {
					err = os.Mkdir(dstPath, info.Mode())
					if err != nil {
						return &taskError{dirCreateFail, err}
					}
				}
			}
		case os.IsNotExist(err):
			if srcIsDir {
				err = os.MkdirAll(dstPath, info.Mode())
				if err != nil {
					return &taskError{dirCreateFail, err}
				}
			} else {
				// check if destination's parent directory exists, create one
				// if required
				_, err = os.Lstat(filepath.Dir(dstPath))
				if err != nil {
					if os.IsNotExist(err) {
						err = os.MkdirAll(filepath.Dir(dstPath), os.ModePerm)
						if err != nil {
							return &taskError{dirCreateFail, err}
						}
					} else {
						return &taskError{dstParentProblem, err}
					}
				}
			}
		default:
			return &taskError{dstProblem, err}
		}

		if srcIsDir {
			return nil
		}

		return copyFile(path, dstPath, info, overwrite)
	}
}

// copyFile is a helper function for Copy() that copies a file from src to dst.
// If dst already exists then it is overwritten as per the value of overcore.
func copyFile(src, dst string, srcFi os.FileInfo, overwrite bool) error {
	dstFile, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0200)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	err = os.Chmod(dstFile.Name(), srcFi.Mode())
	if err != nil {
		return &taskError{fileCopyFail, err}
	}

	srcFile, err := os.OpenFile(src, os.O_RDONLY, 0400)
	if err != nil {
		return &taskError{fileCopyFail, err}
	}
	defer srcFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return &taskError{fileCopyFail, err}
	}
	dstFile.Sync()

	return nil
}
