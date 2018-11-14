package mistertwo

import "fmt"

// BackupTask uses 'source:target' pairs defined in the configuration spec
// to copy the sources to the targets.
func BackupTask(config *Configuration) {
	fmt.Println(withColor(cyan, "bonclay: backup task\n"))

	var errors []string
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

		err = copy(config.Backup.Overwrite, srcPath, dstPath)
		if err != nil {
			errors = append(errors, err.Error())
			printTaskResponse(false, src, dst)
		} else {
			printTaskResponse(true, src, dst)
		}
	}

	if len(errors) > 0 {
		printTaskErrors("backup", errors)
	} else {
		fmt.Println(withColor(green, "\nAll files/directories were successfully backed up."))
	}
}
