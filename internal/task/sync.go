package task

import (
	"github.com/talal/bonclay/internal/core"
	"github.com/talal/bonclay/internal/file"
)

// Sync uses 'source:target' pairs defined in the configuration spec to link
// the sources to the targets.
func Sync(config *core.Configuration) {
	core.WriteTaskHeader("sync")

	var errors []string
	for src, target := range config.Spec {
		errs := file.Link(src, target, config.Sync.Overwrite, config.Sync.Clean)
		if len(errs) > 0 {
			for _, v := range errs {
				errors = append(errors, v.Error())
			}
			core.WriteTaskFailure(src, target)
			continue
		}

		core.WriteTaskSuccess(src, target)
	}

	core.WriteTaskFooter("sync", len(errors) == 0)
	core.WriteTaskErrors(errors)
}
