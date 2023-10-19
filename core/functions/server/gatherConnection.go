package Server

import (
	"fmt"
	"log"
	"net"
	"os"

	ParseJson "Rain/core/functions/json"
	"Rain/core/masters"

	"github.com/fatih/color"
	"golang.org/x/crypto/ssh"
)

func Serve() {

	Listener, error := net.Listen("tcp", ":"+ParseJson.ConfigParse.Masters.MastersPort)
	if error != nil {
		fmt.Println(color.WhiteString(" - ") + color.WhiteString("[") + color.BlueString("STATUS") + color.WhiteString("]") + color.WhiteString(" -> ") + color.WhiteString("[") + color.YellowString("SSH") + color.WhiteString(":") + color.RedString("FATAL") + color.WhiteString("]") + color.WhiteString(" -> ") + color.WhiteString("[") + color.RedString(error.Error()) + color.WhiteString("]"))
		os.Exit(1)
	} else {
		fmt.Println(color.WhiteString(" - ") + color.WhiteString("[") + color.BlueString("STATUS") + color.WhiteString("]") + color.WhiteString(" -> ") + color.WhiteString("[") + color.YellowString("SSH") + color.WhiteString(":") + color.GreenString("SUCCESS") + color.WhiteString("]"))
		fmt.Println(color.WhiteString(" - ") + color.WhiteString("[") + color.BlueString("SSH") + color.WhiteString("]") + color.WhiteString(" Port Opened & Listening") + color.WhiteString(": ") + color.MagentaString(ParseJson.ConfigParse.Masters.MastersPort+"\n"))
	}

	go Users.TitleWorker()

	for {
		conn, error := Listener.Accept()
		if error != nil || conn == nil {
			continue
		}

		go func() {
			Connection, chans, reqs, err := ssh.NewServerConn(conn, New_Server)
			if err != nil {
				log.Printf(color.WhiteString(" - ") + color.WhiteString("[") + color.BlueString("STATUS") + color.WhiteString("]") + color.WhiteString(" -> ") + color.WhiteString("[") + color.YellowString("SSH") + color.WhiteString(":") + color.RedString("FATAL") + color.WhiteString("]") + color.WhiteString(" -> ") + color.WhiteString("[") + color.RedString("Handshake failed with connection from %s", conn.RemoteAddr().String()) + color.WhiteString("]"))
				return
			}

			go ssh.DiscardRequests(reqs)

			go Users.Channels(chans, Connection)
		}()

	}
}
