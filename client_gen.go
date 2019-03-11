/* This file is auto-generated. Do not edit it! */

package kivikmock

import (
	"context"

	"github.com/go-kivik/kivik/driver"
)

var _ = &driver.Attachment{}

func (c *driverClient) AllDBs(ctx context.Context, options map[string]interface{}) ([]string, error) {
	expected := &ExpectedAllDBs{
		commonExpectation: commonExpectation{
			options: options,
		},
	}
	if err := c.nextExpectation(expected); err != nil {
		return nil, err
	}
	if expected.callback != nil {
		return expected.callback(ctx, options)
	}
	return expected.ret0, expected.wait(ctx)
}

func (c *driverClient) Close(ctx context.Context) error {
	expected := &ExpectedClose{}
	if err := c.nextExpectation(expected); err != nil {
		return err
	}
	if expected.callback != nil {
		return expected.callback(ctx)
	}
	return expected.wait(ctx)
}

func (c *driverClient) ClusterSetup(ctx context.Context, arg0 interface{}) error {
	expected := &ExpectedClusterSetup{
		arg0: arg0,
	}
	if err := c.nextExpectation(expected); err != nil {
		return err
	}
	if expected.callback != nil {
		return expected.callback(ctx, arg0)
	}
	return expected.wait(ctx)
}

func (c *driverClient) ClusterStatus(ctx context.Context, options map[string]interface{}) (string, error) {
	expected := &ExpectedClusterStatus{
		commonExpectation: commonExpectation{
			options: options,
		},
	}
	if err := c.nextExpectation(expected); err != nil {
		return "", err
	}
	if expected.callback != nil {
		return expected.callback(ctx, options)
	}
	return expected.ret0, expected.wait(ctx)
}

func (c *driverClient) DBExists(ctx context.Context, arg0 string, options map[string]interface{}) (bool, error) {
	expected := &ExpectedDBExists{
		arg0: arg0,
		commonExpectation: commonExpectation{
			options: options,
		},
	}
	if err := c.nextExpectation(expected); err != nil {
		return false, err
	}
	if expected.callback != nil {
		return expected.callback(ctx, arg0, options)
	}
	return expected.ret0, expected.wait(ctx)
}

func (c *driverClient) DestroyDB(ctx context.Context, arg0 string, options map[string]interface{}) error {
	expected := &ExpectedDestroyDB{
		arg0: arg0,
		commonExpectation: commonExpectation{
			options: options,
		},
	}
	if err := c.nextExpectation(expected); err != nil {
		return err
	}
	if expected.callback != nil {
		return expected.callback(ctx, arg0, options)
	}
	return expected.wait(ctx)
}

func (c *driverClient) Ping(ctx context.Context) (bool, error) {
	expected := &ExpectedPing{}
	if err := c.nextExpectation(expected); err != nil {
		return false, err
	}
	if expected.callback != nil {
		return expected.callback(ctx)
	}
	return expected.ret0, expected.wait(ctx)
}

func (c *driverClient) DB(ctx context.Context, arg0 string, options map[string]interface{}) (driver.DB, error) {
	expected := &ExpectedDB{
		arg0: arg0,
		commonExpectation: commonExpectation{
			options: options,
		},
	}
	if err := c.nextExpectation(expected); err != nil {
		return nil, err
	}
	expected.ret0.name = arg0
	if expected.callback != nil {
		return expected.callback(ctx, arg0, options)
	}
	return &driverDB{DB: expected.ret0}, expected.wait(ctx)
}

func (c *driverClient) DBUpdates(ctx context.Context) (driver.DBUpdates, error) {
	expected := &ExpectedDBUpdates{}
	if err := c.nextExpectation(expected); err != nil {
		return nil, err
	}
	if expected.callback != nil {
		return expected.callback(ctx)
	}
	return &driverDBUpdates{Context: ctx, Updates: expected.ret0}, expected.wait(ctx)
}

func (c *driverClient) DBsStats(ctx context.Context, arg0 []string) ([]*driver.DBStats, error) {
	expected := &ExpectedDBsStats{
		arg0: arg0,
	}
	if err := c.nextExpectation(expected); err != nil {
		return nil, err
	}
	if expected.callback != nil {
		return expected.callback(ctx, arg0)
	}
	return expected.ret0, expected.wait(ctx)
}

func (c *driverClient) Session(ctx context.Context) (*driver.Session, error) {
	expected := &ExpectedSession{}
	if err := c.nextExpectation(expected); err != nil {
		return nil, err
	}
	if expected.callback != nil {
		return expected.callback(ctx)
	}
	return expected.ret0, expected.wait(ctx)
}

func (c *driverClient) Version(ctx context.Context) (*driver.Version, error) {
	expected := &ExpectedVersion{}
	if err := c.nextExpectation(expected); err != nil {
		return nil, err
	}
	if expected.callback != nil {
		return expected.callback(ctx)
	}
	return expected.ret0, expected.wait(ctx)
}
