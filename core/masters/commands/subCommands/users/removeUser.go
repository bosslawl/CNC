package subcommands

import (
	"strings"

	Database "Rain/core/database"
	Sessions "Rain/core/masters/sessions"
	ParseJson "Rain/core/functions/json"
	Execute "Rain/core/config/views/user"

	"golang.org/x/term"
)

func RemoveUser(session *Sessions.Session, cmd []string) error {

	if len(cmd) <= 2 {
		session.Channel.Write([]byte("Invalid Syntax\r\n"))
		session.Channel.Write([]byte("Syntax: users remove <<Username>>\r\n"))
		session.Channel.Write([]byte("Example: users remove root\r\n"))
		return nil
	}

	session.Channel.Write([]byte("\x1b[38;5;15mAre you sure you want to remove that user? This cant be undone\r\n"))
	output := term.NewTerminal(session.Channel, "Y/N >")

	Choice, error := output.ReadLine()
	if error != nil || strings.ToLower(Choice) != "y" {
		return nil
	}

	for LengthControl := 2; LengthControl < len(cmd); LengthControl++ {
		User, boolen := Database.GetUser(cmd[LengthControl])
		if User == nil || !boolen {
			Execute.Execute_CustomTerm("cannot-find-user", session.User, session.Channel, true, nil)
			continue
		}

		if User.Admin {
			session.Channel.Write([]byte("\x1b[38;5;15m" + ParseJson.ConfigParse.App.AppName + " -> User has administrative permissions.\r\n"))
			continue
		}

		error := Database.Remove(cmd[LengthControl])
		if !error {
			session.Channel.Write([]byte("\x1b[38;5;15m" + ParseJson.ConfigParse.App.AppName + " -> Failed to remove user from Database\r\n"))
			continue
		} else {
			session.Channel.Write([]byte("\x1b[38;5;15m\"\x1b[38;5;1m" + User.Username + "\x1b[38;5;15m\" has been removed from the Database\x1b[38;5;15m\r\n"))
			continue
		}
	}

	return nil

}