package Users

import (
	ParseJson "Rain/core/functions/json"

	"golang.org/x/crypto/ssh"
)

func Channels(chans <-chan ssh.NewChannel, Connection *ssh.ServerConn) {

	for newChan := range chans {
		if newChan.ChannelType() != "session" {
			newChan.Reject(ssh.UnknownChannelType, "UnknownChannelType")
			return
		}
		channel, requests, err := newChan.Accept()
		if err != nil {
			return
		}

		go func(in <-chan *ssh.Request) {
			for req := range in {
				switch req.Type {
				case "pty-req":
					req.Reply(true, nil)
					continue
				case "shell":
					req.Reply(true, nil)
					continue
				}
				req.Reply(false, nil)
			}
		}(requests)

		channel.Write([]byte("\033]0;" + ParseJson.ConfigParse.App.AppName + "\007"))
		HandleFunc(channel, Connection)
	}
}
