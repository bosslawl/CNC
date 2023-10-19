package APIAttacks

import (
	ParseJson "Rain/core/functions/json"
	"strings"
)

func CheckBlacklist(tgt string) bool {
	for _, target := range ParseJson.AttacksParse.Blacklists {
		if strings.ToLower(target) == strings.ToLower(tgt) {
			return true
		}
	}
	return false
}
