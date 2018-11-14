package main

import (
	"os"

	"github.com/alecthomas/kingpin"
	"github.com/talal/bonclay/pkg/mistertwo"
)

var (
	// defined by the Makefile at compile time
	version string

	app           = kingpin.New("bonclay", "A fast and minimal backup tool.")
	syncCmd       = app.Command("sync", "Sync creates symbolic links between 'source:target' pairs defined in the configuration spec.")
	syncConfig    = syncCmd.Arg("config-file", "Path to the configuration file (.yaml).").Required().String()
	backupCmd     = app.Command("backup", "Backup uses the `source:target' pairs defined in the configuration spec to copy the sources to the targets.")
	backupConfig  = backupCmd.Arg("config-file", "Path to the configuration file (.yaml).").Required().String()
	restoreCmd    = app.Command("restore", "Restore is the reverse of backup, it copies the targets to the sources.")
	restoreConfig = restoreCmd.Arg("config-file", "Path to the configuration file (.yaml).").Required().String()
)

func main() {
	app.Author("Muhammad Talal Anwar <talalanwar@protonmail.com>")
	app.Version("bonclay version " + version)
	app.VersionFlag.Short('v')
	app.HelpFlag.Short('h')

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case syncCmd.FullCommand():
		c := mistertwo.NewConfiguration(*syncConfig)
		mistertwo.SyncTask(c)
	case backupCmd.FullCommand():
		c := mistertwo.NewConfiguration(*backupConfig)
		mistertwo.BackupTask(c)
	case restoreCmd.FullCommand():
		c := mistertwo.NewConfiguration(*restoreConfig)
		mistertwo.RestoreTask(c)
	}
}
