package task

import (
	"github.com/talal/bonclay/internal/core"
	"github.com/talal/bonclay/internal/file"
)

// Backup uses 'source:target' pairs defined in the configuration spec to copy
// the sources to the targets.
func Backup(config *core.Configuration) {
	core.WriteTaskHeader("backup")

	errors := make([]string, 0, len(config.Spec))
	for src, dst := range config.Spec {
		err := file.Copy(src, dst, config.Backup.Overwrite)
		if err != nil {
			errors = append(errors, err.Error())
			core.WriteTaskFailure(src, dst)
			continue
		}

		core.WriteTaskSuccess(src, dst)
	}

	core.WriteTaskFooter("backup", len(errors) == 0)
	core.WriteTaskErrors(errors)
}
