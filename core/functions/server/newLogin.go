package Server

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	CNC "Rain/core/config/admin"
	External "Rain/core/functions/external"
	ParseJson "Rain/core/functions/json"

	"github.com/fatih/color"
	"golang.org/x/crypto/ssh"
)

var New_Server *ssh.ServerConfig

type Fields struct {
	Username string
	Password string
}

var _ = reflect.TypeOf(Fields{})

func New() {
	Config := &ssh.ServerConfig{

		//PasswordCallback: func(c ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error) {
		//	collection := Database.MCXT.Database("Rain").Collection("Users")
		//	var result Fields
		//	var username string
		//	username = strings.ToLower(c.User())
		//	password := []byte(string(pass))
		//	err := collection.FindOne(context.TODO(), bson.M{"Username": username}).Decode(&result)
		//	if err != nil {
		//		fmt.Println(err)
		//		return nil, errors.New("Failed Username Doesn't Exist")
		//	} else {
		//		err = bcrypt.CompareHashAndPassword([]byte(result.Password), password)
		//		if err != nil {
		//			fmt.Println(err)
		//			return nil, errors.New("Failed login attempt")
		//		}
		//		return nil, nil
		//	}
		//},
		NoClientAuth: true,
	}

	Config.MaxAuthTries = ParseJson.ConfigParse.Masters.MastersMaxAuth

	PrivateKeyFile, error := ioutil.ReadFile(CNC.BuildFolder + "/ssh-key.ppk")
	if error != nil || PrivateKeyFile == nil {
		fmt.Println(color.WhiteString(" - ") + color.WhiteString("[") + color.BlueString("STATUS") + color.WhiteString("]") + color.WhiteString(" -> ") + color.WhiteString("[") + color.YellowString("SSH") + color.WhiteString(":") + color.RedString("FATAL") + color.WhiteString("]") + color.WhiteString(" -> ") + color.WhiteString("[") + color.RedString("Failed to load SSH server key correctly") + color.WhiteString("]"))
		os.Exit(1)
	}

	ParsedKeyByte, error := ssh.ParsePrivateKey(PrivateKeyFile)
	if error != nil || ParsedKeyByte == nil {
		fmt.Println(color.WhiteString(" - ") + color.WhiteString("[") + color.BlueString("STATUS") + color.WhiteString("]") + color.WhiteString(" -> ") + color.WhiteString("[") + color.YellowString("SSH") + color.WhiteString(":") + color.RedString("FATAL") + color.WhiteString("]") + color.WhiteString(" -> ") + color.WhiteString("[") + color.RedString("Failed to load SSH server key correctly") + color.WhiteString("]"))
		os.Exit(1)
	}

	Config.AddHostKey(ParsedKeyByte)

	External.GatherExCommands()

	New_Server = Config

	Serve()
}
