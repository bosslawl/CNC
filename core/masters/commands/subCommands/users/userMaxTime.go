package subcommands

import (
	"strings"
	"strconv"

	Database "Rain/core/database"
	Sessions "Rain/core/masters/sessions"
	ParseJson "Rain/core/functions/json"
	Execute "Rain/core/config/views/user"
)

func UserMaxTime(session *Sessions.Session, sep, cmd []string) error {

	if !strings.Contains(strings.Replace(strings.Join(sep, "="), sep[1], "", -1), "maxtime=") || len(cmd) <= 2 {
		session.Channel.Write([]byte("Invalid Syntax\r\n"))
		session.Channel.Write([]byte("Syntax: users maxtime=<<Integer>> <<Username>>\r\n"))
		session.Channel.Write([]byte("Example: users maxtime=3600 root\r\n"))
		return nil
	}

	TimeChange, error := strconv.Atoi(sep[1])
	if error != nil || TimeChange > 172800 {
		session.Channel.Write([]byte("\x1b[38;5;15m" + ParseJson.ConfigParse.App.AppName + " -> \"" + sep[1] + "\" must be an int or between 0 and 172800\r\n"))
		return nil
	}

	for LengthControl := 2; LengthControl < len(cmd); LengthControl++ {
		User, boolen := Database.GetUser(cmd[LengthControl])
		if User == nil || !boolen {
			Execute.Execute_CustomTerm("cannot-find-user", session.User, session.Channel, true, nil)
			continue
		}

		if User.MaxTime == TimeChange {
			session.Channel.Write([]byte("\x1b[38;5;15m" + ParseJson.ConfigParse.App.AppName + " -> \"" + User.Username + "\" Duration is already set to \"" + sep[1] + "\"\r\n"))
			continue
		}

		Row := Database.EditFeild(cmd[LengthControl], "MaxTime", strconv.Itoa(TimeChange), true)
		if !Row {
			session.Channel.Write([]byte("\x1b[38;5;15m" + ParseJson.ConfigParse.App.AppName + " -> Failed to change users duration to \"" + sep[1] + "\"\r\n"))
			continue
		} else {
			session.Channel.Write([]byte("\x1b[38;5;15m\"\x1b[38;5;11m" + User.Username + "\x1b[38;5;15m\" duration has been changed to \"" + sep[1] + "\" from \"" + strconv.Itoa(User.MaxTime) + "\"\x1b[38;5;15m\r\n"))
		}

		for _, I := range Sessions.Sessions {
			if I.User.Username == User.Username {
				I.Channel.Write([]byte("\x1b[0m\x1b7\x1b[1A\r\x1b[2K"))
				Execute.Execute_CustomTerm("user-maxtime-changed", I.User, I.Channel, true, nil)
				I.Channel.Write([]byte("\x1b[0m\x1b8"))
				I.User.MaxTime = TimeChange
				break
			}
		}

	}

	return nil

}