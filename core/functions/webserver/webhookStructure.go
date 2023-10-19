package Discord

import "reflect"

type Model struct {
	Content string   `json:"content"`
	Embeds  []Embeds `json:"embeds"`
}

type Fields struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type Author struct {
	Name string `json:"name"`
}

type Footer struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url"`
}

type Thumbnail struct {
	URL string `json:"url"`
}

type Embeds struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Color       int       `json:"color"`
	Fields      []Fields  `json:"fields"`
	Author      Author    `json:"author"`
	Footer      Footer    `json:"footer"`
	Thumbnail   Thumbnail `json:"thumbnail"`
}

var _ = reflect.TypeOf(Model{})
var _ = reflect.TypeOf(Fields{})
var _ = reflect.TypeOf(Author{})
var _ = reflect.TypeOf(Footer{})
var _ = reflect.TypeOf(Embeds{})
