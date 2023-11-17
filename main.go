package main

import (
	Branding "Rain/core/config/views/branding"
	Build "Rain/core/database/build"
	External "Rain/core/functions/external"
	ParseJson "Rain/core/functions/json"
	Server "Rain/core/functions/server"
	Attack_Groups "Rain/core/functions/util/attackGroups"
	Handler "Rain/core/masters/handler"

	//"io"

	"fmt"
	//"net/http"
	//"os"
	"strconv"
	//"time"

	"github.com/denisbrodbeck/machineid"

	"github.com/fatih/color"
)

func main() {

	error := ParseJson.Configuration_Parse()

	uuid, err := machineid.ProtectedID("RAIN")

	if err != nil {
		fmt.Println(color.WhiteString(" - ") + color.WhiteString("[") + color.BlueString("STATUS") + color.WhiteString("]") + color.WhiteString(" -> ") + color.WhiteString("[") + color.YellowString("UUID") + color.WhiteString(":") + color.RedString("FATAL") + color.WhiteString("]") + color.WhiteString(" -> ") + color.WhiteString("[") + color.RedString("Failed to get License UUID") + color.WhiteString("]"))
	}

	whilte := color.New(color.FgBlack)
	boldWhite := whilte.Add(color.BgWhite)

	blu := color.New(color.FgBlack)
	boldBlue := blu.Add(color.BgBlue)

	boldWhite.Print(" * ")
	boldBlue.Println(" Loading Configuration File... ")

	if error != nil {
		fmt.Println(color.WhiteString(" - ") + color.WhiteString("[") + color.BlueString("STATUS") + color.WhiteString("]") + color.WhiteString(" -> ") + color.WhiteString("[") + color.YellowString("CONFIG") + color.WhiteString(":") + color.RedString("FATAL") + color.WhiteString("]") + color.WhiteString(" -> ") + color.WhiteString("[") + color.RedString(error.Error()) + color.WhiteString("]"))
	} else {
		fmt.Println(color.WhiteString(" - ") + color.WhiteString("[") + color.BlueString("STATUS") + color.WhiteString("]") + color.WhiteString(" -> ") + color.WhiteString("[") + color.YellowString("CONFIG") + color.WhiteString(":") + color.GreenString("SUCCESS") + color.WhiteString("]"))
	}

	boldWhite.Print("\n * ")
	boldBlue.Println(" Loading License Key... ")

	fmt.Println(color.WhiteString(" - ") + color.WhiteString("[") + color.BlueString("LICENSING") + color.WhiteString("]") + color.WhiteString(" License Key") + color.WhiteString(": ") + color.MagentaString(ParseJson.ConfigParse.App.License))
	fmt.Println(color.WhiteString(" - ") + color.WhiteString("[") + color.BlueString("LICENSING") + color.WhiteString("]") + color.WhiteString(" Device UUID") + color.WhiteString(": ") + color.MagentaString(uuid))

	boldWhite.Print("\n * ")
	boldBlue.Println(" Loading CGI... ")

	error = ParseJson.Attacks_Parse()

	if error != nil {
		fmt.Println(color.WhiteString(" - ") + color.WhiteString("[") + color.BlueString("STATUS") + color.WhiteString("]") + color.WhiteString(" -> ") + color.WhiteString("[") + color.YellowString("ATTACKS") + color.WhiteString(":") + color.RedString("FATAL") + color.WhiteString("]") + color.WhiteString(" -> ") + color.WhiteString("[") + color.RedString(error.Error()) + color.WhiteString("]"))
	} else {
		fmt.Println(color.WhiteString(" - ") + color.WhiteString("[") + color.BlueString("STATUS") + color.WhiteString("]") + color.WhiteString(" -> ") + color.WhiteString("[") + color.YellowString("ATTACKS") + color.WhiteString(":") + color.GreenString("SUCCESS") + color.WhiteString("]"))
	}

	error = Attack_Groups.SortGroups()
	if error != nil {
		fmt.Println(color.WhiteString(" - ") + color.WhiteString("[") + color.BlueString("STATUS") + color.WhiteString("]") + color.WhiteString(" -> ") + color.WhiteString("[") + color.YellowString("GROUPS") + color.WhiteString(":") + color.RedString("FATAL") + color.WhiteString("]") + color.WhiteString(" -> ") + color.WhiteString("[") + color.RedString(error.Error()) + color.WhiteString("]"))
	} else {
		fmt.Println(color.WhiteString(" - ") + color.WhiteString("[") + color.BlueString("STATUS") + color.WhiteString("]") + color.WhiteString(" -> ") + color.WhiteString("[") + color.YellowString("GROUPS") + color.WhiteString(":") + color.GreenString("SUCCESS") + color.WhiteString("]"))
	}

	error = ParseJson.Plans_Parse()
	if error != nil {
		fmt.Println(color.WhiteString(" - ") + color.WhiteString("[") + color.BlueString("STATUS") + color.WhiteString("]") + color.WhiteString(" -> ") + color.WhiteString("[") + color.YellowString("PLANS") + color.WhiteString(":") + color.RedString("FATAL") + color.WhiteString("]") + color.WhiteString(" -> ") + color.WhiteString("[") + color.RedString(error.Error()) + color.WhiteString("]"))
	} else {
		fmt.Println(color.WhiteString(" - ") + color.WhiteString("[") + color.BlueString("STATUS") + color.WhiteString("]") + color.WhiteString(" -> ") + color.WhiteString("[") + color.YellowString("PLANS") + color.WhiteString(":") + color.GreenString("SUCCESS") + color.WhiteString("]"))
	}

	Items_Length, error := Branding.Load_Items()
	if error != nil {
		fmt.Println(color.WhiteString(" - ") + color.WhiteString("[") + color.BlueString("STATUS") + color.WhiteString("]") + color.WhiteString(" -> ") + color.WhiteString("[") + color.YellowString("CGI") + color.WhiteString(":") + color.RedString("FATAL") + color.WhiteString("]") + color.WhiteString(" -> ") + color.WhiteString("[") + color.RedString(error.Error()) + color.WhiteString("]"))
		return
	} else {
		fmt.Println(color.WhiteString(" - ") + color.WhiteString("[") + color.BlueString("STATUS") + color.WhiteString("]") + color.WhiteString(" -> ") + color.WhiteString("[") + color.YellowString("CGI") + color.WhiteString(":") + color.GreenString("SUCCESS") + color.WhiteString("]"))
		fmt.Println(color.WhiteString(" - ") + color.WhiteString("[") + color.BlueString("BRANDING") + color.WhiteString("]") + color.WhiteString(" TFX Files") + color.WhiteString(": ") + color.MagentaString(strconv.Itoa(Items_Length)))
	}

	boldWhite.Println("\n * ")
	boldBlue.Println(" Loading Database... ")

	Boolen := Build.NewMongo()
	if !Boolen {
		fmt.Println(color.WhiteString(" - ") + color.WhiteString("[") + color.BlueString("STATUS") + color.WhiteString("]") + color.WhiteString(" -> ") + color.WhiteString("[") + color.RedString("Failed to locate database tables") + color.WhiteString("]"))

		error = Build.InsertTables()

		if error != nil {
			fmt.Println(color.WhiteString(" - ") + color.WhiteString("[") + color.BlueString("STATUS") + color.WhiteString("]") + color.WhiteString(" -> ") + color.WhiteString("[") + color.RedString("Failed to correctly insert database tables") + color.WhiteString("]"))
			return
		} else {
			fmt.Println(color.WhiteString(" - ") + color.WhiteString("[") + color.BlueString("STATUS") + color.WhiteString("]") + color.WhiteString(" -> ") + color.WhiteString("[") + color.GreenString("Correctly inserted database tables") + color.WhiteString("]"))
		}
	} else {
		fmt.Println(color.WhiteString(" - ") + color.WhiteString("[") + color.BlueString("STATUS") + color.WhiteString("]") + color.WhiteString(" -> ") + color.WhiteString("[") + color.GreenString("Correctly located database tables") + color.WhiteString("]"))
	}
	fmt.Println(color.WhiteString(" - ") + color.WhiteString("[") + color.BlueString("DATABASE") + color.WhiteString("]") + color.WhiteString(" Mongo URL") + color.WhiteString(": ") + color.MagentaString(ParseJson.ConfigParse.Database.MongoURL))

	boldWhite.Print("\n * ")
	boldBlue.Println(" Loading SSH... ")

	Handler.OfflineLoader()

	External.GatherExCommands()

	Server.New()
}
