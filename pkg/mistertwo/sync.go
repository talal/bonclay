package mistertwo

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/talal/go-bits/color"
)

// SyncTask uses 'source:target' pairs defined in the configuration spec
// to link the sources to the targets.
func SyncTask(config *Configuration) {
	color.Println(color.Cyan, "bonclay: sync task\n")

	// non-returned errors are received through a
	// channel (if any) and collected in the errors slice
	var errors []string
	ch := make(chan string)
	go func() {
		for i := range ch {
			errors = append(errors, i)
		}
	}()
	for src, target := range config.Spec {
		srcPath, err := fullPath(src)
		if err != nil {
			errors = append(errors, err.Error())
			printTaskResponse(false, src, target)
			continue
		}
		targetPath, err := fullPath(target)
		if err != nil {
			errors = append(errors, err.Error())
			printTaskResponse(false, src, target)
			continue
		}

		err = sync(config.Sync, srcPath, targetPath, ch)
		if err != nil {
			errors = append(errors, err.Error())
			printTaskResponse(false, src, target)
		} else {
			printTaskResponse(true, src, target)
		}
	}
	close(ch)

	if len(errors) > 0 {
		printTaskErrors("sync", errors)
	} else {
		color.Println(color.Green, "\nAll files/directories were successfully synced.")
	}
}

// sync creates a symbolic link between the source and target. If opts.Clean = true then
// broken symlinks in the source's parent directory are removed.
// Non-returned errors are sent to a channel.
func sync(opts SyncOpts, src, target string, ch chan string) error {
	// validate target and catch common errors
	_, err := os.Lstat(target)
	switch {
	case err == nil:
		// continue down below
	case os.IsNotExist(err):
		return newError(targetNotExists, target)
	default:
		return newError(targetProblem, target)
	}

	srcParentDir := filepath.Dir(src)
	// validate source
	srcFi, err := os.Lstat(src)
	switch {
	case err == nil:
		if !opts.Overwrite {
			if srcFi.Mode().IsDir() {
				return newError(dirSkip, src)
			} else if srcFi.Mode().IsRegular() {
				return newError(fileSkip, src)
			}
		}
		err = os.RemoveAll(src)
		if err != nil {
			return newError(syncFailure, err)
		}
	case os.IsNotExist(err):
		// check if source's parent directory exists
		// Create one, if it doesn't
		_, err := os.Stat(srcParentDir)
		if os.IsNotExist(err) {
			err = os.MkdirAll(srcParentDir, os.ModePerm)
			if err != nil {
				return newError(syncFailure, err)
			}
			opts.Clean = false
		} else if err != nil {
			return newError(srcPDirProblem, err)
		}
	default:
		return newError(srcProblem, target)
	}

	err = clean(opts.Clean, srcParentDir, ch)
	if err != nil {
		ch <- err.Error()
	}

	err = os.Symlink(target, src)
	if err != nil {
		return newError(lCreateFailure, err)
	}

	return nil
}

// clean is a helper function for sync(), it removes broken symbolic links
// inside a directory. An error at the directory level will result in function exit.
// For individual files, the errors are sent to a channel.
func clean(shouldClean bool, path string, ch chan string) error {
	if !shouldClean {
		return nil
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return newError(dCleanFailure, err)
	}

	for _, file := range files {
		if file.Mode()&os.ModeSymlink != 0 {
			// check if link is broken
			linkPath := filepath.Join(path, file.Name())
			linkPointee, err := os.Readlink(linkPath)
			if err != nil {
				ch <- newError(lRemoveFailure, err).Error()
			}

			_, err = os.Stat(linkPointee)
			if os.IsNotExist(err) {
				err := os.Remove(linkPath)
				if err != nil {
					ch <- newError(lRemoveFailure, err).Error()
				}
			}
		}
	}

	return nil
}
