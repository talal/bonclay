package mistertwo

import (
	"fmt"
)

// RestoreTask uses 'source:target' pairs defined in the configuration spec
// to copy the targets to the sources. It is the reverse of BackupTask.
func RestoreTask(config *Configuration) {
	fmt.Println(withColor(cyan, "bonclay: restore task\n"))

	// since copy() is called recursively, therefore non-returned errors are
	// received through a channel (if any) and collected in the errors slice
	var errors []string
	ch := make(chan string)
	go func() {
		for i := range ch {
			errors = append(errors, i)
		}
	}()
	for dst, src := range config.Spec {
		srcPath, err := fullPath(src)
		if err != nil {
			errors = append(errors, err.Error())
			printTaskResponse(false, src, dst)
			continue
		}
		dstPath, err := fullPath(dst)
		if err != nil {
			errors = append(errors, err.Error())
			printTaskResponse(false, src, dst)
			continue
		}

		err = copy(config.Restore.Overwrite, srcPath, dstPath, ch)
		if err != nil {
			errors = append(errors, err.Error())
			printTaskResponse(false, src, dst)
		} else {
			printTaskResponse(true, src, dst)
		}
	}
	close(ch)

	if len(errors) > 0 {
		printTaskErrors("restore", errors)
	} else {
		fmt.Println(withColor(green, "\nAll files/directories were successfully restored."))
	}
}
