package subcommands

import (
	Execute "Rain/core/config/views/user"
	Database "Rain/core/database"
	ParseJson "Rain/core/functions/json"
	Sessions "Rain/core/masters/sessions"
)

func VIPFalse(session *Sessions.Session, cmd []string) error {

	if len(cmd) <= 2 {
		session.Channel.Write([]byte("Invalid Syntax\r\n"))
		session.Channel.Write([]byte("Syntax: users vip=false <<Username>>\r\n"))
		session.Channel.Write([]byte("Example: users vip=false root\r\n"))
		return nil
	}

	for LengthControl := 2; LengthControl < len(cmd); LengthControl++ {
		User, boolen := Database.GetUser(cmd[LengthControl])
		if User == nil || !boolen {
			Execute.Execute_CustomTerm("cannot-find-user", session.User, session.Channel, true, nil)
			continue
		}

		if !User.VIP {
			session.Channel.Write([]byte("\x1b[38;5;15m" + ParseJson.ConfigParse.App.AppName + " -> \"" + User.Username + "\" is already registered as a Non-VIP user\r\n"))
			continue
		}

		Row := Database.EditFeild(cmd[LengthControl], "VIP", "0", false)
		if !Row {
			session.Channel.Write([]byte("\x1b[38;5;15m" + ParseJson.ConfigParse.App.AppName + " -> Failed to update users VIP status\r\n"))
			continue
		} else {
			session.Channel.Write([]byte("\x1b[38;5;15m\"\x1b[38;5;1m" + User.Username + "\x1b[38;5;15m\" is no longer registered as VIP\x1b[0m\r\n"))
		}

		for _, I := range Sessions.Sessions {
			if I.User.Username == User.Username {
				I.Channel.Write([]byte("\x1b[0m\x1b7\x1b[1A\r\x1b[2K"))
				Execute.Execute_CustomTerm("user-vip-demoted", I.User, I.Channel, true, nil)
				I.Channel.Write([]byte("\x1b[0m\x1b8"))
				I.User.VIP = false
				break
			}
		}

	}

	return nil
}
