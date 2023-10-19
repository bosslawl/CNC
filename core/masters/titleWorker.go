package Users

import (
	"Rain/core/masters/sessions"
	Execute "Rain/core/config/views/user"
	"time"
)

func TitleWorker() {

	for {

		for _, Session := range Sessions.Sessions {
			TITLE, _ := Execute.Execute_Standard("title", Session.User, Session.Channel, false, true)

			Session.Channel.Write([]byte("\033]0;" + TITLE + "\007"))
		}

		time.Sleep(1 * time.Second)
	}
}
