package Views

import (
	Execute "Rain/core/config/views/user"
	Database "Rain/core/database"
	ParseJson "Rain/core/functions/json"
	MFA "Rain/core/functions/mfa"
	"fmt"
	"net/url"
	"strings"

	"github.com/xlzd/gotp"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

func MFARequired(channel ssh.Channel, conn *ssh.ServerConn, User *Database.User_Struct) error {

	fmt.Fprintln(channel, "PRESS ENTER")

	Byteess := make([]byte, 4)
	_, error := channel.Read(Byteess)
	channel.Write([]byte("[8;55;100t"))
	channel.Write([]byte("\033[2J\033[;H"))

	secret := MFA.GenTOTPSecret()

	totp := gotp.NewDefaultTOTP(secret)

	qrs := MFA.New()

	qrcode := qrs.Get("otpauth://totp/" + url.QueryEscape(ParseJson.ConfigParse.App.AppName) + ":" + url.QueryEscape(User.Username) + "?secret=" + secret + "&issuer=" + url.QueryEscape(ParseJson.ConfigParse.App.AppName) + "&digits=6&period=30").Sprint()
	fmt.Fprintln(channel, strings.ReplaceAll(qrcode, "\n", "\r\n"))
	fmt.Fprint(channel, "You may scan this code to register your account info a 2FA App, Google Auth, Twilio Authy\r\n")
	fmt.Fprint(channel, "or enter this code> "+secret+"\r\n")
	term := terminal.NewTerminal(channel, "Code>")

	Code, error := term.ReadLine()
	if error != nil {
		channel.Write([]byte("\r\n"))
		return nil
	}

	if totp.Now() != Code {
		New := map[string]string{
			"reason": "Code given is invalid",
		}
		Execute.Execute_CustomTerm("mfa-invalid", User, channel, true, New)
		return nil
	}

	errorss := Database.EditFeild(User.Username, "MFA", "1", false)
	if !errorss {
		fmt.Println((errorss))
	}

	errorsss := Database.EditFeild(User.Username, "MFAToken", secret, false)
	if !errorsss {
		Execute.Execute_Standard("database-error", User, channel, true, false)
		return nil
	} else {
		Execute.Execute_Standard("mfa-enabled", User, channel, true, false)
		channel.Write([]byte("[8;24;80t"))
		return nil
	}
}
