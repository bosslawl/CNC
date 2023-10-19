package Commands

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	Execute "Rain/core/config/views/user"
	Database "Rain/core/database"
	Util "Rain/core/functions/util"
	ParseINI "Rain/core/functions/ini"
	Sessions "Rain/core/masters/sessions"
	"Rain/core/masters/commands/subCommands/sessions"

	"github.com/alexeyco/simpletable"
)

func init() {
	Load_Commands(&Command_interface{
		Name:        "sessions",
		Roles:       []string{"admin"},
		Admin:       true,
		Reseller:    false,
		Description: "Simple lists all open sessions",
		CommandExecution: func(session *Sessions.Session, cmd []string) error {

			if len(cmd) < 2 {
				table := simpletable.New()
				table.Header = &simpletable.Header{
					Cells: []*simpletable.Cell{
						{Align: simpletable.AlignCenter, Text: "Username"},
						{Align: simpletable.AlignCenter, Text: "Reseller"},
						{Align: simpletable.AlignCenter, Text: "Admin"},
						{Align: simpletable.AlignCenter, Text: "Running"},
						{Align: simpletable.AlignCenter, Text: "IPv4"},
						{Align: simpletable.AlignCenter, Text: "Uptime"},
					},
				}

				for _, I := range Sessions.Sessions {
					Running, _ := Database.GetRunningUser(I.User.Username)
					r := []*simpletable.Cell{
						{Align: simpletable.AlignCenter, Text: I.User.Username},
						{Align: simpletable.AlignCenter, Text: Util.ColourizeBoolen(I.User.Reseller, true)},
						{Align: simpletable.AlignCenter, Text: Util.ColourizeBoolen(I.User.Admin, true)},
						{Align: simpletable.AlignCenter, Text: strconv.Itoa(Running) + "/" + strconv.Itoa(I.User.Concurrents)},
						{Align: simpletable.AlignCenter, Text: I.Conn.RemoteAddr().String()},
						{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%.2f mins", time.Since(I.Created).Minutes())},
					}

					table.Body.Cells = append(table.Body.Cells, r)
				}

				ParseINI.ParseTableSessions(table)

				fmt.Fprint(session.Channel, "")
				fmt.Fprintln(session.Channel, strings.ReplaceAll(table.String(), "\n", "\r\n"))
				fmt.Fprint(session.Channel, "\r")
				return nil
			}

			switch cmd[1] {

				case "ids":
					subcommands.SessionsIDs(session)
					return nil
				case "kick":
					subcommands.KickSession(session, cmd)
					return nil
				case "lookup":
					subcommands.LookupSession(session, cmd)
					return nil
				case "message":
					subcommands.MessageSession(session, cmd)
					return nil
			
			}

			custommap := map[string]string{
				"subcommand": cmd[1],
				"command":    cmd[0],
			}
			Execute.Execute_CustomTerm("subcommand-403", session.User, session.Channel, true, custommap)
			return nil
		},
	})
}

