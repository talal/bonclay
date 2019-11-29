package task

import (
	"github.com/talal/bonclay/internal/core"
	"github.com/talal/bonclay/internal/file"
)

// Restore uses 'source:target' pairs defined in the configuration spec to copy
// the targets to the sources. It is the reverse of Backup.
func Restore(config *core.Configuration) {
	core.WriteTaskHeader("restore")

	errors := make([]string, 0, len(config.Spec))
	for dst, src := range config.Spec {
		err := file.Copy(src, dst, config.Restore.Overwrite)
		if err != nil {
			errors = append(errors, err.Error())
			core.WriteTaskFailure(src, dst)
			continue
		}

		core.WriteTaskSuccess(src, dst)
	}

	core.WriteTaskFooter("restore", len(errors) == 0)
	core.WriteTaskErrors(errors)
}
