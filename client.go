package kivikmock

import (
	"context"
	"reflect"

	"github.com/go-kivik/kivik/driver"
)

var _ driver.ClientCloser = &kivikmock{}

func (c *kivikmock) Close(ctx context.Context) error {
	expected := &ExpectedClose{}
	if err := c.nextExpectation(expected); err != nil {
		return err
	}

	return expected.wait(ctx)
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
	expected := &ExpectedAuthenticate{
		authType: reflect.TypeOf(authenticator).Name(),
	}
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

func (c *kivikmock) DBExists(ctx context.Context, name string, options map[string]interface{}) (bool, error) {
	expected := &ExpectedDBExists{
		name:    name,
		options: options,
	}
	if err := c.nextExpectation(expected); err != nil {
		return false, err
	}
	return expected.exists, expected.wait(ctx)
}

func (c *kivikmock) DestroyDB(ctx context.Context, name string, options map[string]interface{}) error {
	expected := &ExpectedDestroyDB{
		name: name,
	}
	if err := c.nextExpectation(expected); err != nil {
		return err
	}
	return expected.wait(ctx)
}

func (c *kivikmock) DBsStats(ctx context.Context, names []string) ([]*driver.DBStats, error) {
	expected := &ExpectedDBsStats{
		names: names,
	}
	if err := c.nextExpectation(expected); err != nil {
		return nil, err
	}
	stats := make([]*driver.DBStats, len(expected.stats))
	for i, s := range expected.stats {
		stats[i] = kivikStats2driverStats(s)
	}
	return stats, expected.wait(ctx)
}

var _ driver.DBsStatser = &kivikmock{}
