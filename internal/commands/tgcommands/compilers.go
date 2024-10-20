package tgcommands

import (
	"fmt"
	"regexp"
	"tgapi/internal/commands"
)

const (
	generateCMDName = `generate`
	helpCMDName     = `help`
)

const (
	substringPosName int = iota + 1
	substringPosMessageStart
)

var mutualCommandCompiler = regexp.MustCompile(`(?i)^/(\w+)(.*)$`)

func NewCommandFromName(name string) commands.Cmd {
	return BotCommand{
		command: name,
		regexp:  compileRegexpFromName(name),
	}
}

func compileRegexp(s string) *regexp.Regexp {
	return regexp.MustCompile(s)
}

func compileRegexpFromName(s string) *regexp.Regexp {
	template := "(?i)^\\/(%s) (.*)$"
	return regexp.MustCompile(
		fmt.Sprintf(
			template,
			s,
		),
	)
}

func ExtractSubstrings(s string) []string {
	substrings := mutualCommandCompiler.FindStringSubmatch(s)
	//fmt.Println(substrings, "SUBSTRINGS")

	return substrings
}
