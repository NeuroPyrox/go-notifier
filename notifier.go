package notifier

import (
	"sync"
)

type Notifier interface {
	Notify()
	Listen() Listener
}

type notifier struct {
	sync.RWMutex
	current *state
}

func New() Notifier {
	return &notifier{
		current: newState(),
	}
}

func (n *notifier) Notify() {
	n.Lock()
	defer n.Unlock()
	n.current = n.current.update()
}

func (n *notifier) Listen() Listener {
	n.RLock()
	defer n.RUnlock()
	return newListener(n.current)
}
