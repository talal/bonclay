package commands

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/talal/bonclay/pkg/file"
	"github.com/talal/bonclay/pkg/mistertwo"
	"github.com/talal/bonclay/pkg/write"
	"github.com/talal/go-bits/cli"
)

const syncDesc = `Sync files/directories.`

const syncUsage = `
Usage: bonclay sync <config-file>

Sync creates symbolic links between 'source:target' pairs defined in
the configuration spec.
`

// Sync represents the sync command.
var Sync = makeSyncCmd()

func makeSyncCmd() cli.Command {
	fs := flag.NewFlagSet("bonclay sync", flag.ExitOnError)

	fs.Usage = func() {
		fmt.Println(strings.TrimSpace(syncUsage))
	}

	return cli.Command{
		Name: "sync",
		Desc: syncDesc,
		Action: func(args []string) {
			fs.Parse(args)
			args = fs.Args()

			if len(args) != 1 {
				fs.Usage()
				os.Exit(1)
			}

			cfg := mistertwo.NewConfiguration(args[0])
			syncTask(cfg)
		},
	}
}

// syncTask uses 'source:target' pairs defined in the configuration spec
// to link the sources to the targets.
func syncTask(config *mistertwo.Configuration) {
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
