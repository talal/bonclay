package commands

import (
	"github.com/talal/bonclay/pkg/file"
	"github.com/talal/bonclay/pkg/mistertwo"
	"github.com/talal/bonclay/pkg/write"
)

// BackupTask uses 'source:target' pairs defined in the configuration spec to copy
// the sources to the targets.
func BackupTask(config *mistertwo.Configuration) {
	write.TaskHeader("backup")

	errors := make([]string, 0, len(config.Spec))
	for src, dst := range config.Spec {
		err := file.Copy(src, dst, config.Backup.Overwrite)
		if err != nil {
			errors = append(errors, err.Error())
			write.TaskFailure(src, dst)
			continue
		}

		write.TaskSuccess(src, dst)
	}

	write.TaskFooter("backup", len(errors) == 0)
	write.TaskErrors(errors)
}
