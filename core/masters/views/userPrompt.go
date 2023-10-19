package Views

import (
	"fmt"
	"strings"
	"time"

	"golang.org/x/term"

	Execute "Rain/core/config/views/user"
	External "Rain/core/functions/external"
	Replireadings "Rain/core/functions/rep"
	APIAttacks "Rain/core/masters/attacks"
	Commands "Rain/core/masters/commands"
	Sessions "Rain/core/masters/sessions"
	Tools "Rain/tools"
)

func Prompt(session *Sessions.Session) {

	_, error := Execute.Execute_Standard("home-splash", session.User, session.Channel, true, false)
	if error != nil {
		return
	}

	prompt := term.NewTerminal(session.Channel, "")
	for {
		PromptBuild, lenline := Execute.PromptBuild("prompt", session)

		Fade, _, _ := Replireadings.GetFunctions("prompt")
		if !Fade {
			var PassedSplit bool = false
			for U := 0; U < lenline-1; U++ {
				if PassedSplit {
					session.Channel.Write([]byte(PromptBuild[U] + "\r\n"))
				}

				if strings.Contains(PromptBuild[U], "=============== SPLIT ===============") {
					PassedSplit = true
				}
			}
			_, error := session.Channel.Write([]byte(PromptBuild[lenline-1]))
			if error != nil {
				return
			}
		} else {
			var PassedSplit bool = false

			for U := 0; U < lenline-1; U++ {

				if PassedSplit {
					Out, _ := Tools.Fade(PromptBuild[U], session)
					session.Channel.Write([]byte(Out + "\r\n"))
				}

				if strings.Contains(PromptBuild[U], "=============== SPLIT ===============") {
					PassedSplit = true
				}

			}

			GradMain, _ := Tools.Fade(PromptBuild[lenline-1], session)
			_, error := session.Channel.Write([]byte(GradMain + "\x1b[38;5;254m"))
			if error != nil {
				return
			}
		}
		command, error := prompt.ReadLine()
		if error != nil {
			Execute.Execute_CustomTerm("exit-splash", session.User, session.Channel, true, nil)
			time.Sleep(1 * time.Second)
			session.Channel.Close()
			return
		}

		Execute_command(command, session)
	}
}

func Execute_command(command string, session *Sessions.Session) {

	Command := strings.Split(command, " ")

	if Command[0] == "" {
		return
	}

	Command_interface_Main := Commands.Get_Name(strings.ToLower(Command[0]))
	if Command_interface_Main == nil {
		CommandingOut := External.Command[strings.ToLower(Command[0])]
		if CommandingOut != nil {

			if CommandingOut.Admin && session.User.Admin {
				Execute.ExecuteString(strings.Join(CommandingOut.Banner, "\r\n"), session.User, session.Channel, true, false)
				return
			} else if CommandingOut.Reseller && session.User.Reseller || session.User.Admin {
				Execute.ExecuteString(strings.Join(CommandingOut.Banner, "\r\n"), session.User, session.Channel, true, false)
				return
			} else if CommandingOut.VIP || session.User.Reseller || session.User.Admin {
				Execute.ExecuteString(strings.Join(CommandingOut.Banner, "\r\n"), session.User, session.Channel, true, false)
				return
			} else if !CommandingOut.Admin && !CommandingOut.Reseller && !CommandingOut.VIP {
				Execute.ExecuteString(strings.Join(CommandingOut.Banner, "\r\n"), session.User, session.Channel, true, false)
				return
			}
		}
		Method := APIAttacks.Get(strings.ToLower(Command[0]))
		if Method == nil {
			custommap := map[string]string{
				"command": Command[0],
			}
			Execute.Execute_CustomTerm("command-403", session.User, session.Channel, true, custommap)
			return
		}

		APIAttacks.CreateAttack(Command, session)
		return
	}

	var Allow = Commands.Permissions{
		Admin:    session.User.Admin,
		Reseller: session.User.Reseller,
		VIP:      session.User.VIP,
		Raw:      session.User.Raw,
		Holder:   session.User.Holder,
	}

	Allowed := Command_interface_Main.Allowed_Permissions(&Allow)
	if Allowed {
		error := Command_interface_Main.CommandExecution(session, Command)
		if error != nil {
			if session.User.Admin {
				session.Channel.Write([]byte("error: " + fmt.Sprint(error) + "\r\n"))
				return
			}
			session.Channel.Write([]byte("failed to execute that command correctly"))
			return
		}
	} else {
		Execute.Execute_CustomTerm("command-404", session.User, session.Channel, true, nil)
	}

}
