package tgcommands_test

import (
	"testing"
	"tgapi/internal/commands/tgcommands"
)

func TestCommands_New(t *testing.T) {
	var cmds tgcommands.Commands

	cmds.New()
}

func TestCommands_ExtractSubstrings(t *testing.T) {
	str := "/generate"

	substrings := tgcommands.ExtractSubstrings(str)

	t.Log(substrings)
}

func TestCommands_ExtractCommands(t *testing.T) {
	commd := tgcommands.NewCommandFromName("generate")

	t.Log(commd.Command())
	t.Log(commd.Message("/generate"))
	t.Log(commd.Compiler())
}
