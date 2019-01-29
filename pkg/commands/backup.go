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

const backupDesc = `Backup files/directories to their target location.`

const backupUsage = `
Usage: bonclay backup <config-file>

Backup uses the 'source:target' pairs defined in the configuration spec
to copy the sources to the targets.
`

// Backup represents the backup command.
var Backup = makeBackupCmd()

func makeBackupCmd() cli.Command {
	fs := flag.NewFlagSet("bonclay backup", flag.ExitOnError)

	fs.Usage = func() {
		fmt.Println(strings.TrimSpace(backupUsage))
	}

	return cli.Command{
		Name: "backup",
		Desc: backupDesc,
		Action: func(args []string) {
			fs.Parse(args)
			args = fs.Args()

			if len(args) != 1 {
				fs.Usage()
				os.Exit(1)
			}

			cfg := mistertwo.NewConfiguration(args[0])
			backupTask(cfg)
		},
	}
}

// backupTask uses 'source:target' pairs defined in the configuration spec to copy
// the sources to the targets.
func backupTask(config *mistertwo.Configuration) {
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
