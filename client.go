package kivikmock

import (
	"context"
	"errors"
	"fmt"

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

	expected := &ExpectedAllDBs{
		options: opts,
	}
	if err := c.nextExpectation(expected); err != nil {
		return nil, err
	}

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
