package Commands

import (
	Execute "Rain/core/config/views/user"
	Database "Rain/core/database"
	ParseConfig "Rain/core/functions/json"
	MFA "Rain/core/functions/mfa"
	Sessions "Rain/core/masters/sessions"
	"fmt"
	"net/url"
	"strings"

	"github.com/xlzd/gotp"
	"golang.org/x/crypto/ssh/terminal"
)

func init() {
	Load_Commands(&Command_interface{
		Name:        "mfa",
		Roles:       []string{"everyone"},
		Admin:       false,
		Reseller:    false,
		Description: "Enable MFA for your account",
		CommandExecution: func(session *Sessions.Session, cmd []string) error {
			channel := session.Channel
			if len(cmd) < 2 {
				return nil
			}

			switch cmd[1] {

			case "on":
				fmt.Fprintln(channel, "PRESS ENTER")

				Byteess := make([]byte, 4)
				_, error := channel.Read(Byteess)
				channel.Write([]byte("[8;55;100t"))
				channel.Write([]byte("\033[2J\033[;H"))

				secret := MFA.GenTOTPSecret()

				totp := gotp.NewDefaultTOTP(secret)

				qrs := MFA.New()

				qrcode := qrs.Get("otpauth://totp/" + url.QueryEscape(ParseConfig.ConfigParse.App.AppName) + ":" + url.QueryEscape(session.User.Username) + "?secret=" + secret + "&issuer=" + url.QueryEscape(ParseConfig.ConfigParse.App.AppName) + "&digits=6&period=30").Sprint()
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
					New := map[string]string {
						"reason": "Code given is invalid",
					}
					Execute.Execute_CustomTerm("mfa-invalid", session.User, session.Channel, true, New)
					return nil
				}

				errorss := Database.EditFeild(session.User.Username, "MFA", "1", false)
				if !errorss {
					fmt.Println((errorss))
				}

				errorsss := Database.EditFeild(session.User.Username, "MFAToken", secret, false)
				if !errorsss {
					Execute.Execute_Standard("database-error", session.User, session.Channel, true, false)
					return nil
				} else {
					Execute.Execute_Standard("mfa-enabled", session.User, session.Channel, true, false)
					channel.Write([]byte("[8;24;80t"))
					return nil
				}

			case "off":

				User, _ := Database.GetUser(session.User.Username)
				Term := terminal.NewTerminal(channel, "MFA Code> ")

				Code, error := Term.ReadLine()
				if error != nil {
					channel.Close()
					return error
				}

				TOTP := gotp.NewDefaultTOTP(User.MFAToken)
				if TOTP.Now() != Code {
					New := map[string]string {
						"reason": "Code given is invalid",
					}
					Execute.Execute_CustomTerm("mfa-invalid", session.User, session.Channel, true, New)
					return nil
				} else {
					error := Database.EditFeild(session.User.Username, "MFA", "0", false)
					if !error {
						Execute.Execute_Standard("database-error", session.User, session.Channel, true, false)
						return nil
					}

					Execute.Execute_Standard("mfa-disabled", session.User, session.Channel, true, false)
					return nil
				}
			}
			return nil
		},
	})
}
