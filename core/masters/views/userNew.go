package Views

import (
	"errors"

	Term "Rain/core/config/views/term"
	Execute "Rain/core/config/views/user"
	Database "Rain/core/database"
	Util "Rain/core/functions/util"

	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

func NewAccountLogin(channel ssh.Channel, conn *ssh.ServerConn, User *Database.User_Struct) error {

	_, error := Execute.Execute_Standard("login-new-user", User, channel, true, false)
	if error != nil {
		return error
	}

	channel.Write([]byte("\033[" + Term.Pos2 + ";" + Term.Pos1 + "f"))
	NewPwd := term.NewTerminal(channel, "\x1b[30m\x1b[47m")

	NewPassword, error := NewPwd.ReadLine()
	if error != nil {
		return nil
	}

	channel.Write([]byte("\033[" + Term.Pos4 + ";" + Term.Pos3 + "f"))
	NewCPwd := term.NewTerminal(channel, "\x1b[30m\x1b[47m")

	NewConfirmPassword, error := NewCPwd.ReadLine()
	if error != nil {
		return nil
	}

	channel.Write([]byte("\x1b[0m"))

	if NewPassword != NewConfirmPassword {
		Execute.Execute_Standard("password-match", User, channel, true, false)
		return errors.New("passwords dont match")
	}

	if len(NewConfirmPassword) < 4 {
		Execute.Execute_Standard("password-length", User, channel, true, false)
		return errors.New("invaild password length")
	}

	pwd := Util.PasswordHash(NewConfirmPassword)
	Raw := Database.EditFeild(User.Username, "Password", pwd, false)
	if !Raw {
		Execute.Execute_Standard("password-failed", User, channel, true, false)
		return errors.New("failed to update password correctly")
	}

	Raw = Database.EditFeild(User.Username, "Newuser", "0", false)
	if !Raw {
		Execute.Execute_Standard("password-failed", User, channel, true, false)
		return errors.New("failed to update password correctly")
	}

	Execute.Execute_Standard("password-success", User, channel, true, false)

	return nil
}
