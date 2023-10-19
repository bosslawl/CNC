package Views

import (
	Execute "Rain/core/config/views/user"
	Database "Rain/core/database"
	"fmt"
	"time"

	"github.com/xlzd/gotp"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/crypto/ssh"
)

func MFANeeded(channel ssh.Channel, conn *ssh.ServerConn, User *Database.User_Struct) error {
	_, error := Execute.Execute_Standard("mfa-needed", User, channel, true, false)
	if error != nil {
		fmt.Println(error)
		return error
	}

	Term := terminal.NewTerminal(channel, "MFA Code> ")

	Code, error := Term.ReadLine()
	if error != nil {
		fmt.Println(error)
		channel.Close()
		return error
	}

	TOTP := gotp.NewDefaultTOTP(User.MFAToken)
	if TOTP.Now() != Code {
		New := map[string]string {
			"reason": "Code given is invalid",
		}
		Execute.Execute_CustomTerm("mfa-invalid", User, channel, true, New)
		time.Sleep(4 * time.Second)
		channel.Close()
		return nil
	}
	return nil
}
