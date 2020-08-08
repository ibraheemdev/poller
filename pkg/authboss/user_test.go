package authboss

import "testing"

type testAssertionFailUser struct{}

func (testAssertionFailUser) GetPID() string { return "" }
func (testAssertionFailUser) PutPID(string)  {}

func TestUserAssertions(t *testing.T) {
	t.Parallel()

	u := &mockUser{}
	fu := testAssertionFailUser{}

	paniced := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				paniced = true
			}
		}()

		MustBeAuthable(u)
		MustBeConfirmable(u)
		MustBeLockable(u)
		MustBeRecoverable(u)
	}()

	if paniced {
		t.Error("The mock user should have included all interfaces and should not panic")
	}

	didPanic := func(f func()) (paniced bool) {
		defer func() {
			if r := recover(); r != nil {
				paniced = true
			}
		}()

		f()
		return paniced
	}

	if !didPanic(func() { MustBeAuthable(fu) }) {
		t.Error("should have panic'd")
	}
	if !didPanic(func() { MustBeConfirmable(fu) }) {
		t.Error("should have panic'd")
	}
	if !didPanic(func() { MustBeLockable(fu) }) {
		t.Error("should have panic'd")
	}
	if !didPanic(func() { MustBeRecoverable(fu) }) {
		t.Error("should have panic'd")
	}
}
