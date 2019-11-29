package commands

import (
	"github.com/talal/bonclay/pkg/file"
	"github.com/talal/bonclay/pkg/mistertwo"
	"github.com/talal/bonclay/pkg/write"
)

// SyncTask uses 'source:target' pairs defined in the configuration spec
// to link the sources to the targets.
func SyncTask(config *mistertwo.Configuration) {
	write.TaskHeader("sync")

	var errors []string
	for src, target := range config.Spec {
		errs := file.Link(src, target, config.Sync.Overwrite, config.Sync.Clean)
		if len(errs) > 0 {
			for _, v := range errs {
				errors = append(errors, v.Error())
			}
			write.TaskFailure(src, target)
			continue
		}

		write.TaskSuccess(src, target)
	}

	write.TaskFooter("sync", len(errors) == 0)
	write.TaskErrors(errors)
}
