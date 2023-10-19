package Execute

import (
	"bufio"
	"strings"

	viewsMapped "Rain/core/config/views"
	Term "Rain/core/config/views/term"
	Termfx "Rain/core/config/views/tfx"
	Database "Rain/core/database"

	"golang.org/x/crypto/ssh"
)

func Execute_Standard(Name string, User *Database.User_Struct, Channel ssh.Channel, Colourize bool, Title bool) (string, error) {

	User, _ = Database.GetUser(User.Username)

	Cli := Termfx.New()

	Cli = Term.Standard_Struct(Cli, User, Colourize, nil)

	View := viewsMapped.Branding[Name]

	ReadScan := bufio.NewScanner(strings.NewReader(View))

	var title string
	var Line int
	for ReadScan.Scan() {
		CliBrand, error := Cli.ExecuteString(ReadScan.Text())
		if error != nil {
			continue
		}

		CliBrand = strings.Replace(CliBrand, "\\x1b", "\x1b", -1)
		CliBrand = strings.Replace(CliBrand, "\\033", "\x1b", -1)

		Line++

		if Title {
			if Line == 1 {
				title = CliBrand
				continue
			}

		} else {
			Channel.Write([]byte(CliBrand + "\r\n"))
		}
		continue
	}

	return title, nil
}

func ExecuteNewUser(Name string, User *Database.User_Struct, Channel ssh.Channel, Colourize bool, Title bool) (string, error) {

	UserS := &Database.User_Struct{
		Username: "none",
		Password: "none",

		Admin:           false,
		Powersaving:     false,
		Bypassblacklist: false,
		Reseller:        false,
		Newuser:         false,
		Banned:          false,
		VIP:             false,
		Raw:             false,
		Holder:          false,
		MFA:             false,
		MFAToken:        "",

		MaxTime:     0,
		Cooldown:    0,
		Concurrents: 0,

		Expiry: 0,
	}

	Cli := Termfx.New()
	Cli = Term.Standard_Struct(Cli, UserS, Colourize, nil)

	View := viewsMapped.Branding[Name]

	ReadScan := bufio.NewScanner(strings.NewReader(View))

	var title string
	var Line int
	for ReadScan.Scan() {
		CliBrand, error := Cli.ExecuteString(ReadScan.Text())
		if error != nil {
			continue
		}

		CliBrand = strings.Replace(CliBrand, "\\x1b", "\x1b", -1)
		CliBrand = strings.Replace(CliBrand, "\\033", "\x1b", -1)

		Line++

		if Title {
			if Line == 1 {
				title = CliBrand
				continue
			}

		} else {
			Channel.Write([]byte(CliBrand + "\r\n"))
		}
		continue
	}

	return title, nil
}

func ExecuteString(String string, User *Database.User_Struct, Channel ssh.Channel, Colourize bool, Title bool) (string, error) {

	User, _ = Database.GetUser(User.Username)

	Cli := Termfx.New()

	Cli = Term.Standard_Struct(Cli, User, Colourize, nil)

	ReadScan := bufio.NewScanner(strings.NewReader(String))

	var title string
	var Line int
	for ReadScan.Scan() {
		CliBrand, error := Cli.ExecuteString(ReadScan.Text())
		if error != nil {
			continue
		}

		CliBrand = strings.Replace(CliBrand, "\\x1b", "\x1b", -1)
		CliBrand = strings.Replace(CliBrand, "\\033", "\x1b", -1)

		Line++

		if Title {
			if Line == 1 {
				title = CliBrand
				continue
			}

		} else {
			Channel.Write([]byte(CliBrand + "\r\n"))
		}
		continue
	}

	return title, nil
}
