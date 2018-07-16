package notifier

type Notifier interface {
	Notify()
	Listen() Listener
}

type notifier struct {
	current *state
}

func New() Notifier {
	return &notifier{
		current: newState(),
	}
}

func (n *notifier) Notify() {
	n.current = n.current.update()
}

func (n *notifier) Listen() Listener {
	return newListener(n.current)
}
