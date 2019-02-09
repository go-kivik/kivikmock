package kivikmock

import (
	"context"
	"fmt"
	"reflect"
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

// nextExpectation accepts the expected value `e`, checks that this is a valid
// expectation, and if so, populates e with the matching expectation. If the
// expectation is not expected, an error is returned.
func (c *kivikmock) nextExpectation(actual expectation) error {
	c.drv.Lock()
	defer c.drv.Unlock()

	var expected expectation
	var fulfilled int
	for _, next := range c.expected {
		next.Lock()
		if next.fulfilled() {
			next.Unlock()
			fulfilled++
			continue
		}

		if c.ordered {
			if reflect.TypeOf(actual).Elem().Name() == reflect.TypeOf(next).Elem().Name() {
				if meets(actual, next) {
					expected = next
					break
				}
				next.Unlock()
				return fmt.Errorf("Expectation not met:\nExpected: %s\n  Actual: %s",
					next, actual)
			}
			next.Unlock()
			return fmt.Errorf("call to %s was not expected. Next expectation is: %s", actual.method(false), next.method(false))
		}
		if meets(actual, next) {
			expected = next
			break
		}

		next.Unlock()
	}

	if expected == nil {
		if fulfilled == len(c.expected) {
			return fmt.Errorf("call to %s was not expected, all expectations already fulfilled", actual.method(false))
		}
		return fmt.Errorf("call to %s was not expected", actual.method(!c.ordered))
	}

	defer expected.Unlock()
	expected.fulfill()

	reflect.ValueOf(actual).Elem().Set(reflect.ValueOf(expected).Elem())
	return nil
}

func meets(a, e expectation) bool {
	if reflect.TypeOf(a).Elem().Name() != reflect.TypeOf(e).Elem().Name() {
		return false
	}
	return a.met(e)
}
