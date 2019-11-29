package main

import (
	"os"

	"github.com/talal/bonclay/internal/core"
	"github.com/talal/bonclay/internal/task"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	// set by the Makefile at linking time
	version string

	app = kingpin.New("bonclay", "A fast and minimal backup tool")

	initCmd = app.Command("init", "Create a new config file in the current directory")

	backupCmd     = app.Command("backup", "Backup files/directories to their target location")
	backupCfgFile = backupCmd.Arg("config-file", "Path to the config file").Required().String()

	restoreCmd     = app.Command("restore", "Restore files/directories to their original location")
	restoreCfgFile = restoreCmd.Arg("config-file", "Path to the config file").Required().String()

	syncCmd     = app.Command("sync", "Sync files/directories")
	syncCfgFile = syncCmd.Arg("config-file", "Path to the config file").Required().String()
)

func main() {
	app.Version("bonclay version " + version)
	app.VersionFlag.Short('v')
	app.HelpFlag.Short('h')

	// parse all command-line args and flags
	cmd := kingpin.MustParse(app.Parse(os.Args[1:]))
	switch cmd {
	case initCmd.FullCommand():
		task.Init()
	case backupCmd.FullCommand():
		task.Backup(core.NewConfiguration(*backupCfgFile))
	case restoreCmd.FullCommand():
		task.Restore(core.NewConfiguration(*restoreCfgFile))
	case syncCmd.FullCommand():
		task.Sync(core.NewConfiguration(*syncCfgFile))
	}
}
