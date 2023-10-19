package ParseJson

import (
	"encoding/json"
	"io/ioutil"
	"os"

	CNC "Rain/core/config/admin"
	ParsedJson "Rain/core/functions/json/meta"
)

var AttacksParse ParsedJson.AttacksFile

func Attacks_Parse() error {

	NewFile, error := os.Open(CNC.BuildFolder + "/attacks/attacks.json")
	if error != nil {
		return error
	}

	defer NewFile.Close()

	ByteValFile, error := ioutil.ReadAll(NewFile)
	if error != nil {
		return error
	}

	var New ParsedJson.AttacksFile
	error = json.Unmarshal(ByteValFile, &New)
	if error != nil {
		return error
	}

	AttacksParse = New

	return nil
}
