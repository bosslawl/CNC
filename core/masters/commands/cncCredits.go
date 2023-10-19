package Commands

import 	"Rain/core/masters/sessions"

func init() {
	Load_Commands(&Command_interface{
		Name:        "credits",
		Roles:       []string{"everyone"},
		Admin:       false,
		Reseller:    false,
		Description: "Displays CNC credits",
		CommandExecution: func(session *Sessions.Session, cmd []string) error {

			session.Channel.Write([]byte("Rain is a modular cnc that is designed to be easy to use.\r\n"))
			session.Channel.Write([]byte("Rain is written in Go and is designed to be cross platform.\r\n"))
			session.Channel.Write([]byte("Paragon was originally made by FB in golang for private use.\r\n"))
			session.Channel.Write([]byte("Paragon was later rewritten by boss & athena into Rain.\r\n"))
			session.Channel.Write([]byte("Rain is now a project that is now maintained soely by boss & athena.\r\n"))
			
			return nil
		},
	})
}
