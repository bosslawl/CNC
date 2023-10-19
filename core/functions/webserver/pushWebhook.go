package Discord

import (
	"bytes"
	"encoding/json"

	// "strconv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/fatih/color"

	ParseJson "Rain/core/functions/json"
)

func DiscordAttackAlert(user string, target string, port string, duration string, method string) error {

	return LogAlert("New Attack", []Fields{
		{
			Name:   "**Target**",
			Value:  target,
			Inline: false,
		},

		{
			Name:   "**Duration**",
			Value:  duration,
			Inline: false,
		},

		{
			Name:   "**Port**",
			Value:  port,
			Inline: false,
		},

		{
			Name:   "**Method**",
			Value:  method,
			Inline: true,
		},

		{
			Name:   "**Username**",
			Value:  user,
			Inline: true,
		},
	})
}

/* func telelog(key string, target string, port int, duration int, method string, client string, cons int, current int, bot string, chatid string, myClient *http.Client) { // Telegram logging
    tele := "https://api.telegram.org/bot"+bot+"/sendMessage?chat_id="+chatid+"&text=Attack%20Successfully%20Sent%0AKey:%20"+key+"%0ATarget:%20"+target+"%0APort:%20"+strconv.Itoa(port)+"%0ADuration:%20"+strconv.Itoa(duration)+"%0AMethod:%20"
    _, err := myClient.Get(tele)
    if err != nil {
		log.Fatal(color.WhiteString(" - ") + color.WhiteString("[") + color.BlueString("STATUS") + color.WhiteString("]") + color.WhiteString(" -> ") + color.WhiteString("[") + color.YellowString("LOGGING") + color.WhiteString(":") + color.RedString("FATAL") + color.WhiteString("]") + color.WhiteString(" -> ") + color.WhiteString("[") + color.RedString("Cannot push telegram log") + color.WhiteString("]"))
    }
} */

func LogAlert(alert_name string, F []Fields) error {
	year, month, day := time.Now().Date()
	date := fmt.Sprintf("%d/%s/%d", day, month, year)
	data, err := json.Marshal(Model{
		Embeds: []Embeds{
			{
				Title:       fmt.Sprintf("Attack Distributed Successfully"),
				Description: fmt.Sprintf("Date: %s", date),
				Color:       0x77f59c,
				Fields:      F,
				Author: Author{
					fmt.Sprintf(ParseJson.ConfigParse.App.AppName),
				},
			},
		},
	})
	if err != nil {
		return err
	}
	rq, err := http.Post(ParseJson.ConfigParse.App.WebHook, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	defer rq.Body.Close()

	if rq.StatusCode > 204 {
		log.Print(err)
		_, err := ioutil.ReadAll(rq.Body)
		if err != nil {
			return err
		}
		log.Fatal(color.WhiteString(" - ") + color.WhiteString("[") + color.BlueString("STATUS") + color.WhiteString("]") + color.WhiteString(" -> ") + color.WhiteString("[") + color.YellowString("LOGGING") + color.WhiteString(":") + color.RedString("FATAL") + color.WhiteString("]") + color.WhiteString(" -> ") + color.WhiteString("[") + color.RedString("Cannot push discord log") + color.WhiteString("]"))
	}
	return nil
}
