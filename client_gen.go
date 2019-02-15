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

func (c *driverClient) DBUpdates() (driver.DBUpdates, error) {
	expected := &ExpectedDBUpdates{}
	if err := c.nextExpectation(expected); err != nil {
		return nil, err
	}
	return expected.ret0, expected.err
}

func (c *driverClient) DBsStats(ctx context.Context, arg0 []string) ([]*driver.DBStats, error) {
	expected := &ExpectedDBsStats{
		arg0: arg0,
	}
	if err := c.nextExpectation(expected); err != nil {
		return nil, err
	}
	return expected.ret0, expected.wait(ctx)
}

func (c *driverClient) GetReplications(ctx context.Context, options map[string]interface{}) ([]driver.Replication, error) {
	expected := &ExpectedGetReplications{
		commonExpectation: commonExpectation{
			options: options,
		},
	}
	if err := c.nextExpectation(expected); err != nil {
		return nil, err
	}
	return expected.ret0, expected.wait(ctx)
}

func (c *driverClient) Replicate(ctx context.Context, arg0 string, arg1 string, options map[string]interface{}) (driver.Replication, error) {
	expected := &ExpectedReplicate{
		arg0: arg0,
		arg1: arg1,
		commonExpectation: commonExpectation{
			options: options,
		},
	}
	if err := c.nextExpectation(expected); err != nil {
		return nil, err
	}
	return expected.ret0, expected.wait(ctx)
}

func (c *driverClient) Session(ctx context.Context) (*driver.Session, error) {
	expected := &ExpectedSession{}
	if err := c.nextExpectation(expected); err != nil {
		return nil, err
	}
	return expected.ret0, expected.wait(ctx)
}

func (c *driverClient) Version(ctx context.Context) (*driver.Version, error) {
	expected := &ExpectedVersion{}
	if err := c.nextExpectation(expected); err != nil {
		return nil, err
	}
	return expected.ret0, expected.wait(ctx)
}
