package Views

import (
	"time"

	Database "Rain/core/database"
	Execute "Rain/core/config/views/user"

	"golang.org/x/crypto/ssh"
)

func PlanFinished(channel ssh.Channel, conn *ssh.ServerConn, User *Database.User_Struct) error {
	_, error := Execute.Execute_Standard("login-expired", User, channel, true, false)
	if error != nil {
		return error
	}

	time.Sleep(10 * time.Second)
	channel.Close()
	return nil
}
