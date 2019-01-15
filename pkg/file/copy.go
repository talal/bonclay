package file

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Copy copies a src to dst, where src is either a file or a directory.
// Directories are copied recursively. if dst already exists then it is
// overwritten as per the value of overwrite.
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
				return makeError(srcNotExists, path)
			}

			return makeError(srcProblem, err)
		}

		// source should not be a symlink
		if info.Mode()&os.ModeSymlink != 0 {
			return makeError(linkSkip, path)
		}

		// validate destination
		srcIsDir := info.Mode().IsDir()
		dstPath := filepath.Join(dst, strings.TrimPrefix(path, src))
		dstFi, err := os.Lstat(dstPath)
		if err != nil {
			if os.IsNotExist(err) {
				if srcIsDir {
					err = os.MkdirAll(dstPath, info.Mode())
				} else {
					err = os.MkdirAll(filepath.Dir(dstPath), os.ModePerm)
				}

				if err != nil {
					return makeError(dirCreateFail, err)
				}
			}

			return makeError(dstProblem, err)
		}

		if !overwrite {
			return makeError(dstExists, dstPath)
		}

		// if destination is a symlink then delete it, and create an empty
		// directory if needed
		if dstFi.Mode()&os.ModeSymlink != 0 {
			err = os.Remove(dstPath)
			if err != nil {
				return makeError(existDeleteFail, err)
			}

			if srcIsDir {
				err = os.Mkdir(dstPath, info.Mode())
				if err != nil {
					return makeError(dirCreateFail, err)
				}
			}
		}

		if srcIsDir {
			return nil
		}

		return copyFile(path, dstPath, info, overwrite)
	}
}

// copyFile is a helper function for copy() that copies a file from src to dst.
// If dst already exists then it is overwritten as per the value of overwrite.
func copyFile(src, dst string, srcFi os.FileInfo, overwrite bool) error {
	dstFile, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0200)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	err = os.Chmod(dstFile.Name(), srcFi.Mode())
	if err != nil {
		return makeError(fileCopyFail, err)
	}

	srcFile, err := os.OpenFile(src, os.O_RDONLY, 0400)
	if err != nil {
		return makeError(fileCopyFail, err)
	}
	defer srcFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return makeError(fileCopyFail, err)
	}

	return nil
}
