package subcommands

import (
	"strconv"
	"strings"

	ParseJson "Rain/core/functions/json"
	Sessions "Rain/core/masters/sessions"
)

func MessageSession(session *Sessions.Session, cmd []string) error {

	if len(cmd) <= 2 {
		session.Channel.Write([]byte("Invalid Syntax\r\n"))
		session.Channel.Write([]byte("Syntax: sessions message <<Message>> <Username>>@<<ID>>\r\n"))
		session.Channel.Write([]byte("Example: sessions message hello! root@1\r\n"))
		return nil
	}

	var Message []string

	for U := 2; U < len(cmd); U++ {
		CommandDebug := strings.Split(cmd[U], "@")
		if len(CommandDebug) <= 1 {
			Message = append(Message, cmd[U])
			continue
		}

		Turn, error := strconv.Atoi(CommandDebug[1])
		if error != nil {
			session.Channel.Write([]byte("Failed to message \"" + CommandDebug[1] + "\" due to invaild string split\r\n"))
			continue
		}

		SessionID := Sessions.Sessions[int64(Turn)]
		if SessionID == nil {
			session.Channel.Write([]byte("Failed to find session for \"" + CommandDebug[0] + "\" with id of \"" + CommandDebug[1] + "\"\r\n"))
			continue
		}

		if SessionID.Key == session.Key {
			session.Channel.Write([]byte(ParseJson.ConfigParse.App.AppName + " -> You can't message yourself\r\n"))
			continue
		}

		SessionID.Channel.Write([]byte("\x1b[0m\x1b7\x1b[1A\r\x1b[2K \x1b[38;5;15mPrivate Message from " + session.User.Username + "\x1b[38;5;105m>\x1b[38;5;15m " + strings.Join(Message, " ") + "\x1b[0m\x1b8"))
		continue

	}

	session.Channel.Write([]byte(ParseJson.ConfigParse.App.AppName + " -> Message has been sent to user(s)\r\n"))

	return nil

}
