package Commands

import (
	Execute "Rain/core/config/views/user"
	"Rain/core/masters/sessions"
)

func init() {
	Load_Commands(&Command_interface{
		Name:        "cls",
		Roles:       []string{"everyone"},
		Admin:       false,
		Reseller:    false,
		Description: "Simple clear terminal screen command",
		CommandExecution: func(session *Sessions.Session, cmd []string) error {

			_, error := Execute.Execute_Standard("clear-splash", session.User, session.Channel, true, false)
			if error != nil {
				return error
			}
			return nil
		},
	})
}
