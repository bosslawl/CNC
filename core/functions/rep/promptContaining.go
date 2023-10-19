package Replireadings

import (
	"bufio"
	"strconv"
	"strings"

	viewsMapped "Rain/core/config/views"
)

func GetFunctions(name string) (bool, string, string) {
	Scanner := bufio.NewScanner(strings.NewReader(viewsMapped.Branding[name]))

	var PassedSplit bool = false

	var ItemsFound = make(map[string]string)
	for Scanner.Scan() {
		if !PassedSplit {
			Q, A := Eval(Scanner.Text())
			ItemsFound[Q] = A
		}

		if strings.Contains(Scanner.Text(), "=============== SPLIT ===============") {
			PassedSplit = true
		}
	}
	lol, _ := strconv.ParseBool(ItemsFound["Fade"])
	return lol, ItemsFound["PresetOne"], ItemsFound["PresetTwo"]
}

// Returns the question and the answer involded in a line
func Eval(line string) (string, string) {
	Line := strings.Split(line, "=")
	if len(Line) < 1 {
		return "nil", "nil"
	}

	return Line[0], Line[1]
}
