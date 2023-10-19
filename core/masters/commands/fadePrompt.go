package Commands

import (
	"strings"

	Replireadings "Rain/core/functions/rep"
	Execute "Rain/core/config/views/user"
	"Rain/core/masters/sessions"

	"golang.org/x/term"
)

func init() {
	Load_Commands(&Command_interface{
		Name:        "promptfade",
		Roles:       []string{"everyone"},
		Admin:       false,
		Reseller:    false,
		Description: "Change the fade of prompt",
		CommandExecution: func(session *Sessions.Session, cmd []string) error {

			if Active, _, _ := Replireadings.GetFunctions("prompt"); !Active {
				custommap := map[string]string{
					"command": cmd[0],
				}
				Execute.Execute_CustomTerm("command-403", session.User, session.Channel, true, custommap)
				return nil
			}

			ColourOneB := term.NewTerminal(session.Channel, "\x1b[38;5;255mColour One (Start of Fade)>\x1b[38;5;227m ")
			NewSesColourOne, error := ColourOneB.ReadLine()
			if error != nil {
				session.Channel.Write([]byte("\r\n"))
				return nil
			}

			ColourTwoB := term.NewTerminal(session.Channel, "\x1b[38;5;255mColour Two (End of Fade)>\x1b[38;5;227m ")
			NewSesColourTwo, error := ColourTwoB.ReadLine()
			if error != nil {
				session.Channel.Write([]byte("\r\n"))
				return nil
			}

			session.ColourOne = strings.Trim(NewSesColourOne, "#")
			session.ColourTwo = strings.Trim(NewSesColourTwo, "#")

			return nil
		},
	})
}
