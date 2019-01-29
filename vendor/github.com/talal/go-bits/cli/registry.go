// Copyright for some parts of the code are held by
// Genadi Samokovarov <https://github.com/gsamokovarov>
// under the MIT License (MIT)

package cli

import "sort"

// CommandRegistry contains all the application's commands.
type CommandRegistry map[string]Command

// Add adds a command to a CommandRegistry.
func (r *CommandRegistry) Add(c Command) {
	(*r)[c.Name] = c
}

// Commands returns all of the registered commands, sorted alphabetically.
func (r *CommandRegistry) Commands() []Command {
	var commands []Command

	for _, cmdName := range r.sortedKeys() {
		commands = append(commands, (*r)[cmdName])
	}

	return commands
}

func (r *CommandRegistry) sortedKeys() []string {
	var keys []string

	for key := range *r {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	return keys
}
