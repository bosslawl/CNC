package ParseJson

import (
	"encoding/json"
	"io/ioutil"
	"os"

	CNC "Rain/core/config/admin"
	ParsedJson "Rain/core/functions/json/meta"
)

var ConfigParse ParsedJson.ConfigFile

func Configuration_Parse() error {

	NewFile, error := os.Open(CNC.BuildFolder + "/config.json")
	if error != nil {
		return error
	}

	defer NewFile.Close()

	ByteValFile, error := ioutil.ReadAll(NewFile)
	if error != nil {
		return error
	}

	var New ParsedJson.ConfigFile
	error = json.Unmarshal(ByteValFile, &New)
	if error != nil {
		return error
	}

	ConfigParse = New

	return nil
}
