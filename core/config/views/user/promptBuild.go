package Execute

import (
	"bufio"
	"strings"

	viewsMapped "Rain/core/config/views"
	Sessions "Rain/core/masters/sessions"
)

func PromptBuild(Name string, session *Sessions.Session) ([]string, int) {

	scanner := bufio.NewScanner(strings.NewReader(viewsMapped.Branding[Name]))

	var Line []string

	for scanner.Scan() {
		Test := strings.Replace(scanner.Text(), "<<$username>>", session.User.Username, -1)
		Test = strings.Replace(Test, "\\x1b", "\x1b", -1)
		Test = strings.Replace(Test, "\\033", "\x1b", -1)
		Line = append(Line, Test)
	}

	return Line, len(Line)
}
