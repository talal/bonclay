package task

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const exampleConfig = `
# For a detailed description on the different options and how to configure them,
# refer to the user guide: https://github.com/talal/bonclay/blob/master/doc/guide.md
backup:
  overwrite: false

restore:
  overwrite: false

sync:
  clean: true
  overwrite: true

spec:
  # ~/examplefile: file
  # ../../examplefile: ../file
  # ~/example dir: dir
  # ../example dir/some other dir: ../new dir
`

// Init creates a sample config file (bonclay.conf.yaml) in the current
// directory.
func Init() {
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
