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
	return expected.ret0, expected.wait(ctx)
}

func (c *driverClient) Close(ctx context.Context) error {
	expected := &ExpectedClose{}
	if err := c.nextExpectation(expected); err != nil {
		return err
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
	return expected.wait(ctx)
}

func (c *driverClient) Ping(ctx context.Context) (bool, error) {
	expected := &ExpectedPing{}
	if err := c.nextExpectation(expected); err != nil {
		return false, err
	}
	return expected.ret0, expected.wait(ctx)
}
