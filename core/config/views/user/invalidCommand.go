package Execute

import (
	"bufio"
	"strings"

	Database "Rain/core/database"
	viewsMapped "Rain/core/config/views"
	Termfx "Rain/core/config/views/tfx"
	"Rain/core/config/views/term"

	"golang.org/x/crypto/ssh"
)

func Execute_CustomTerm(Name string, User *Database.User_Struct, Channel ssh.Channel, Colourize bool, mapped map[string]string) (string, error) {

	User, _ = Database.GetUser(User.Username)

	Cli := Termfx.New()

	Cli = Term.Standard_Struct(Cli, User, Colourize, mapped)

	View := viewsMapped.Branding[Name]

	ReadScan := bufio.NewScanner(strings.NewReader(View))

	var title string
	for ReadScan.Scan() {
		CliBrand, error := Cli.ExecuteString(ReadScan.Text())
		if error != nil {
			continue
		}

		CliBrand = strings.Replace(CliBrand, "\x1b", "\x1b", -1)
		CliBrand = strings.Replace(CliBrand, "\033", "\x1b", -1)

		Channel.Write([]byte(CliBrand + "\r\n"))

	}

	return title, nil
}
