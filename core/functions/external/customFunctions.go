package External

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	CNC "Rain/core/config/admin"
)

func GatherExCommands() []error {

	for key, _ := range Command {
		delete(Command, key)
	}

	var Errors []error
	Walk, error := ioutil.ReadDir(CNC.CommandsFolder)
	if error != nil {
		return append(Errors, error)
	}

	for _, File := range Walk {
		BannerFile := GetFileContain(File)

		Defines, Banner := SplitBySplit(BannerFile)

		var New = Storage{}

		New.Banner = Banner
		New = *GetDefines(&New, Defines)

		Mutex.Lock()
		Command[New.Name] = &New
		Mutex.Unlock()
	}

	return Errors

}

func GetDefines(cmd *Storage, Define []string) *Storage {

	for _, Text := range Define {

		switch strings.Split(Text, "=")[0] {
		case "name":
			cmd.Name = strings.Split(Text, "=")[1]
		case "description":
			cmd.Description = strings.Split(Text, "=")[1]
		case "admin":
			AdminBoolen, _ := strconv.ParseBool(strings.Split(Text, "=")[1])
			cmd.Admin = AdminBoolen
		case "reseller":
			ResellerBoolen, _ := strconv.ParseBool(strings.Split(Text, "=")[1])
			cmd.Reseller = ResellerBoolen
		case "vip":
			VipBoolen, _ := strconv.ParseBool(strings.Split(Text, "=")[1])
			cmd.VIP = VipBoolen
		case "raw":
			RawBoolen, _ := strconv.ParseBool(strings.Split(Text, "=")[1])
			cmd.Raw = RawBoolen
		case "holder":
			HolderBoolen, _ := strconv.ParseBool(strings.Split(Text, "=")[1])
			cmd.Holder = HolderBoolen
		}
	}

	return cmd
}

func GetFileContain(File os.FileInfo) []string {
	FileLocker, error := os.Open(CNC.CommandsFolder + File.Name())
	if error != nil {
		return nil
	}

	Scanner := bufio.NewScanner(FileLocker)

	var Banner []string

	for Scanner.Scan() {
		Banner = append(Banner, Scanner.Text())
	}

	return Banner
}

// MENU SPLIT DONE
func SplitBySplit(Banner []string) ([]string, []string) {
	var Pass_split bool = false

	var Defines []string
	var BannerLine []string

	for _, Text := range Banner {

		if strings.Contains(Text, "MENU SPLIT DONE") {
			Pass_split = true
			continue
		}

		if Pass_split {
			BannerLine = append(BannerLine, Text)
		} else if !Pass_split {
			Defines = append(Defines, Text)
		}
	}

	return Defines, BannerLine
}
