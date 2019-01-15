package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/talal/bonclay/pkg/mistertwo"
	"github.com/talal/bonclay/pkg/tasks"
)

// set by the Makefile at linking time
var version string

const usageInfo = `
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
      Restore is the reverse of backup, it uses the 'source:target' pairs defined
      in the configuration spec to copy the targets to the sources.
`

const errMsg = `
error: The following required arguments were not provided:
    <command> <config-file>

Usage:
  bonclay <command> <config-file>

For more information try --help
`

var versionFlag bool

func init() {
	flag.BoolVar(&versionFlag, "version", false, "Show application version.")
	flag.BoolVar(&versionFlag, "v", false, "Show application version.")
}

func main() {
	flag.Usage = func() { printAndExit(strings.TrimSpace(usageInfo), 0) }
	flag.Parse()

	if versionFlag {
		printAndExit("bonclay "+version, 0)
	}

	args := flag.Args()

	if len(args) != 2 {
		printAndExit(strings.TrimSpace(errMsg), 1)
	}

	switch args[0] {
	case "sync":
		c := mistertwo.NewConfiguration(args[1])
		tasks.Sync(c)
	case "backup":
		c := mistertwo.NewConfiguration(args[1])
		tasks.Backup(c)
	case "restore":
		c := mistertwo.NewConfiguration(args[1])
		tasks.Restore(c)
	default:
		printAndExit(fmt.Sprintf(
			"error: '%s' is not a valid command, try '--help' for more information", args[0]), 1)
	}
}

func printAndExit(str string, exitCode int) {
	fmt.Println(str)
	os.Exit(exitCode)
}
