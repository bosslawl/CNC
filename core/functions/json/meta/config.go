package ParsedJson

import "reflect"

type ConfigFile struct {
	Masters struct {
		MastersPort    string `json:"MastersPort"`
		MastersMaxAuth int    `json:"MastersMaxAuth"`
		MastersKeyFile string `json:"MastersKeyFile"`
	} `json:"Masters"`

	Database struct {
		MongoURL string `json:"MongoURL"`
	} `json:"Database"`

	App struct {
		AppName string `json:"AppName"`
		WebHook string `json:"Webhook"`
		License string `json:"License"`
	} `json:"App"`

	Controls struct {
		Catpcha struct {
			Status          bool   `json:"Status"`
			Header          string `json:"Header"`
			AllowedAttempts int    `json:"AllowedAttempts"`
			AdminBypass     bool   `json:"AdminBypass"`
			Question        struct {
				MinGen int `json:"MinGen"`
				MaxGen int `json:"MaxGen"`
			} `json:"Question"`
		} `json:"Catpcha"`

		MFA struct {
			Status      bool `json:"Status"`
			AdminBypass bool `json:"AdminBypass"`
		} `json:"MFA"`
	} `json:"Controls"`
}

var _ = reflect.TypeOf(ConfigFile{})
