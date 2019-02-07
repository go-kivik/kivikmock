package kivikmock

import (
	"fmt"
	"sync"
)

type expectation interface {
	fulfilled() bool
	Lock()
	Unlock()
	String() string
}

// commonExpectation satisfies the expectation interface, except the String()
// method.
type commonExpectation struct {
	sync.Mutex
	triggered bool
	err       error
}

func (e *commonExpectation) fulfilled() bool {
	return e.triggered
}

// ExpectedClose is used to manage *kivik.Client.Close expectation returned
// by Mock.ExpectClose.
type ExpectedClose struct {
	commonExpectation
}

// WillReturnError allows setting an error for *kivik.Client.Close action.
func (e *ExpectedClose) WillReturnError(err error) *ExpectedClose {
	e.err = err
	return e
}

func (e *ExpectedClose) String() string {
	msg := "ExpectedClose => expecting client Close"
	if e.err != nil {
		return fmt.Sprintf("%s, which should return error: %s", msg, e.err)
	}
	return msg
}
