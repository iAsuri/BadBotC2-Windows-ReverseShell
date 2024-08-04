package slave

import "sync"

var (
	mut     sync.Mutex
	Clients = make(map[string]Client) // Mapping Users Session in memory
)
