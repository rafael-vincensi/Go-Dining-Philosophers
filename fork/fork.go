package fork

import "sync"

type Fork struct {
	ID    int
	Mutex sync.Mutex
}
