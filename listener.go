package notifier

import (
	"sync"
)

type Listener interface {
	Check() bool
}

type listener struct {
	sync.Mutex
	last *state
}

func newListener(s *state) Listener {
	return &listener{
		last: s,
	}
}

func (l *listener) Check() bool {
	l.Lock()
	current := l.last.getCurrent()
	if current == l.last {
		l.Unlock()
		return false
	}
	l.last = current
	l.Unlock()
	return true
}
