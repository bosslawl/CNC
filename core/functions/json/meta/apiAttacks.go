package ParsedJson

import "reflect"

type AttacksFile struct {
	Blacklists []string       `json:"Blacklists"`
	Method     []AttackMethod `json:"Methods"`
}

type AttackMethod struct {
	Name        string   `json:"Name"`
	Description string   `json:"Description"`
	API         string   `json:"API"`
	Enabled     bool     `json:"Enabled"`
	Links       []Target `json:"Targets"`

	Management struct {
		DefaultPort    int  `json:"DefaultPort"`
		MaxConcurrents int  `json:"MaxConcurrents"`
		AdminMethod    bool `json:"AdminMethod"`
		VIPMethod      bool `json:"VIPMethod"`
		RawMethod      bool `json:"RawMethod"`
		HolderMethod   bool `json:"HolderMethod"`
		Timeout        int  `json:"Timeout"`
	} `json:"Management"`
}

type Target struct {
	Enabled   bool   `json:"Enabled"`
	Method    string `json:"Method"`
	Target    string `json:"Target"`
	URLEncode bool   `json:"URLEncode"`
	Debugging bool   `json:"Debugging"`
}

var _ = reflect.TypeOf(AttacksFile{})
var _ = reflect.TypeOf(AttackMethod{})
var _ = reflect.TypeOf(Target{})
