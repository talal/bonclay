package mistertwo

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

// copy copies a src to dst, where src is either a file or a directory. Directories
// are copied recursively. if dst already exists then it is overwritten as per the
// value of overwrite.
// Non-returned errors are sent to a channel.
func copy(overwrite bool, src, dst string) error {
	return filepath.Walk(src, makeWalkFunc(overwrite, src, dst))
}

func makeWalkFunc(overwrite bool, src, dst string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		// validate source
		switch {
		case err == nil:
			// continue down below
		case os.IsNotExist(err):
			return newError(srcNotExists, path)
		default:
			return newError(srcProblem, err)
		}

		// ideally, a source should never be a symlink but in case it is...
		if info.Mode()&os.ModeSymlink != 0 {
			return newError(linkSkip, path)
		}

		// validate destination
		srcIsDir := info.Mode().IsDir()
		dstPath := filepath.Join(dst, strings.TrimPrefix(path, src))
		dstFi, err := os.Lstat(dstPath)
		switch {
		case err == nil:
			if !overwrite {
				return newError(dstExists, dstPath)
			}
			if dstFi.Mode()&os.ModeSymlink != 0 {
				err = os.Remove(dstPath)
				if err != nil {
					return newError(existDeleteFailure, err)
				}
				if srcIsDir {
					err = os.Mkdir(dstPath, info.Mode())
					return newError(dCreateFailure, err)
				}
			}
		case os.IsNotExist(err):
			if srcIsDir {
				err = os.MkdirAll(dstPath, info.Mode())
			} else {
				err = os.MkdirAll(filepath.Dir(dstPath), os.ModePerm)
			}
			if err != nil {
				return newError(dCreateFailure, err)
			}
		default:
			return newError(dstProblem, err)
		}

		if srcIsDir {
			return nil
		}
		return copyFile(overwrite, path, dstPath, info)
	}
}

// copyFile is a helper function for copy() that copies a file from src to dst.
// If dst already exists then it is overwritten as per the value of overwrite.
func copyFile(overwrite bool, src, dst string, srcFi os.FileInfo) error {
	dstFile, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0200)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	err = os.Chmod(dstFile.Name(), srcFi.Mode())
	if err != nil {
		return newError(fCopyFailure, err)
	}

	srcFile, err := os.OpenFile(src, os.O_RDONLY, 0400)
	if err != nil {
		return newError(fCopyFailure, err)
	}
	defer srcFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return newError(fCopyFailure, err)
	}

	return nil
}
