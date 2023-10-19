package Build

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"reflect"
	"time"

	CNC "Rain/core/config/admin"
	Database "Rain/core/database"
	Util "Rain/core/functions/util"

	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/bson"
)

type details struct {
	Username string
	Password string
}

var _ = reflect.TypeOf(details{})

func InsertTables() error {
	var Password string

	Password = MakeRandom(CNC.NewMongo_Passwordlen)
	Pwd := Util.PasswordHash(Password)

	data := details{
		Username: CNC.NewMongo_Username,
		Password: Password,
	}

	File, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile("./build/login-info.json", File, 0644)

	var NewUser = Database.User_Struct{
		Username:        CNC.NewMongo_Username,
		Password:        Pwd,
		Admin:           true,
		Powersaving:     false,
		Bypassblacklist: false,
		Reseller:        true,
		Newuser:         true,
		Banned:          false,
		VIP:             true,
		Raw:             true,
		Holder:          true,
		MaxSessions:     32,
		MaxTime:         7200,
		MFA:             false,
		MFAToken:        "",
		Cooldown:        0,
		Concurrents:     15,
		Expiry:          time.Now().Add((time.Hour * 24) * time.Duration(365)).Unix(),
	}

	db := Database.MCXT.Database("Rain")

	command := bson.D{{"create", "Attacks"}}
	command2 := bson.D{{"create", "Users"}}

	var result bson.M
	if err := db.RunCommand(context.TODO(), command).Decode(&result); err != nil {
		fmt.Println("Attacks already exist")
	}

	var result2 bson.M
	if err2 := db.RunCommand(context.TODO(), command2).Decode(&result2); err2 != nil {
		fmt.Println("Users already exist")
	}

	error := Database.CreateNewUser(&NewUser)
	if error == nil {
		fmt.Println(color.WhiteString(" - ") + color.WhiteString("[") + color.BlueString("DATABASE") + color.WhiteString("]") + color.WhiteString(" Username") + color.WhiteString(": ") + color.MagentaString(CNC.NewMongo_Username))
		fmt.Println(color.WhiteString(" - ") + color.WhiteString("[") + color.BlueString("DATABASE") + color.WhiteString("]") + color.WhiteString(" Password") + color.WhiteString(": ") + color.MagentaString(Password))
		return nil
	}

	return error
}

func MakeRandom(n int) string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
