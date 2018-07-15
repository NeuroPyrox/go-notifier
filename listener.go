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
	last *state
}

func newListener(s *state) Listener {
	return &listener{
		last: s,
	}
}

func (l *listener) Check() bool {
	current := l.last.getCurrent()
	if current == l.last {
		return false
	}
	l.last = current
	return true
}

func (l *listener) Wait() <-chan struct{} {
	return l.last.done
}

func (l *listener) Advance() {
	l.Lock()
	l.last = l.last.getCurrent()
	l.Unlock()
}
