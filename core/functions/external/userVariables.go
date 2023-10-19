package External

import "sync"

var (
	Command = make(map[string]*Storage)
	Mutex sync.Mutex
)

type Storage struct {
	Name        string
	Description string
	Admin       bool
	Reseller    bool
	VIP         bool 
	Raw         bool
	Holder      bool
	
	Banner      []string
}
