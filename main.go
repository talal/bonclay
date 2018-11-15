package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/talal/bonclay/pkg/mistertwo"
)

var (
	// set by the Makefile at linking time
	version string

	versionFlag bool

	usage = strings.Replace(strings.Replace(strings.TrimSpace(`
bonclay <version>
A fast and minimal backup tool.

Usage:
\tbonclay <command> <config-file>

Flags:
\t-h, --help
\t\t\tShow usage information.
\t-v, --version
\t\t\tShow application version.

Commands:
\tsync <config-file>
\t\t\tSync creates symbolic links between 'source:target' pairs defined in the
\t\t\tconfiguration spec.

\tbackup <config-file>
\t\t\tBackup uses the 'source:target' pairs defined in the configuration spec
\t\t\tto copy the sources to the targets.

\trestore <config-file>
\t\t\tRestore is the reverse of backup, it copies the targets to the sources.
`), `\t`, "  ", -1), "<version>", version, -1)

	errStr = strings.Replace(strings.TrimSpace(`
error: The following required arguments were not provided:
\t\t<command> <config-file>

Usage:
\tbonclay <command> <config-file>

For more information try --help
`), `\t`, "  ", -1)
)

func init() {
	flag.BoolVar(&versionFlag, "version", false, "Show application version.")
	flag.BoolVar(&versionFlag, "v", false, "Show application version.")
}

func main() {
	flag.Usage = func() { printAndExit(usage, 0) }
	flag.Parse()

	if versionFlag {
		printAndExit("bonclay "+version, 0)
	}

	args := flag.Args()

	if len(args) != 2 {
		printAndExit(errStr, 1)
	}

	switch args[0] {
	case "sync":
		c := mistertwo.NewConfiguration(args[1])
		mistertwo.SyncTask(c)
	case "backup":
		c := mistertwo.NewConfiguration(args[1])
		mistertwo.BackupTask(c)
	case "restore":
		c := mistertwo.NewConfiguration(args[1])
		mistertwo.RestoreTask(c)
	default:
		printAndExit(fmt.Sprintf(
			"error: '%s' is not a valid command, try '--help' for more information", args[0]), 1)
	}
}

func printAndExit(str string, exitCode int) {
	fmt.Println(str)
	os.Exit(exitCode)
}
