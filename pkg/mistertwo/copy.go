package mistertwo

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// copy copies a src to dst, where src is either a file or a directory. Directories
// are copied recursively. if dst already exists then it is overwritten as per the
// value of overwrite.
// Non-returned errors are sent to a channel.
func copy(overwrite bool, src, dst string, ch chan string) error {
	// validate src and catch common errors
	srcFi, err := os.Lstat(src)
	switch {
	case err == nil:
		// continue down below
	case os.IsNotExist(err):
		return newError(srcNotExists, src)
	default:
		return newError(srcProblem, err)
	}

	// copy src to dst as per its type
	switch mode := srcFi.Mode(); {
	// ideally, a source should never be a symlink but in case it is...
	case mode&os.ModeSymlink != 0:
		return newError(linkSkip, src)
	case mode.IsRegular():
		return copyFile(overwrite, src, dst, srcFi)
	case mode.IsDir():
		// copyDir calls copy() recursively, therefore we send
		// the channel itself to catch any errors
		return copyDir(overwrite, src, dst, srcFi, ch)
	default:
		return newError(unknownType, src)
	}
}

// copyDir is a helper function for copy() that copies a directory recursively from src to dst.
// If dst already exists then its contents are overwritten as per the value of overwrite.
// Non-returned errors are sent to a channel.
func copyDir(overwrite bool, src, dst string, srcFi os.FileInfo, ch chan string) error {
	dstFi, err := os.Lstat(dst)
	switch {
	case err == nil:
		if !overwrite {
			return newError(dstExists, dst)
		}
		if dstFi.Mode()&os.ModeSymlink != 0 {
			err = os.Remove(dst)
			if err != nil {
				return newError(dCopyFailure, err)
			}
			err = os.Mkdir(dst, os.ModePerm)
			if err != nil {
				return newError(dCopyFailure, err)
			}
		}
	case os.IsNotExist(err):
		err = os.MkdirAll(dst, os.ModePerm)
		if err != nil {
			return newError(dCopyFailure, err)
		}
	default:
		return newError(dstProblem, err)
	}

	files, err := ioutil.ReadDir(src)
	if err != nil {
		return newError(dCopyFailure, err)
	}

	for _, file := range files {
		srcPath := filepath.Join(src, file.Name())
		dstPath := filepath.Join(dst, file.Name())
		err = copy(overwrite, srcPath, dstPath, ch)
		if err != nil {
			ch <- err.Error()
		}
	}

	return nil
}

// copyFile is a helper function for copy() that copies a file from src to dst.
// If dst already exists then it is overwritten as per the value of overwrite.
func copyFile(overwrite bool, src, dst string, srcFi os.FileInfo) error {
	dstFi, err := os.Lstat(dst)
	switch {
	case err == nil:
		if !overwrite {
			return newError(dstExists, dst)
		}
		if dstFi.Mode()&os.ModeSymlink != 0 {
			err = os.Remove(dst)
			if err != nil {
				return newError(fCopyFailure, err)
			}
		}
	case os.IsNotExist(err):
		err = os.MkdirAll(filepath.Dir(dst), os.ModePerm)
		if err != nil {
			return newError(fCopyFailure, err)
		}
	default:
		return newError(dstProblem, err)
	}

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
