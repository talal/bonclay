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

	usage = strings.Replace(strings.TrimSpace(`
bonclay <version>
A fast and minimal backup tool.

Usage:
  bonclay <command> <config-file>

Flags:
  -h, --help
      Show usage information.
  -v, --version
      Show application version.

Commands:
  sync <config-file>
      Sync creates symbolic links between 'source:target' pairs defined in the
      configuration spec.

  backup <config-file>
      Backup uses the 'source:target' pairs defined in the configuration spec
      to copy the sources to the targets.

  restore <config-file>
      Restore is the reverse of backup, it copies the targets to the sources.
`), "<version>", version, -1)

	errStr = strings.TrimSpace(`
error: The following required arguments were not provided:
    <command> <config-file>

Usage:
  bonclay <command> <config-file>

For more information try --help
`)
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
