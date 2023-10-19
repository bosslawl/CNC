package Commands

import (
	"strings"

	Sessions "Rain/core/masters/sessions"
)

func init() {
	Load_Commands(&Command_interface{
		Name:        "broadcast",
		Roles:       []string{"admin"},
		Admin:       true,
		Reseller:    false,
		Description: "Broadcast a message to open sessions",
		CommandExecution: func(session *Sessions.Session, cmd []string) error {

			if len(cmd) < 2 {
				session.Channel.Write([]byte("Invalid Syntax\r\n"))
				session.Channel.Write([]byte("Syntax: broadcast <<Message>>\r\n"))
				session.Channel.Write([]byte("Example: broadcast Hello!\r\n"))
				return nil
			}

			message := strings.Join(cmd, " ")

			message = strings.Replace(message, "broadcast", "", -1)

			Sessions.Broadcast([]byte("\x1b[0m\x1b7\x1b[1A\r\x1b[2K\x1b[38;5;15m" + session.User.Username + "\x1b[38;5;105m>\x1b[38;5;15m " + message + "\x1b[0m\x1b8"))
			return nil
		},
	})
}
