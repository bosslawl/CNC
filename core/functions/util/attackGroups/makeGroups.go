package Attack_Groups

import (
	"reflect"
	"strconv"
	"sync"

	ParseJson "Rain/core/functions/json"
)

type AttackGroup struct {
	Methods []Method
	Voting  []Vote
}

type Vote struct {
	Name      string
	Type_vote string
}

type Method struct {
	Name           string
	Description    string
	VIP            bool
	Raw            bool
	Holder         bool
	DefaultPort    string
	Concurrents    int
	MaxConcurrents int
}

var _ = reflect.TypeOf(AttackGroup{})
var _ = reflect.TypeOf(Vote{})
var _ = reflect.TypeOf(Method{})

var (
	Attk_groups = make(map[string]*AttackGroup)
	MuxLock     sync.Mutex
)

func Sort_Voting(api string) (int, int) {
	api_group := Attk_groups[api]
	if api_group == nil {
		return 0, 0
	}

	var voted []string

	var pos int = 0
	var posvoter []string

	var neg int = 0
	var negvoter []string

	for _, l := range api_group.Voting {
		if l.Type_vote == "positive" {
			if !Check(l.Name, negvoter) && !Check(l.Name, voted) {
				posvoter = append(posvoter, l.Name)
				voted = append(voted, l.Name)
				pos++
			}
		} else {
			if !Check(l.Name, posvoter) && !Check(l.Name, voted) {
				negvoter = append(negvoter, l.Name)
				voted = append(voted, l.Name)
				neg++
			}

		}
	}

	return pos, neg
}

func Check(val string, array []string) bool {
	for _, a := range array {
		if a == val {
			return true
		}
	}

	return false
}

func SortGroups() error {

	for Cli, _ := range Attk_groups {
		delete(Attk_groups, Cli)
	}

	for I := 0; I < len(ParseJson.AttacksParse.Method); I++ {
		Cli, Get := Attk_groups[ParseJson.AttacksParse.Method[I].API]
		if !Get {
			var New = AttackGroup{
				Methods: []Method{
					Method{
						Name:           ParseJson.AttacksParse.Method[I].Name,
						Description:    ParseJson.AttacksParse.Method[I].Description,
						VIP:            ParseJson.AttacksParse.Method[I].Management.VIPMethod,
						Raw:            ParseJson.AttacksParse.Method[I].Management.RawMethod,
						Holder:         ParseJson.AttacksParse.Method[I].Management.HolderMethod,
						DefaultPort:    ":" + strconv.Itoa(ParseJson.AttacksParse.Method[I].Management.DefaultPort),
						Concurrents:    len(ParseJson.AttacksParse.Method[I].Links),
						MaxConcurrents: ParseJson.AttacksParse.Method[I].Management.MaxConcurrents,
					},
				},
			}

			MuxLock.Lock()
			Attk_groups[ParseJson.AttacksParse.Method[I].API] = &New
			MuxLock.Unlock()
		} else {
			var New = Method{
				Name:           ParseJson.AttacksParse.Method[I].Name,
				Description:    ParseJson.AttacksParse.Method[I].Description,
				VIP:            ParseJson.AttacksParse.Method[I].Management.VIPMethod,
				Raw:            ParseJson.AttacksParse.Method[I].Management.RawMethod,
				Holder:         ParseJson.AttacksParse.Method[I].Management.HolderMethod,
				DefaultPort:    ":" + strconv.Itoa(ParseJson.AttacksParse.Method[I].Management.DefaultPort),
				Concurrents:    len(ParseJson.AttacksParse.Method[I].Links),
				MaxConcurrents: ParseJson.AttacksParse.Method[I].Management.MaxConcurrents,
			}
			Cli.Methods = append(Cli.Methods, New)
		}
	}

	return nil
}
