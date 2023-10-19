package Sessions

import (
	"errors"

	"golang.org/x/crypto/ssh"
)

func Broadcast(payload []byte) (int, error) {
	for _, Session := range Sessions {
		Session.Channel.Write([]byte(payload))
	}

	return len(payload), nil
}

func Message(payload []byte, identification int64, user string) (int, error) {
	session := Sessions[identification]
	if session == nil || session.User.Username != user {
		return 0, errors.New("unknown session")
	}

	return session.Channel.Write(payload)
}

func Disconnect(identification int64, user string) error {
	session := Sessions[identification]
	if session == nil || session.User.Username != user {
		return errors.New("unknown session")
	}

	return session.Channel.Close()
}

func AwaitClose(session *Session) {
	error := session.Conn.Wait()
	if error != nil {
		delete(Sessions, session.Key)
		return
	}
	delete(Sessions, session.Key)
	return
}

//checks how many sessions a user has open (session locker)
func Open(conn *ssh.ServerConn) int {
	Ammount := 0
	// loops throughout the session list checking the ammount of sessions open by that person
	for _, s := range Sessions {
		if s.User.Username == conn.User() {
			Ammount++
		}
	}
	return Ammount
}