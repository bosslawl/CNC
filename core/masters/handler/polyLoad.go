package Handler

import (
	"bufio"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
	"sync"

	CNC "Rain/core/config/admin"

	ini "github.com/vaughan0/go-ini"
	"golang.org/x/crypto/ssh"
)

var (
	BetaMapHandler = make(map[string]*CommandText)
	Handle         sync.Mutex
)

type CommandText struct {
	CommandName        string
	CommandAdmin       bool
	CommandReseller    bool
	CommandVip         bool
	CommandRaw         bool
	CommandHolder      bool
	CommandContains    string
	CommandDescription string
}

var _ = reflect.TypeOf(CommandText{})

var CommandFile string = CNC.CommandsFolder

func PolyLoader(channel ssh.Channel) bool {

	for i, _ := range BetaMapHandler {
		delete(BetaMapHandler, i)
	}

	Files, err := ioutil.ReadDir(CommandFile)
	if err != nil {
		return false
	}
	loaded := 0

	for _, f := range Files {
		file, _ := ini.LoadFile(CommandFile + f.Name())

		CommandNameLoad, ok := file.Get("", "name")
		if !ok {
			channel.Write([]byte("\x1b[0m[\x1b[1;31mFAILED\x1b[0m] Failed to get name from \"" + f.Name() + "\"\r\n"))
			continue
		}

		CommandDescriptionLoad, ok := file.Get("", "description")
		if !ok {
			channel.Write([]byte("\x1b[0m[\x1b[1;31mFAILED\x1b[0m] Failed to get Description from \"" + f.Name() + "\"\r\n"))
			continue
		}

		CommandAdminLoad, ok := file.Get("", "admin")
		if !ok {
			channel.Write([]byte("\x1b[0m[\x1b[1;31mFAILED\x1b[0m] Failed to Admin boolen from \"" + f.Name() + "\"\r\n"))
			continue
		}

		CommandResellerLoad, ok := file.Get("", "reseller")
		if !ok {
			channel.Write([]byte("\x1b[0m[\x1b[1;31mFAILED\x1b[0m] Failed to Reseller boolen from \"" + f.Name() + "\"\r\n"))
			continue
		}

		CommandVipLoad, ok := file.Get("", "vip")
		if !ok {
			channel.Write([]byte("\x1b[0m[\x1b[1;31mFAILED\x1b[0m] Failed to VIP boolen from \"" + f.Name() + "\"\r\n"))
			continue
		}

		CommandRawLoad, ok := file.Get("", "raw")
		if !ok {
			channel.Write([]byte("\x1b[0m[\x1b[1;31mFAILED\x1b[0m] Failed to Raw boolen from \"" + f.Name() + "\"\r\n"))
			continue
		}

		CommandHolderLoad, ok := file.Get("", "holder")
		if !ok {
			channel.Write([]byte("\x1b[0m[\x1b[1;31mFAILED\x1b[0m] Failed to Holder boolen from \"" + f.Name() + "\"\r\n"))
			continue
		}

		CommandAdminType, error := strconv.ParseBool(CommandAdminLoad)
		if error != nil {
			channel.Write([]byte("\x1b[0m[\x1b[1;31mFAILED\x1b[0m] Failed to Admin boolen from \"" + f.Name() + "\"\r\n"))
			continue
		}

		CommandResellerType, error := strconv.ParseBool(CommandResellerLoad)
		if error != nil {
			channel.Write([]byte("\x1b[0m[\x1b[1;31mFAILED\x1b[0m] Failed to Reseller boolen from \"" + f.Name() + "\"\r\n"))
			continue
		}

		CommandVipType, error := strconv.ParseBool(CommandVipLoad)
		if error != nil {
			channel.Write([]byte("\x1b[0m[\x1b[1;31mFAILED\x1b[0m] Failed to VIP boolen from \"" + f.Name() + "\"\r\n"))
			continue
		}

		CommandRawType, error := strconv.ParseBool(CommandRawLoad)
		if error != nil {
			channel.Write([]byte("\x1b[0m[\x1b[1;31mFAILED\x1b[0m] Failed to Raw boolen from \"" + f.Name() + "\"\r\n"))
			continue
		}

		CommandHolderType, error := strconv.ParseBool(CommandHolderLoad)
		if error != nil {
			channel.Write([]byte("\x1b[0m[\x1b[1;31mFAILED\x1b[0m] Failed to Holder boolen from \"" + f.Name() + "\"\r\n"))
			continue
		}

		Files, err := ioutil.ReadFile(CommandFile + f.Name())
		if err != nil {
			channel.Write([]byte("\x1b[0m[\x1b[1;31mFAILED\x1b[0m] Failed to load Texture from \"" + f.Name() + "\"\r\n"))
			continue
		}

		var Banner string = ""
		var FoundEnd bool = false
		scan := bufio.NewScanner(strings.NewReader(string(Files)))
		for scan.Scan() {

			if FoundEnd != true {
				if strings.Contains(scan.Text(), "MENU SPLIT DONE") {
					FoundEnd = true
					continue
				}
			} else {
				if Banner == "" {
					Banner = scan.Text() + "\n"
					continue
				}
				Banner = Banner + scan.Text() + "\n"
			}
		}

		var MapCommand = CommandText{
			CommandName:        CommandNameLoad,
			CommandAdmin:       CommandAdminType,
			CommandReseller:    CommandResellerType,
			CommandVip:         CommandVipType,
			CommandRaw:         CommandRawType,
			CommandHolder:      CommandHolderType,
			CommandContains:    Banner,
			CommandDescription: CommandDescriptionLoad,
		}

		Handle.Lock()
		BetaMapHandler[MapCommand.CommandName] = &MapCommand
		Handle.Unlock()

		loaded++
	}

	return true
}
