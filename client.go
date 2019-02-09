package kivikmock

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/go-kivik/kivik/driver"
)

var _ driver.ClientCloser = &kivikmock{}

func (c *kivikmock) Close(_ context.Context) error {
	c.drv.Lock()
	defer c.drv.Unlock()

	c.opened--
	if c.opened == 0 {
		delete(c.drv.clients, c.dsn)
	}

	var expected *ExpectedClose
	var fulfilled int
	var ok bool
	for _, next := range c.expected {
		next.Lock()
		if next.fulfilled() {
			next.Unlock()
			fulfilled++
			continue
		}

		if expected, ok = next.(*ExpectedClose); ok {
			break
		}
		next.Unlock()
		if c.ordered {
			return fmt.Errorf("call to client Close was not expected. Next expectation is: %s", next)
		}
	}
	if expected == nil {
		msg := "call to client Close was not expected"
		if fulfilled == len(c.expected) {
			msg = "all expectations were already fulfilled, " + msg
		}
		return errors.New(msg)
	}

	expected.triggered = true
	expected.Unlock()
	return expected.err
}

func (c *kivikmock) AllDBs(ctx context.Context, opts map[string]interface{}) ([]string, error) {
	c.drv.Lock()
	defer c.drv.Unlock()

	errMsg := func() string {
		if opts != nil {
			return fmt.Sprintf("call to AllDBs with options %+v", opts)
		}
		return "call to AllDBs"
	}

	var expected *ExpectedAllDBs
	var fulfilled int
	var ok bool
	for _, next := range c.expected {
		next.Lock()
		if next.fulfilled() {
			next.Unlock()
			fulfilled++
			continue
		}

		if c.ordered {
			if expected, ok = next.(*ExpectedAllDBs); ok {
				break
			}
			next.Unlock()
			msg := errMsg()
			return nil, fmt.Errorf(msg+" was not expected. Next expectation is: %s", next)
		}
		if e, ok := next.(*ExpectedAllDBs); ok {
			if reflect.DeepEqual(opts, e.options) {
				expected = e
				break
			}
		}
		next.Unlock()
	}

	if expected == nil {
		msg := errMsg()
		msg += " was not expected"
		if fulfilled == len(c.expected) {
			msg = "all expectations were already fulfilled, " + msg
		}
		return nil, errors.New(msg)
	}

	defer expected.Unlock()
	expected.triggered = true

	return expected.results, expected.err
}

var _ driver.Authenticator = &kivikmock{}

func (c *kivikmock) Authenticate(ctx context.Context, authenticator interface{}) error {
	c.drv.Lock()
	defer c.drv.Unlock()

	expected := &ExpectedAuthenticate{}
	if err := c.nextExpectation(expected); err != nil {
		return err
	}

	return expected.err
}
