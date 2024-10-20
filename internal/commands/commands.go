package commands

type Commands interface {
	New() Commands
	Extract(t string) Cmd
}

type Cmd interface {
	Command() string
	Message(t string) string
	Compiler() string
}

// New will define and initialize chosen implementation.
// Function call should take empty struct as an argument. Call should look like this:
//
//	cmds := commands.New(new(tgcommands.Commands))
func New(implementator Commands) Commands {
	return implementator.New()
}
