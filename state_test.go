package notifier

import (
	"testing"
)

func TestState(t *testing.T) {
	current := newState()
	if current.getCurrent() != current {
		t.Error("Current state of new state is not self!")
	}
	states := []*state{current}
	numTestStates := 5
	for i := 1; i < numTestStates; i++ {
		current = current.update()
		states = append(states, current)
		for _, testState := range states {
			if testState.getCurrent() != current {
				t.Error("Current state of test state is wrong")
			}
		}
	}
}
