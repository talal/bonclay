/*
Package cli provides types and functions for creating simple command line
interfaces.

Here is one example of how it can be used.

main.go:
	package main

	import "github.com/myapp/commands"
	import "github.com/talal/go-bits/cli"

	var commandRegistry = make(cli.CommandRegistry)

	func init() {
		commandRegistry.Add(commands.CmdOne)
		commandRegistry.Add(commands.CmdTwo)
	}

	func main() {
		fs := flag.NewFlagSet("myapp", flag.ExitOnError)

		fs.Parse(os.Args[1:])

		args := fs.Args()
		if len(args) == 0 {
			fs.Usage()
			os.Exit(1)
		}

		if cmd, ok := commandRegistry[args[0]]; !ok {
			fmt.Fprintf(os.Stderr, "myapp: error: unknown command: %s\n", args[0])
			os.Exit(1)
		} else {
			cmd.Action(args[1:])
		}
	}


commands/cmdone.go:
	package commands

	import "github.com/talal/go-bits/cli"

	var CmdOne = makeCmdOne()

	func makeCmdOne() cli.Command {
		fs := flag.NewFlagSet("myapp cmdone", flag.ExitOnError)

		cmdOneDesc := "Command one does just that."

		return cli.Command{
			Name: "cmdone",
			Desc: cmdOneDesc,
			Action: func(args []string) {
				fs.Parse(args)
				args = fs.Args()

				if len(args) != 1 {
					fs.Usage()
					os.Exit(1)
				}

				cfg := somepackage.NewConfiguration(args[0])
				doSomething(cfg)
			},
		}
	}

	func doSomething(cfg *somePackage.SomeType) {}

commands/cmdtwo.go:
	package commands

	import "github.com/talal/go-bits/cli"

	var CmdTwo = makeCmdTwo()

	func makeCmdTwo() cli.Command {
		fs := flag.NewFlagSet("myapp cmdtwo", flag.ExitOnError)

		cmdTwoDesc := "Command two does just that."

		return cli.Command{
			Name: "cmdtwo",
			Desc: cmdTwoDesc,
			Action: func(args []string) {
				fs.Parse(args)
				args = fs.Args()

				if len(args) > 2 {
					fs.Usage()
					os.Exit(1)
				}

				doSomethingElse()
			},
		}
	}

	func doSomethingElse() {}
*/
package cli
