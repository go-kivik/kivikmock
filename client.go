package kivikmock

import (
	"context"

	"github.com/go-kivik/kivik/driver"
)

var _ driver.ClientCloser = &kivikmock{}

func (c *kivikmock) Close(_ context.Context) error {
	c.drv.Lock()
	defer c.drv.Unlock()

	expected := &ExpectedClose{}
	if err := c.nextExpectation(expected); err != nil {
		return err
	}

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
