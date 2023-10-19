package Views

import (
	ParseJson "Rain/core/functions/json"
	Execute "Rain/core/config/views/user"
	Database "Rain/core/database"
	"math/rand"
	"strconv"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

// catpcha screen for clients
func Catpcha(channel ssh.Channel, conn *ssh.ServerConn, User *Database.User_Struct) error {
	var Used int

	fromthetop:

	if ParseJson.ConfigParse.Controls.Catpcha.AllowedAttempts == Used {
		Execute.Execute_Standard("captcha-failed", User, channel, true, false)
		time.Sleep(10 * time.Second)
		channel.Close()
		return nil
	}

	NumOne := rand.Intn(ParseJson.ConfigParse.Controls.Catpcha.Question.MaxGen - ParseJson.ConfigParse.Controls.Catpcha.Question.MinGen) + ParseJson.ConfigParse.Controls.Catpcha.Question.MinGen
	NumTwo := rand.Intn(ParseJson.ConfigParse.Controls.Catpcha.Question.MaxGen - ParseJson.ConfigParse.Controls.Catpcha.Question.MinGen) + ParseJson.ConfigParse.Controls.Catpcha.Question.MinGen
	
	
	_, err := Execute.Execute_Standard("captcha-banner", User, channel, true, false)
	if err != nil {
		channel.Write([]byte("Answer This Question to gain access -> "+strconv.Itoa(NumOne) + " + " + strconv.Itoa(NumTwo)))
	}
	channel.Write([]byte("Answer This Question to gain access -> "+strconv.Itoa(NumOne) + " + " + strconv.Itoa(NumTwo)))

	Answer := terminal.NewTerminal(channel, "\r\nAnswer> ")

	RAnswer,err := Answer.ReadLine(); if err != nil {
		channel.Write([]byte("\r\nVoiding\r\n"))
		channel.Close()
		return err
	}

	AnswerQ := NumOne + NumTwo

	if RAnswer != strconv.Itoa(AnswerQ) {
		Used++
		goto fromthetop
	} else {
		return nil
	}
}