package subcommands

import (
	"strings"

	Database "Rain/core/database"
	Sessions "Rain/core/masters/sessions"
	ParseJson "Rain/core/functions/json"
	Execute "Rain/core/config/views/user"

	"golang.org/x/term"
)

func BanUser(session *Sessions.Session, cmd []string) error {
	if len(cmd) <= 2 {
		session.Channel.Write([]byte("Invalid Syntax\r\n"))
		session.Channel.Write([]byte("Syntax: users ban <<Username>>\r\n"))
		session.Channel.Write([]byte("Example: users ban root\r\n"))
		return nil
	}

	session.Channel.Write([]byte("\x1b[38;5;15mAre you sure you want to ban that user?\r\n"))
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

		if User.Banned {
			session.Channel.Write([]byte("\x1b[38;5;15m" + ParseJson.ConfigParse.App.AppName + " -> \"" + User.Username + "\" is already banned\r\n"))
			continue
		}

		Row := Database.EditFeild(cmd[LengthControl], "Banned", "1", false)
		if !Row {
			session.Channel.Write([]byte("\x1b[38;5;15m" + ParseJson.ConfigParse.App.AppName + " -> Failed to ban the user\r\n"))
			continue
		} else {
			session.Channel.Write([]byte("\x1b[38;5;15m\"\x1b[38;5;1m" + User.Username + "\x1b[38;5;15m\" is now banned\x1b[0m\r\n"))
		}

		for _, I := range Sessions.Sessions {
			if I.User.Username == User.Username {
				I.Channel.Write([]byte("\x1b[0m\x1b7\x1b[1A\r\x1b[2K"))
				Execute.Execute_CustomTerm("user-banned", I.User, I.Channel, true, nil)
				I.Channel.Write([]byte("\x1b[0m\x1b8"))
				I.Channel.Close()
				I.User.Banned = true
				break
			}
		}
	}

	return nil

}