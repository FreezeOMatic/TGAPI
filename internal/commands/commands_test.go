package commands_test

import (
	"testing"
	"tgapi/internal/commands"
	"tgapi/internal/commands/tgcommands"
)

func TestNew(t *testing.T) {

	tgCommands := commands.New(new(tgcommands.Commands))

	t.Log(tgCommands.Extract("/generate me some asnwer ti wquestui ga gaa gag"))
}
