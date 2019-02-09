package kivikmock

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type expectation interface {
	fulfill()
	fulfilled() bool
	Lock()
	Unlock()
	fmt.Stringer
	// method should return the name of the method that would trigger this
	// condition. If verbose is true, the output should disambiguate between
	// different calls to the same method.
	method(verbose bool) string
	error() error
	wait(context.Context) error
	// met is called on the actual value, and returns true if the expectation
	// is met.
	met(expectation) bool
}

// commonExpectation satisfies the expectation interface, except the String()
// and method() methods.
type commonExpectation struct {
	sync.Mutex
	triggered bool
	err       error // nolint: structcheck
	delay     time.Duration
}

func (e *commonExpectation) fulfill() {
	e.triggered = true
}

func (e *commonExpectation) fulfilled() bool {
	return e.triggered
}

func (e *commonExpectation) error() error {
	return e.err
}

// wait blocks until e.delay expires, or ctx is cancelled. If e.delay expires,
// e.err is returned, otherwise ctx.Err() is returned.
func (e *commonExpectation) wait(ctx context.Context) error {
	if e.delay == 0 {
		return e.err
	}
	t := time.NewTimer(e.delay)
	defer t.Stop()
	select {
	case <-t.C:
		return e.err
	case <-ctx.Done():
		return ctx.Err()
	}
}
