package subcommands

import (
	Database "Rain/core/database"
	Sessions "Rain/core/masters/sessions"
	ParseJson "Rain/core/functions/json"
	Execute "Rain/core/config/views/user"
)

func PowerSavingTrue(session *Sessions.Session, cmd []string) error {

	
	if len(cmd) <= 2 {
		session.Channel.Write([]byte("Invalid Syntax\r\n"))
		session.Channel.Write([]byte("Syntax: users powersaving=true <<Username>>\r\n"))
		session.Channel.Write([]byte("Example: users powersaving=true root\r\n"))
		return nil
	}

	for LengthControl := 2; LengthControl < len(cmd); LengthControl++ {
		User, boolen := Database.GetUser(cmd[LengthControl])
		if User == nil || !boolen {
			Execute.Execute_CustomTerm("cannot-find-user", session.User, session.Channel, true, nil)
			continue
		}

		if User.Powersaving {
			session.Channel.Write([]byte("\x1b[38;5;15m" + ParseJson.ConfigParse.App.AppName + " -> \"" + User.Username + "\" is already registered as a powersaving user\r\n"))
			continue
		}

		Row := Database.EditFeild(cmd[LengthControl], "Powersaving", "1", false)
		if !Row {
			session.Channel.Write([]byte("\x1b[38;5;15m" + ParseJson.ConfigParse.App.AppName + " -> Failed to update users Powersaving Status\r\n"))
			continue
		} else {
			session.Channel.Write([]byte("\x1b[38;5;15m\"\x1b[38;5;2m" + User.Username + "\x1b[38;5;15m\" is now a Powersaving User\x1b[38;5;15m\r\n"))
		}

		for _, I := range Sessions.Sessions {
			if I.User.Username == User.Username {
				I.Channel.Write([]byte("\x1b[0m\x1b7\x1b[1A\r\x1b[2K"))
				Execute.Execute_CustomTerm("user-powersaving-enabled", I.User, I.Channel, true, nil)
				I.Channel.Write([]byte("\x1b[0m\x1b8"))
				I.User.Powersaving = true
				break
			}
		}

	}

	return nil

}