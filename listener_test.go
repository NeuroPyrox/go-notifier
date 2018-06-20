package notifier

import (
	"testing"
)

func TestSingleListener(t *testing.T) {
	testNotifier := New()
	testNotifier.Notify()
	testListener := testNotifier.Listen()
	if testListener.Check() {
		t.Error("Shouldn't detect any notifications before being created!")
	}
	testNotifier.Notify()
	if !testListener.Check() {
		t.Error("Should've detected notification!")
	}
	if testListener.Check() {
		t.Error("Shouldn't detect a notification twice!")
	}
	testNotifier.Notify()
	testNotifier.Notify()
	testNotifier.Notify()
	if !testListener.Check() {
		t.Error("Should've detected notification!")
	}
	if testListener.Check() {
		t.Error("Should've cleared all notifications!")
	}
}

func TestMultipleListeners(t *testing.T) {
	testNotifier := New()
	testListener1 := testNotifier.Listen()
	testNotifier.Notify()
	testListener2 := testNotifier.Listen()
	if testListener2.Check() {
		t.Error("Shouldn't detect any notifications before being created!")
	}
	if !testListener1.Check() {
		t.Error("Should've detected notification!")
	}
	testNotifier.Notify()
	if !testListener1.Check() {
		t.Error("Should've detected notification!")
	}
	if !testListener2.Check() {
		t.Error("Should've detected notification!")
	}
}
