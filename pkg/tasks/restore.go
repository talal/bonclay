package tasks

import (
	"github.com/talal/bonclay/pkg/file"
	"github.com/talal/bonclay/pkg/mistertwo"
	"github.com/talal/bonclay/pkg/write"
)

// Restore uses 'source:target' pairs defined in the configuration spec
// to copy the targets to the sources. It is the reverse of Backup.
func Restore(config *mistertwo.Configuration) {
	write.TaskHeader("restore")

	errors := make([]string, 0, len(config.Spec))
	for dst, src := range config.Spec {
		err := file.Copy(src, dst, config.Restore.Overwrite)
		if err != nil {
			errors = append(errors, err.Error())
			write.TaskFailure(src, dst)
			continue
		}

		write.TaskSuccess(src, dst)
	}

	write.TaskFooter("restore", len(errors) == 0)
	write.TaskErrors(errors)
}
