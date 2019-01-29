package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/talal/bonclay/pkg/commands"
	"github.com/talal/go-bits/cli"
)

// set by the Makefile at linking time
var version string

const appUsageTemplate = `
Usage: bonclay [OPTIONS] <COMMAND> [COMMAND ARGS...]

Bonclay is a fast and minimal backup tool.

Options:{{range .Options}}
  {{printf "%-12s %s" .Name .Desc}}{{end}}

Commands:{{range .Commands}}
  {{printf "%-12s %s" .Name .Desc}}{{end}}
{{"\n"}}
`

var commandRegistry = make(cli.CommandRegistry)

func init() {
	commandRegistry.Add(commands.Backup)
	commandRegistry.Add(commands.Restore)
	commandRegistry.Add(commands.Sync)
	commandRegistry.Add(commands.Init)
}

func main() {
	fs := flag.NewFlagSet("bonclay", flag.ExitOnError)
	versionFlag := fs.Bool("version", false, "Show application version.")

	fs.Usage = func() {
		var data struct {
			Commands []cli.Command
			Options  []cli.Command
		}
		data.Commands = commandRegistry.Commands()
		data.Options = []cli.Command{
			{Name: "--help", Desc: "Show this screen."},
			{Name: "--version", Desc: "Show application version."},
		}

		tmpl := template.Must(template.New("appUsage").Parse(strings.TrimSpace(appUsageTemplate)))
		tmpl.Execute(os.Stdout, data)
	}

	fs.Parse(os.Args[1:])

	if *versionFlag {
		fmt.Println("bonclay " + version)
		return
	}

	args := fs.Args()
	if len(args) == 0 {
		fs.Usage()
		os.Exit(1)
	}

	if cmd, ok := commandRegistry[args[0]]; !ok {
		fmt.Fprintf(os.Stderr, "bonclay: error: unknown command: %s\n", args[0])
		os.Exit(1)
	} else {
		cmd.Action(args[1:])
	}
}
