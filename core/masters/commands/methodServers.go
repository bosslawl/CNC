package Commands

import (
	"fmt"
	"strconv"
	"strings"

	Database "Rain/core/database"
	ParseINI "Rain/core/functions/ini"
	ParseJson "Rain/core/functions/json"
	Util "Rain/core/functions/util"
	Attack_Groups "Rain/core/functions/util/attackGroups"
	Sessions "Rain/core/masters/sessions"

	"github.com/alexeyco/simpletable"
)

func init() {
	Load_Commands(&Command_interface{
		Name:        "servers",
		Roles:       []string{"everyone"},
		Admin:       false,
		Reseller:    false,
		Description: "View servers information",
		CommandExecution: func(session *Sessions.Session, cmd []string) error {
			// Convert the command to lower case
			for i := range cmd {
				cmd[i] = strings.ToLower(cmd[i])
			}

			if len(cmd) < 2 {
				session.Channel.Write([]byte("\x1b[0m"))
				table := simpletable.New()
				table.Header = &simpletable.Header{
					Cells: []*simpletable.Cell{
						{Align: simpletable.AlignCenter, Text: "#"},
						{Align: simpletable.AlignCenter, Text: "Name"},
						{Align: simpletable.AlignCenter, Text: "Likes:Dislikes"},
						{Align: simpletable.AlignCenter, Text: "Methods"},
						{Align: simpletable.AlignCenter, Text: "Running"},
					},
				}

				count := 0
				for Name, More := range Attack_Groups.Attk_groups {
					count++

					Running, _ := Database.GetAPI_Running(Name)
					Pos, neg := Attack_Groups.Sort_Voting(Name)
					r := []*simpletable.Cell{
						{Align: simpletable.AlignCenter, Text: strconv.Itoa(count)},
						{Align: simpletable.AlignCenter, Text: Name},
						{Align: simpletable.AlignCenter, Text: strconv.Itoa(neg) + ":" + strconv.Itoa(Pos)},
						{Align: simpletable.AlignCenter, Text: strconv.Itoa(len(More.Methods))},
						{Align: simpletable.AlignCenter, Text: strconv.Itoa(Running)},
					}

					table.Body.Cells = append(table.Body.Cells, r)
				}

				ParseINI.ParseTableServers(table)

				fmt.Fprint(session.Channel, "")
				fmt.Fprintln(session.Channel, strings.ReplaceAll(table.String(), "\n", "\r\n"))
				fmt.Fprint(session.Channel, "\r")

				return nil
			}

			switch cmd[1] {

			case "like":
				if len(cmd) <= 2 {
					session.Channel.Write([]byte(ParseJson.ConfigParse.App.AppName + " -> Command Example: servers like <api name>\r\n"))
					return nil
				}

				Row := Attack_Groups.Attk_groups[cmd[2]]
				if Row == nil {
					session.Channel.Write([]byte(ParseJson.ConfigParse.App.AppName + " -> Failed to find that server\r\n"))
					return nil
				}

				Row.Voting = append(Row.Voting, Attack_Groups.Vote{
					Name:      session.User.Username,
					Type_vote: "postive",
				})

				session.Channel.Write([]byte(ParseJson.ConfigParse.App.AppName + " -> Voted for that server correctly\r\n"))

				return nil

			case "dislike":

				if len(cmd) <= 2 {
					session.Channel.Write([]byte(ParseJson.ConfigParse.App.AppName + " -> Command Example: servers dislike <servername>\r\n"))
					return nil
				}

				Row := Attack_Groups.Attk_groups[cmd[2]]
				if Row == nil {
					session.Channel.Write([]byte(ParseJson.ConfigParse.App.AppName + " -> Failed to find that server\r\n"))
					return nil
				}

				Row.Voting = append(Row.Voting, Attack_Groups.Vote{
					Name:      session.User.Username,
					Type_vote: "dislike",
				})

				session.Channel.Write([]byte(ParseJson.ConfigParse.App.AppName + " -> Voted for that server correctly\r\n"))

				return nil

			case "view":
				if len(cmd) <= 2 {
					session.Channel.Write([]byte(ParseJson.ConfigParse.App.AppName + " -> Command Example: servers methods <servername>\r\n"))
					return nil
				}

				methods_serve := Attack_Groups.Attk_groups[cmd[2]]
				if methods_serve == nil {
					session.Channel.Write([]byte(ParseJson.ConfigParse.App.AppName + " -> Failed to find that server\r\n"))
					return nil
				}

				session.Channel.Write([]byte("\x1b[0m"))
				table := simpletable.New()
				table.Header = &simpletable.Header{
					Cells: []*simpletable.Cell{
						{Align: simpletable.AlignCenter, Text: "#"},
						{Align: simpletable.AlignCenter, Text: "Name"},
						{Align: simpletable.AlignCenter, Text: "Description"},
						{Align: simpletable.AlignCenter, Text: "VIP"},
						{Align: simpletable.AlignCenter, Text: "Running"},
					},
				}

				count := 0
				for _, I := range methods_serve.Methods {
					Running, error := Database.GetAPI_RunningMethod(cmd[2], I.Name)
					if error != nil {
						continue
					}
					count++
					r := []*simpletable.Cell{
						{Align: simpletable.AlignCenter, Text: strconv.Itoa(count)},
						{Align: simpletable.AlignCenter, Text: I.Name},
						{Align: simpletable.AlignCenter, Text: I.Description},
						{Align: simpletable.AlignCenter, Text: Util.ColourizeBoolen(I.VIP, true)},
						{Align: simpletable.AlignCenter, Text: strconv.Itoa(Running) + "/" + strconv.Itoa(I.MaxConcurrents)},
					}

					table.Body.Cells = append(table.Body.Cells, r)
				}

				ParseINI.ParseTableServers(table)

				fmt.Fprint(session.Channel, "")
				fmt.Fprintln(session.Channel, strings.ReplaceAll(table.String(), "\n", "\r\n"))
				fmt.Fprint(session.Channel, "\r")
			}
			return nil
		},
	})
}
