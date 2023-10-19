package viewsMapped

import "sync"

var (
	Branding = make(map[string]string)
	NyxMux   sync.Mutex
)