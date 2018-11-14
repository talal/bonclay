package mistertwo

import "fmt"

// BackupTask uses 'source:target' pairs defined in the configuration spec
// to copy the sources to the targets.
func BackupTask(config *Configuration) {
	fmt.Println(withColor(cyan, "bonclay: backup task\n"))

	// since copy() is called recursively, therefore non-returned errors are
	// received through a channel (if any) and collected in the errors slice
	var errors []string
	ch := make(chan string)
	go func() {
		for i := range ch {
			errors = append(errors, i)
		}
	}()
	for src, dst := range config.Spec {
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

		err = copy(config.Backup.Overwrite, srcPath, dstPath, ch)
		if err != nil {
			errors = append(errors, err.Error())
			printTaskResponse(false, src, dst)
		} else {
			printTaskResponse(true, src, dst)
		}
	}
	close(ch)

	if len(errors) > 0 {
		printTaskErrors("backup", errors)
	} else {
		fmt.Println(withColor(green, "\nAll files/directories were successfully backed up."))
	}
}
