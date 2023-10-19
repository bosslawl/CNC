package subcommands

import (
	"strconv"
	"strings"
	"fmt"
	"time"

	Sessions "Rain/core/masters/sessions"
)

func LookupSession(session *Sessions.Session, cmd []string) error {

	if len(cmd) <= 2 {
		session.Channel.Write([]byte("Invalid Syntax\r\n"))
		session.Channel.Write([]byte("Syntax: sessions lookup <Username>>@<<ID>>\r\n"))
		session.Channel.Write([]byte("Example: sessions lookup root@1\r\n"))
		return nil
	}

	for U := 2; U < len(cmd); U++ {
		CommandDebug := strings.Split(cmd[U], "@")
		if len(CommandDebug) <= 1 {
			session.Channel.Write([]byte("Failed to lookup session \"" + CommandDebug[1] + "\" due to invaild string split\r\n"))
			continue
		}

		Turn, error := strconv.Atoi(CommandDebug[1])
		if error != nil {
			session.Channel.Write([]byte("Failed to lookup session \"" + CommandDebug[1] + "\" due to invaild string split\r\n"))
			continue
		}

		SessionID := Sessions.Sessions[int64(Turn)]
		if SessionID == nil {
			session.Channel.Write([]byte("Failed to find session for \"" + CommandDebug[0] + "\" with id of \"" + CommandDebug[1] + "\"\r\n"))
			continue
		}

		session.Channel.Write([]byte("User > " + SessionID.User.Username + " | IPv4 > " + session.Conn.RemoteAddr().String() + " | Since > " + fmt.Sprintf("%.2f mins", time.Since(SessionID.Created).Minutes()) + " | Client > " + ClientVersion(string(session.Conn.ClientVersion())) + "\r\n"))

	}

	return nil

}

func ClientVersion(Version string) string {
	var KeyWords = []string{"windows", "Ubuntu", "KiTTY", "PuTTY"}

	for _, Keyword := range KeyWords {
		if strings.Contains(Version, Keyword) {
			return Keyword
		}
	}

	return "Unknown"
}