package ParsedJson

import "reflect"

type PlansFile struct {
	Plan []Plans `json:"Plans"`
}

type Plans struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Price       string `json:"Price"`

	Admin           bool `json:"Admin"`
	PowerSaving     bool `json:"PowerSaving"`
	Bypassblacklist bool `json:"Bypassblacklist"`
	Reseller        bool `json:"Reseller"`
	Banned          bool `json:"Banned"`
	VIP             bool `json:"VIP"`
	Raw             bool `json:"Raw"`
	Holder          bool `json:"Holder"`

	MaxSessions int `json:"MaxSessions"`
	MaxTime     int `json:"Maxime"`
	Cooldown    int `json:"Cooldown"`
	Concurrents int `json:"Concurrents"`
	DaysActive  int `json:"DaysActive"`
}

var _ = reflect.TypeOf(PlansFile{})
var _ = reflect.TypeOf(Plans{})
