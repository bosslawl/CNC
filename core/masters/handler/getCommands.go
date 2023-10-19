package Handler

func Command(name string) *CommandText {
	for _, s := range BetaMapHandler {
		if s.CommandName == name {
			return s
		}
	}
	return nil
}