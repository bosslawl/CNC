package Views

import (
	Database "Rain/core/database"
	"context"
	"reflect"
	"strings"

	Term "Rain/core/config/views/term"
	Execute "Rain/core/config/views/user"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
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

func NewLogin(channel ssh.Channel, conn *ssh.ServerConn, User *Database.User_Struct) (*Database.User_Struct, error) {

	_, error := Execute.ExecuteNewUser("new-login", User, channel, true, false)
	if error != nil {
		return nil, error
	}

	channel.Write([]byte("\033[" + Term.Pos2 + ";" + Term.Pos1 + "f"))
	Usern := term.NewTerminal(channel, "\x1b[30m\x1b[47m")

	Username, error := Usern.ReadLine()
	if error != nil {
		return nil, error
	}

	channel.Write([]byte("\033[" + Term.Pos4 + ";" + Term.Pos3 + "f"))
	Pwd := term.NewTerminal(channel, "\x1b[30m\x1b[47m")

	Password, error := Pwd.ReadLine()
	if error != nil {
		return nil, error
	}

	channel.Write([]byte("\x1b[0m"))

	collection := Database.MCXT.Database("Rain").Collection("Users")
	var result User_Struct
	var username string
	username = strings.ToLower(Username)
	password := []byte(string(Password))
	err := collection.FindOne(context.TODO(), bson.M{"Username": username}).Decode(&result)
	if err != nil {
		Execute.ExecuteNewUser("login-incorrect", User, channel, true, false)
		time.Sleep(2 * time.Second)
		channel.Close()
		return nil, err
	} else {
		err = bcrypt.CompareHashAndPassword([]byte(result.Password), password)
		if err != nil {
			Execute.ExecuteNewUser("login-incorrect", User, channel, true, false)
			time.Sleep(2 * time.Second)
			channel.Close()
			return nil, err
		}
		User = &Database.User_Struct{
			Username:        result.Username,
			Password:        result.Password,
			Admin:           result.Admin,
			Powersaving:     result.Powersaving,
			Bypassblacklist: result.Bypassblacklist,
			Reseller:        result.Reseller,
			Newuser:         result.Newuser,
			Banned:          result.Banned,
			VIP:             result.VIP,
			Raw:             result.Raw,
			Holder:          result.Holder,
			MaxSessions:     result.MaxSessions,
			MaxTime:         result.MaxTime,
			Cooldown:        result.Cooldown,
			Concurrents:     result.Concurrents,
			MFA:             result.MFA,
			MFAToken:        result.MFAToken,
			Expiry:          result.Expiry,
		}
		return User, nil
	}

}
