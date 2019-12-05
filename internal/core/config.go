package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

// Configuration contains all the data from a configuration file.
type Configuration struct {
	Backup struct {
		Overwrite bool `yaml:"overwrite"`
	} `yaml:"backup"`
	Restore struct {
		Overwrite bool `yaml:"overwrite"`
	} `yaml:"restore"`
	Sync struct {
		Clean     bool `yaml:"clean"`
		Overwrite bool `yaml:"overwrite"`
	} `yaml:"sync"`
	// Spec is a map of source to target, where source/target are
	// the path to a file or a directory.
	Spec map[string]string `yaml:"spec"`
}

// NewConfiguration reads and validates a configuration file.
// Errors are written to os.Stderr and will result in program termination.
func NewConfiguration(path string) (config *Configuration) {
	path = strings.Replace(path, "~", os.Getenv("HOME"), 1)

	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "bonclay: error: could not load config file: %v\n", err)
		os.Exit(1)
	}

	err = yaml.Unmarshal(b, &config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "bonclay: error: could not parse config file: %v\n", err)
		os.Exit(1)
	}

	if len(config.Spec) == 0 {
		fmt.Fprintln(os.Stderr, "error: invalid config file: no files/directories specified in the config file's spec.")
		os.Exit(1)
	}

	return
}
