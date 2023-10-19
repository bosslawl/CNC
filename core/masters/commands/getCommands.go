package Commands

import (
	"fmt"
	"strings"

	Util "Rain/core/functions/util"
	Handler "Rain/core/masters/handler"
	Sessions "Rain/core/masters/sessions"
	ParseINI "Rain/core/functions/ini"

	"github.com/alexeyco/simpletable"
)

func init() {
	Load_Commands(&Command_interface{
		Name:        "commands",
		Roles:       []string{"everyone"},
		Admin:       false,
		Reseller:    false,
		Description: "Display all commands",
		CommandExecution: func(Session *Sessions.Session, cmd []string) error {

			Session.Channel.Write([]byte("\x1b[0m"))
			table := simpletable.New()
			table.Header = &simpletable.Header{
				Cells: []*simpletable.Cell{
					{Align: simpletable.AlignCenter, Text: "Name"},
					{Align: simpletable.AlignCenter, Text: "Description"},
					{Align: simpletable.AlignCenter, Text: "Admin"},
					{Align: simpletable.AlignCenter, Text: "Reseller"},
				},
			}

			for _, C := range Command {

				r := []*simpletable.Cell{
					{Align: simpletable.AlignCenter, Text: C.Name},
					{Align: simpletable.AlignCenter, Text: C.Description},
					{Align: simpletable.AlignCenter, Text: Util.ColourizeBoolen(C.Admin, true)},
					{Align: simpletable.AlignCenter, Text: Util.ColourizeBoolen(C.Reseller, true)},
				}

				table.Body.Cells = append(table.Body.Cells, r)
			}

			for _, C := range Handler.BetaMapHandler {

				r := []*simpletable.Cell{
					{Align: simpletable.AlignCenter, Text: C.CommandName},
					{Align: simpletable.AlignCenter, Text: C.CommandDescription},
					{Align: simpletable.AlignCenter, Text: Util.ColourizeBoolen(C.CommandAdmin, true)},
					{Align: simpletable.AlignCenter, Text: Util.ColourizeBoolen(C.CommandReseller, true)},
				}

				table.Body.Cells = append(table.Body.Cells, r)
			}

			ParseINI.ParseTableCommands(table)

			fmt.Fprint(Session.Channel, "")
			fmt.Fprintln(Session.Channel, strings.ReplaceAll(table.String(), "\n", "\r\n"))
			fmt.Fprint(Session.Channel, "\r")
			return nil
		},
	})
}
