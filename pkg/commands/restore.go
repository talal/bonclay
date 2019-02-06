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

const restoreDesc = `Restore files/directories to their original location.`

const restoreUsage = `
Usage: bonclay restore <config-file>

Restore is the reverse of backup, it uses the 'source:target' pairs
defined in the configuration spec to copy the targets to the sources.
`

// Restore represents the restore command.
var Restore = makeRestoreCmd()

func makeRestoreCmd() cli.Command {
	fs := flag.NewFlagSet("bonclay restore", flag.ExitOnError)

	fs.Usage = func() {
		fmt.Println(strings.TrimSpace(restoreUsage))
	}

	return cli.Command{
		Name: "restore",
		Desc: restoreDesc,
		Action: func(args []string) {
			fs.Parse(args)
			args = fs.Args()

			if len(args) != 1 {
				fs.Usage()
				os.Exit(1)
			}

			cfg := mistertwo.NewConfiguration(args[0])
			restoreTask(cfg)
		},
	}
}

// restoreTask uses 'source:target' pairs defined in the configuration spec
// to copy the targets to the sources. It is the reverse of Backup.
func restoreTask(config *mistertwo.Configuration) {
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
