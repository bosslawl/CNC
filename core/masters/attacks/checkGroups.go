package APIAttacks

import (
	ParseJson "Rain/core/functions/json"
	ParsedJson "Rain/core/functions/json/meta"
)

func Get(name string) *ParsedJson.AttackMethod {
	for I := 0; I < len(ParseJson.AttacksParse.Method); I++ {
		if ParseJson.AttacksParse.Method[I].Name == name {
			return &ParseJson.AttacksParse.Method[I]
		}
	}
	return nil
}
