package tasks

import (
	"github.com/talal/bonclay/pkg/file"
	"github.com/talal/bonclay/pkg/mistertwo"
	"github.com/talal/bonclay/pkg/write"
)

// Sync uses 'source:target' pairs defined in the configuration spec
// to link the sources to the targets.
func Sync(config *mistertwo.Configuration) {
	write.TaskHeader("sync")

	// non-returned errors are received through a
	// channel (if any) and collected in the errors slice
	var errors []string
	ch := make(chan string)
	go func() {
		for i := range ch {
			errors = append(errors, i)
		}
	}()

	for src, target := range config.Spec {
		err := file.Link(src, target, config.Sync.Overwrite, config.Sync.Clean, ch)
		if err != nil {
			errors = append(errors, err.Error())
			write.TaskFailure(src, target)
			continue
		}

		write.TaskSuccess(src, target)
	}

	close(ch)

	write.TaskFooter("sync", len(errors) == 0)
	write.TaskErrors(errors)
}
