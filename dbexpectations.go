package kivikmock

import "time"

// ExpectedDBClose is used to manage *kivik.Client.Close expectation returned
// by Mock.ExpectClose.
type ExpectedDBClose struct {
	commonExpectation
}

func (e *ExpectedDBClose) method(v bool) string {
	if v {
		return "DB.Close(ctx)"
	}
	return "DB.Close()"
}

func (e *ExpectedDBClose) met(_ expectation) bool { return true }

// WillReturnError allows setting an error for *kivik.Client.Close action.
func (e *ExpectedDBClose) WillReturnError(err error) *ExpectedDBClose {
	e.err = err
	return e
}

func (e *ExpectedDBClose) String() string {
	extra := delayString(e.delay) + errorString(e.err)
	msg := "call to DB.Close()"
	if extra != "" {
		msg += " which:" + extra
	}
	return msg
}

// WillDelay will cause execution of Close() to delay by duration d.
func (e *ExpectedDBClose) WillDelay(d time.Duration) *ExpectedDBClose {
	e.delay = d
	return e
}
