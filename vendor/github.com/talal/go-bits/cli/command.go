// Copyright for some parts of the code are held by
// Genadi Samokovarov <https://github.com/gsamokovarov>
// under the MIT License (MIT)

package cli

// CommandFn represents a command handler function.
type CommandFn func(args []string)

// Command represents a command line action.
type Command struct {
	Name   string
	Desc   string
	Action CommandFn
}
