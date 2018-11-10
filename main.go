/**********************************************************************************
*
* Copyright 2018 Muhammad Talal Anwar <talalanwar@protonmail.com>
*
* This program is free software: you can redistribute it and/or modify it under
* the terms of the GNU General Public License as published by the Free Software
* Foundation, either version 3 of the License, or (at your option) any later
* version.
*
* This program is distributed in the hope that it will be useful, but WITHOUT ANY
* WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR
* A PARTICULAR PURPOSE. See the GNU General Public License for more details.
*
* You should have received a copy of the GNU General Public License along with
* this program. If not, see <http://www.gnu.org/licenses/>.
*
**********************************************************************************/

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
