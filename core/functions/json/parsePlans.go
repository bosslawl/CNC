package ParseJson

import (
	"encoding/json"
	"io/ioutil"
	"os"

	CNC "Rain/core/config/admin"
	ParsedJson "Rain/core/functions/json/meta"
)

var PlansParse ParsedJson.PlansFile

func Plans_Parse() error {

	NewFile, error := os.Open(CNC.BuildFolder + "/plan-presets.json")
	if error != nil {
		return error
	}

	defer NewFile.Close()

	ByteValFile, error := ioutil.ReadAll(NewFile)
	if error != nil {
		return error
	}

	var New ParsedJson.PlansFile
	error = json.Unmarshal(ByteValFile, &New)
	if error != nil {
		return error
	}

	PlansParse = New

	return nil
}
