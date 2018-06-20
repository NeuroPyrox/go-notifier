package notifier

type state struct {
	next *state
	done chan struct{}
}

func newState() *state {
	return &state{
		done: make(chan struct{}),
	}
}

func (s *state) update() *state {
	s.next = newState()
	close(s.done)
	return s.next
}

func (s *state) getCurrent() *state {
	select {
	case <-s.done:
		return s.next.getCurrent()
	default:
		return s
	}
}
