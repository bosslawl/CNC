package subcommands

import (
	"fmt"
	"strconv"
	"strings"

	Sessions "Rain/core/masters/sessions"

	"github.com/alexeyco/simpletable"
)

func SessionsIDs(session *Sessions.Session) error {

	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "ID"},
			{Align: simpletable.AlignCenter, Text: "Username"},
		},
	}

	for _, I := range Sessions.Sessions {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: strconv.Itoa(int(I.Key))},
			{Align: simpletable.AlignCenter, Text: I.User.Username},
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.SetStyle(simpletable.StyleUnicode)

	fmt.Fprint(session.Channel, "")
	fmt.Fprintln(session.Channel, strings.ReplaceAll(table.String(), "\n", "\r\n"))
	fmt.Fprint(session.Channel, "\r")

	return nil
}