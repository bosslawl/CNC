package subcommands

import (
	"strings"
	"time"
	"strconv"

	Database "Rain/core/database"
	Sessions "Rain/core/masters/sessions"
	ParseJson "Rain/core/functions/json"
	Execute "Rain/core/config/views/user"
)

func UserAddDays(session *Sessions.Session, sep, cmd []string) error {

	if !strings.Contains(strings.Replace(strings.Join(sep, "="), sep[1], "", -1), "add-days=") || len(cmd) <= 2 {
		session.Channel.Write([]byte("Invalid Syntax\r\n"))
		session.Channel.Write([]byte("Syntax: users add-days=<<Integer>> <<Username>>\r\n"))
		session.Channel.Write([]byte("Example: users add-days=30 root\r\n"))
		return nil
	}

	TimeChange, error := strconv.Atoi(sep[1])
	if error != nil {
		session.Channel.Write([]byte("\x1b[38;5;15m" + ParseJson.ConfigParse.App.AppName + " -> \"" + sep[1] + "\" must be an int\r\n"))
		return nil
	}

	for LengthControl := 2; LengthControl < len(cmd); LengthControl++ {
		User, boolen := Database.GetUser(cmd[LengthControl])
		if User == nil || !boolen {
			Execute.Execute_CustomTerm("cannot-find-user", session.User, session.Channel, true, nil)
			continue
		}

		Row := Database.AddTime(cmd[LengthControl], (time.Hour*24)*time.Duration(TimeChange))

		if !Row {
			session.Channel.Write([]byte("\x1b[38;5;15m" + ParseJson.ConfigParse.App.AppName + " -> Failed to change users days left to \"" + sep[1] + "\"\r\n"))
			continue
		} else {
			session.Channel.Write([]byte("\x1b[38;5;15m\"\x1b[38;5;11m" + User.Username + "\x1b[38;5;15m\" expiry has been edited\x1b[38;5;15m\r\n"))
		}

		//Row := Database.EditFeild(cmd[LengthControl], "Expiry", strconv.Itoa(TimeChange))
		//if !Row {
		//	session.Channel.Write([]byte("\x1b[38;5;15m" + ParseJson.ConfigParse.App.AppName + " -> Failed to change users days left to \"" + sep[1] + "\"\r\n"))
		//	continue
		//} else {
		//	session.Channel.Write([]byte("\x1b[38;5;15m\"\x1b[38;5;11m" + User.Username + "\x1b[38;5;15m\" expiry has been edited\x1b[38;5;15m\r\n"))
		//}

		for _, I := range Sessions.Sessions {
			if I.User.Username == User.Username {
				I.Channel.Write([]byte("\x1b[0m\x1b7\x1b[1A\r\x1b[2K"))
				Execute.Execute_CustomTerm("user-expiry-changed", I.User, I.Channel, true, nil)
				I.Channel.Write([]byte("\x1b[0m\x1b8"))
				End := time.Unix(session.User.Expiry, 0)

				NewEnd := End.Add((time.Hour * 24) * time.Duration(TimeChange)).Unix()
				I.User.Expiry = NewEnd
				break
			}
		}
	}

	return nil

}