package kivikmock

import (
	"context"

	"github.com/go-kivik/kivik/driver"
)

var _ driver.ClientCloser = &kivikmock{}

func (c *kivikmock) Close(_ context.Context) error {
	expected := &ExpectedClose{}
	if err := c.nextExpectation(expected); err != nil {
		return err
	}

	return expected.err
}

func (c *kivikmock) AllDBs(ctx context.Context, opts map[string]interface{}) ([]string, error) {
	expected := &ExpectedAllDBs{
		options: opts,
	}
	if err := c.nextExpectation(expected); err != nil {
		return nil, err
	}

	return expected.results, expected.wait(ctx)
}

var _ driver.Authenticator = &kivikmock{}

func (c *kivikmock) Authenticate(ctx context.Context, authenticator interface{}) error {
	expected := &ExpectedAuthenticate{}
	if err := c.nextExpectation(expected); err != nil {
		return err
	}

	return expected.wait(ctx)
}

var _ driver.Cluster = &kivikmock{}

func (c *kivikmock) ClusterSetup(ctx context.Context, action interface{}) error {
	expected := &ExpectedClusterSetup{
		action: action,
	}
	if err := c.nextExpectation(expected); err != nil {
		return err
	}
	return expected.wait(ctx)
}

func (c *kivikmock) ClusterStatus(ctx context.Context, options map[string]interface{}) (string, error) {
	expected := &ExpectedClusterStatus{
		options: options,
	}
	if err := c.nextExpectation(expected); err != nil {
		return "", err
	}
	return expected.status, expected.wait(ctx)
}
