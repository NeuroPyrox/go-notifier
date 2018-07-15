package notifier

import (
	"testing"
)

func waitBlocks(l Listener) bool {
	select {
	case <-l.Wait():
		return false
	default:
		return true
	}
}

func TestListener_CheckShouldBeFalse_WhenFirstCreated(t *testing.T) {
	testNotifier := New()
	testListener1 := testNotifier.Listen()
	if testListener1.Check() {
		t.Error("Shouldn't detect any notifications when first created!")
	}
	testNotifier.Notify()
	testListener2 := testNotifier.Listen()
	if testListener2.Check() {
		t.Error("Shouldn't detect any notifications when first created!")
	}
}

func TestListener_CheckShouldBeTrue_AfterNotify(t *testing.T) {
	testNotifier := New()
	testListener := testNotifier.Listen()
	testNotifier.Notify()
	if !testListener.Check() {
		t.Error("Should've detected notification!")
	}
}

func TestListener_CheckShouldBeFalse_WhenCalledAgain(t *testing.T) {
	testNotifier := New()
	testListener := testNotifier.Listen()
	testListener.Check()
	if testListener.Check() {
		t.Error("Shouldn't detect any notifications when already checked!")
	}
	testNotifier.Notify()
	if !testListener.Check() {
		t.Error("Should've detected notification!")
	}
	if testListener.Check() {
		t.Error("Shouldn't detect any notifications when already checked!")
	}
}

func TestListener_CheckShouldOnlyBeTrueOnce_AfterMultipleNotifications(t *testing.T) {
	testNotifier := New()
	testListener := testNotifier.Listen()
	testNotifier.Notify()
	testNotifier.Notify()
	testNotifier.Notify()
	if !testListener.Check() {
		t.Error("Should've detected notification!")
	}
	if testListener.Check() {
		t.Error("Shouldn't detect any notifications when already checked!")
	}
}

func TestMultipleListeners_CheckShouldOnlyBeTrueForOne_WhenNotifyBetweenCreation(t *testing.T) {
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
}

func TestMultipleListeners_CheckShouldBeTrueForBoth_AfterNotify(t *testing.T) {
	testNotifier := New()
	testListener1 := testNotifier.Listen()
	testListener2 := testNotifier.Listen()
	testNotifier.Notify()
	if !testListener1.Check() {
		t.Error("Should've detected notification!")
	}
	if !testListener2.Check() {
		t.Error("Should've detected notification!")
	}
}

func TestMultipleListeners_CheckShouldOnlyBeTrueOnce_AfterMultipleNotifications(t *testing.T) {
	testNotifier := New()
	testListener1 := testNotifier.Listen()
	testListener2 := testNotifier.Listen()
	testNotifier.Notify()
	testNotifier.Notify()
	testNotifier.Notify()
	if !testListener1.Check() || !testListener2.Check() {
		t.Error("Should've detected notifications!")
	}
	if testListener1.Check() {
		t.Error("Shouldn't detect any notifications when already checked!")
	}
	if testListener2.Check() {
		t.Error("Shouldn't detect any notifications when already checked!")
	}
}

func TestListener_WaitShouldBlock_WhenFirstCreated(t *testing.T) {
	testNotifier := New()
	testListener := testNotifier.Listen()
	if !waitBlocks(testListener) {
		t.Error("Wait should block when first created!")
	}
}

func TestListener_WaitShouldNotBlock_AfterNotify(t *testing.T) {
	testNotifier := New()
	testListener := testNotifier.Listen()
	testNotifier.Notify()
	if waitBlocks(testListener) {
		t.Error("Wait should not block after notify!")
	}
}

func TestListener_WaitShouldNotStartBlocking_WhenCalledAgain(t *testing.T) {
	testNotifier := New()
	testListener := testNotifier.Listen()
	testNotifier.Notify()
	testListener.Wait()
	if waitBlocks(testListener) {
		t.Error("Wait should not start blocking when called again!")
	}
}

func TestListener_WaitShouldStartBlocking_AfterCheck(t *testing.T) {
	testNotifier := New()
	testListener := testNotifier.Listen()
	testNotifier.Notify()
	testListener.Check()
	if !waitBlocks(testListener) {
		t.Error("Wait should start blocking after check!")
	}
}

func TestListener_CheckShouldBeFalse_AfterAdvance(t *testing.T) {
	testNotifier := New()
	testListener := testNotifier.Listen()
	testNotifier.Notify()
	testNotifier.Notify()
	testNotifier.Notify()
	testListener.Advance()
	if testListener.Check() {
		t.Error("Check should be false after advance!")
	}
}

func TestListener_CheckShouldBeTrue_AfterAdvanceThenNotify(t *testing.T) {
	testNotifier := New()
	testListener := testNotifier.Listen()
	testListener.Advance()
	testNotifier.Notify()
	if !testListener.Check() {
		t.Error("Check should be true after advance then notify!")
	}
}

func TestListener_WaitShouldBlock_AfterAdvance(t *testing.T) {
	testNotifier := New()
	testListener := testNotifier.Listen()
	testNotifier.Notify()
	testListener.Advance()
	if !waitBlocks(testListener) {
		t.Error("Wait should block after advance!")
	}
}
