package tasks

import (
	"github.com/talal/bonclay/pkg/file"
	"github.com/talal/bonclay/pkg/mistertwo"
	"github.com/talal/bonclay/pkg/write"
)

// Backup uses 'source:target' pairs defined in the configuration spec to copy
// the sources to the targets.
func Backup(config *mistertwo.Configuration) {
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
