package commands

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/talal/go-bits/cli"
)

const initDesc = `Create a new config file in the current directory.`

const initUsage = `
Usage: bonclay init

Create a new config file in the current directory with sane defaults and
example spec section.
`

const exampleConfig = `
# For a detailed description on the different options, take a look at the
# user guide: https://github.com/talal/bonclay/blob/master/doc/guide.md
backup:
  overwrite: false

restore:
  overwrite: false

sync:
  clean: true
  overwrite: true

spec:
  # ~/examplefile: file
  # ~/exampledir: dir
  # ~/dir/another one: another one
`

// Init represents the init command.
var Init = makeInitCmd()

func makeInitCmd() cli.Command {
	fs := flag.NewFlagSet("bonclay init", flag.ExitOnError)

	fs.Usage = func() {
		fmt.Println(strings.TrimSpace(initUsage))
	}

	return cli.Command{
		Name: "init",
		Desc: initDesc,
		Action: func(args []string) {
			fs.Parse(args)
			args = fs.Args()

			if len(args) > 0 {
				fs.Usage()
				os.Exit(1)
			}

			initTask()
		},
	}
}

// initTask creates a sample config file (bonclay.conf.yaml) in the current
// directory.
func initTask() {
	cwd, err := os.Getwd()
	if err != nil {
		cwd = os.Getenv("PWD")
	}

	cfgPath := filepath.Join(filepath.Clean(cwd), "bonclay.conf.yaml")

	_, err = os.Lstat(cfgPath)
	if err == nil {
		fmt.Fprintf(os.Stderr, "bonclay: error: config file already exists: %s\n", cfgPath)
		os.Exit(1)
	}

	f, err := os.Create(cfgPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "bonclay: error: could not create config file: %s\n", cfgPath)
		os.Exit(1)
	}
	defer f.Close()

	_, err = f.WriteString(strings.TrimSpace(exampleConfig))
	if err != nil {
		fmt.Fprintf(os.Stderr, "bonclay: error: could not create config file: %s\n", cfgPath)
		os.Exit(1)
	}
	f.Sync()

	fmt.Printf("bonclay: config file created: %s\n", cfgPath)
}
