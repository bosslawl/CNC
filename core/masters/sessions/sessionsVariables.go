package Sessions

import (
	"reflect"
	"sync"
	"time"

	Database "Rain/core/database"

	"golang.org/x/crypto/ssh"
)

var (
	Sessions = make(map[int64]*Session)
	NycMux   sync.Mutex
)

type Session struct {
	Key       int64
	User      *Database.User_Struct
	Channel   ssh.Channel
	Conn      *ssh.ServerConn
	Chat      bool
	Created   time.Time
	ColourOne string
	ColourTwo string
}

var _ = reflect.TypeOf(Session{})

func (s *Session) Remove() {
	NycMux.Lock()
	delete(Sessions, s.Key)
	NycMux.Unlock()
}

func (s *Session) Check(conn *ssh.ServerConn) {
	error := conn.Wait()
	if error != nil {
		s.Remove()
	}
	return
}

func (s *Session) Open(username string) int {
	Ammount := 0
	for _, s := range Sessions {
		if s.User.Username == username {
			Ammount++
		}
	}
	return Ammount
}

func Online() int {
	return len(Sessions)
}

// direct message all sessions open by a user
func DirectMessage(user string, payload string) error {
	for _, s := range Sessions {
		if s.User.Username == user {
			_, err := s.Channel.Write([]byte(payload))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
