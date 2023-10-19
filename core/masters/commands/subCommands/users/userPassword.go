package subcommands

import (
	"strings"

	Database "Rain/core/database"
	Sessions "Rain/core/masters/sessions"
	ParseJson "Rain/core/functions/json"
	Execute "Rain/core/config/views/user"
	Util "Rain/core/functions/util"
)

func ChangeUserPassword(session *Sessions.Session, sep, cmd []string) error {

	if !strings.Contains(strings.Replace(strings.Join(sep, "="), sep[1], "", -1), "password=") || len(cmd) <= 2 {
		session.Channel.Write([]byte("Invalid Syntax\r\n"))
		session.Channel.Write([]byte("Syntax: users password=<<Passowrd>> <<Username(s)>>\r\n"))
		session.Channel.Write([]byte("Example: users password=abcdefghijklmnopqrstuvwxyz root\r\n"))
		return nil
	}

	for LengthControl := 2; LengthControl < len(cmd); LengthControl++ {
		User, boolen := Database.GetUser(cmd[LengthControl])
		if User == nil || !boolen {
			Execute.Execute_CustomTerm("cannot-find-user", session.User, session.Channel, true, nil)
			continue
		}

		if User.Password == sep[1] {
			session.Channel.Write([]byte("\x1b[38;5;15m" + ParseJson.ConfigParse.App.AppName + " -> \"" + sep[1] + "\" Password is already set to that!\r\n"))
			continue
		}
		pwd := Util.PasswordHash(sep[1])
		Row := Database.EditFeild(cmd[LengthControl], "Password", pwd, false)
		if !Row {
			session.Channel.Write([]byte("\x1b[38;5;15m" + ParseJson.ConfigParse.App.AppName + " -> Failed to change users password correctly\r\n"))
			continue
		} else {
			session.Channel.Write([]byte("\x1b[38;5;15m\"\x1b[38;5;105m" + User.Username + "\x1b[38;5;15m\" password has been changed\x1b[38;5;15m\r\n"))
		}
	}

	return nil

}