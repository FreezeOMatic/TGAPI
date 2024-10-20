package tgcommands

import (
	"log"
	"regexp"
	"strings"
	"tgapi/internal/commands"
)

const tgBotActionCodeGenerate = 1

type Commands struct {
	help        commands.Cmd
	subcommands map[string]commands.Cmd
}

func (c Commands) New() commands.Commands {
	return Commands{
		subcommands: map[string]commands.Cmd{
			generateCMDName: NewCommandFromName(generateCMDName),
		},
		help: NewCommandFromName(helpCMDName),
	}
}

func (c Commands) Extract(msg string) commands.Cmd {

	commandString := ExtractSubstrings(msg)
	var (
		name string
	)
	// fmt.Println("NAME ", commandString[substringPosName])
	// fmt.Println("LEN ", len(commandString))

	//for i, k := range commandString {
	//	println("i ", i, "k ", k)
	//}

	if len(commandString) > 0 {
		name = strings.TrimLeft(commandString[substringPosName], "/")
	}

	var (
		cmd commands.Cmd
		ok  bool
	)

	if cmd, ok = c.subcommands[name]; !ok {
		log.Printf("[ERROR] unknown command: %s\n", name)
		return nil
	}

	return cmd
}

type BotCommand struct {
	command string
	regexp  *regexp.Regexp
}

func (c BotCommand) ActionCode() int {
	return tgBotActionCodeGenerate
}

func (c BotCommand) Command() string {
	return c.command
}

func (c BotCommand) Message(t string) string {
	commandString := ExtractSubstrings(t)

	var (
		message string
	)

	if len(commandString) == 0 {
		return message
	}

	if len(commandString) < 2 {
		return ""
	}

	return strings.Join(commandString[2:], " ")
}

func (c BotCommand) Compiler() string {
	return c.regexp.String()
}
