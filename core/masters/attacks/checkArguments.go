package APIAttacks

import (
	ParsedJson "Rain/core/functions/json/meta"
	"net/url"
	"strconv"
	"strings"
)

func Get_Links(Method *ParsedJson.AttackMethod, Attack *MethodAttacks) []string {

	var TargetLinks []string
	for _, I := range Method.Links {
		Link := I.Target
		if I.Enabled {
			if I.URLEncode {
				Link = strings.Replace(Link, "[target]", url.QueryEscape(Attack.Target), -1)
				Link = strings.Replace(Link, "[port]", url.QueryEscape(strconv.Itoa(Attack.Port)), -1)
				Link = strings.Replace(Link, "[duration]", url.QueryEscape(strconv.Itoa(Attack.Duration)), -1)
				Link = strings.Replace(Link, "[method]", url.QueryEscape(I.Method), -1)
			} else {
				Link = strings.Replace(Link, "[target]", Attack.Target, -1)
				Link = strings.Replace(Link, "[port]", strconv.Itoa(Attack.Port), -1)
				Link = strings.Replace(Link, "[duration]", strconv.Itoa(Attack.Duration), -1)
				Link = strings.Replace(Link, "[method]", I.Method, -1)
			}

			TargetLinks = append(TargetLinks, Link)

		}
	}

	return TargetLinks
}
