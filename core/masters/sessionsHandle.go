package Users

import (
	"fmt"
	"reflect"
	"time"

	Execute "Rain/core/config/views/user"
	Database "Rain/core/database"
	ParseJson "Rain/core/functions/json"
	Replireadings "Rain/core/functions/rep"
	Sessions "Rain/core/masters/sessions"
	Views "Rain/core/masters/views"

	"golang.org/x/crypto/ssh"
)

type User_Struct struct {
	Username string
	Password string

	Admin           bool
	Powersaving     bool
	Bypassblacklist bool
	Reseller        bool
	Newuser         bool
	Banned          bool
	VIP             bool
	Raw             bool
	Holder          bool
	MaxSessions     int

	MaxTime     int
	Cooldown    int
	Concurrents int
	MFA         bool
	MFAToken    string

	Expiry int64
}

var _ = reflect.TypeOf(User_Struct{})

func HandleFunc(channel ssh.Channel, conn *ssh.ServerConn) {

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
		MaxSessions:     1,
		MFAToken:        "",

		MaxTime:     0,
		Cooldown:    0,
		Concurrents: 0,

		Expiry: 0,
	}

	User, error := Views.NewLogin(channel, conn, UserS)
	if error != nil {
		channel.Write([]byte(ParseJson.ConfigParse.App.AppName + " is currently experiencing issues, please come back later!"))
		time.Sleep(5 * time.Second)
		channel.Close()
		return
	}

	if User.Banned {
		Row := Views.Banned(channel, conn, User)
		if Row != nil {
			channel.Close()
			return
		}
	}

	if User.Newuser {
		error := Views.NewAccountLogin(channel, conn, User)
		if error != nil {
			time.Sleep(10 * time.Second)
			channel.Close()
			return
		}
	}

	if User.Expiry <= time.Now().Unix() {
		fmt.Println(User.Expiry, time.Now().Unix())
		error := Views.PlanFinished(channel, conn, User)
		if error != nil {
			channel.Close()
			return
		}
		return
	}

	if ParseJson.ConfigParse.Controls.Catpcha.Status {
		if ParseJson.ConfigParse.Controls.Catpcha.AdminBypass && User.Admin {
		} else {
			Views.Catpcha(channel, conn, User)
			if error != nil {
				channel.Close()
				return
			}
		}
	}

	var session = &Sessions.Session{
		Key:       time.Now().Unix(),
		User:      User,
		Channel:   channel,
		Conn:      conn,
		Chat:      false,
		Created:   time.Now(),
		ColourOne: "",
		ColourTwo: "",
	}

	Sessions.NycMux.Lock()
	Sessions.Sessions[session.Key] = session
	Sessions.NycMux.Unlock()

	go session.Check(conn)
	open := session.Open(User.Username)

	if open > User.MaxSessions {
		Execute.Execute_Standard("max-sessions", User, channel, true, false)
		time.Sleep(5 * time.Second)
		channel.Close()
		return
	}

	if User.MFA {
		error := Views.MFANeeded(channel, conn, User)
		if error != nil {
			channel.Close()
			return
		}
	}

	if !User.MFA && ParseJson.ConfigParse.Controls.MFA.Status {
		if ParseJson.ConfigParse.Controls.MFA.AdminBypass && User.Admin {
		} else {
			error := Views.MFARequired(channel, conn, User)
			if error != nil {
				channel.Close()
				return
			}
		}
	}

	_, Pre1, Pre2 := Replireadings.GetFunctions("prompt")

	var New = Sessions.Session{

		Key: time.Now().Unix(),

		User: User,

		Channel: channel,
		Conn:    conn,

		Chat: false,

		Created: time.Now(),

		ColourOne: Pre1,
		ColourTwo: Pre2,
	}

	Sessions.NycMux.Lock()
	Sessions.Sessions[New.Key] = &New
	Sessions.NycMux.Unlock()

	go Sessions.AwaitClose(&New)

	Views.Prompt(&New)
}
