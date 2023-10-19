package Commands

import (
	"strconv"

	"Rain/core/functions/external"
	"Rain/core/functions/util/attackGroups"
	"Rain/core/config/views/branding"
	"Rain/core/masters/sessions"
	"Rain/core/functions/json"
)

func init() {
	Load_Commands(&Command_interface{
		Name:        "reload",
		Roles:       []string{"admin"},
		Admin:       true,
		Reseller:    false,
		Description: "Reload's all assets and ofsets",
		CommandExecution: func(session *Sessions.Session, cmd []string) error {

			load, error := Branding.Load_Items()
			if error != nil {
				session.Channel.Write([]byte("\x1b[38;5;15m[\x1b[38;5;1mFATAL\x1b[38;5;15m] CNC couldn't find your branding file\r\n"))
				return nil
			} else {
				session.Channel.Write([]byte("\x1b[38;5;15m[\x1b[38;5;2mOK\x1b[38;5;15m] CNC correctly reloaded " + strconv.Itoa(load) + " items of branding correctly\r\n"))
			}

			error = ParseJson.Configuration_Parse()
			if error != nil {
				session.Channel.Write([]byte("\x1b[38;5;15m[\x1b[38;5;1mFATAL\x1b[38;5;15m] CNC couldn't find your `config.json` file\r\n"))
				return nil
			} else {
				session.Channel.Write([]byte("\x1b[38;5;15m[\x1b[38;5;2mOK\x1b[38;5;15m] CNC correctly reloaded your `config.json`\r\n"))
			}

			error = ParseJson.Attacks_Parse()
			if error != nil {
				session.Channel.Write([]byte("\x1b[38;5;15m[\x1b[38;5;1mFATAL\x1b[38;5;15m] CNC couldn't find your `api-attack.json` file\r\n"))
				return nil
			} else {
				session.Channel.Write([]byte("\x1b[38;5;15m[\x1b[38;5;2mOK\x1b[38;5;15m] CNC correctly reloaded your `api-attack.json`\r\n"))
			}

			error = Attack_Groups.SortGroups()
			if error != nil {
				session.Channel.Write([]byte("\x1b[38;5;15m[\x1b[38;5;1mFATAL\x1b[38;5;15m] CNC couldn't sort your attack groups\r\n"))
				return nil
			} else {
				session.Channel.Write([]byte("\x1b[38;5;15m[\x1b[38;5;2mOK\x1b[38;5;15m] CNC correctly sorted your attack groups\r\n"))
			}

			External.GatherExCommands()

			return nil
		},
	})
}
