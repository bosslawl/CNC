package subcommands

import (
	"strings"
	"strconv"

	Sessions "Rain/core/masters/sessions"
)

func KickSession(session *Sessions.Session, cmd []string) error {

	if len(cmd) <= 2 {
		session.Channel.Write([]byte("Invalid Syntax\r\n"))
		session.Channel.Write([]byte("Syntax: sessions kick <Username>>@<<ID>>\r\n"))
		session.Channel.Write([]byte("Example: sessions kick root@1\r\n"))
		return nil
	}

	for U := 2; U < len(cmd); U++ {
		CommandDebug := strings.Split(cmd[U], "@")
		if len(CommandDebug) <= 1 {
			session.Channel.Write([]byte("Failed to kick \"" + CommandDebug[0] + "\" due to invaild string split\r\n"))
			continue
		}

		Turn, error := strconv.Atoi(CommandDebug[1])
		if error != nil {
			session.Channel.Write([]byte("Failed to kick \"" + CommandDebug[1] + "\" due to invaild string split\r\n"))
			continue
		}

		SessionID := Sessions.Sessions[int64(Turn)]
		if SessionID == nil {
			session.Channel.Write([]byte("Failed to find session for \"" + CommandDebug[0] + "\" with id of \"" + CommandDebug[1] + "\"\r\n"))
			continue
		}

		if SessionID.Key == session.Key {
			session.Channel.Write([]byte("Failed to find session for \"" + CommandDebug[0] + "\" with id of \"" + CommandDebug[1] + "\"\r\n"))
			continue
		}

		if SessionID.User.Username == CommandDebug[0] && SessionID.Key == int64(Turn) {

			error := SessionID.Channel.Close()
			if error != nil {
				SessionID.Conn.Close()
				session.Channel.Write([]byte("Failed to kick \"" + CommandDebug[0] + "\"\r\n"))
				continue
			} else {
				SessionID.Conn.Close()
				session.Channel.Write([]byte("Successfully kicked \"" + CommandDebug[0] + "\"\r\n"))
				continue
			}

		}
	}

	return nil

}