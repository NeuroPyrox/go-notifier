package notifier

import (
	"sync"
)

type Listener interface {
	Check() bool
	Wait() <-chan struct{}
	Advance()
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

func (l *listener) Wait() <-chan struct{} {
	l.Lock()
	defer l.Unlock()
	return l.last.done
}

func (l *listener) Advance() {
	l.Lock()
	l.last = l.last.getCurrent()
	l.Unlock()
}
