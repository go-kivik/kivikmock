package kivikmock

import (
	"context"
	"reflect"

	"github.com/go-kivik/kivik/driver"
)

type driverClient struct {
	*MockClient
}

var _ driver.Client = &driverClient{}
var _ driver.ClientCloser = &driverClient{}
var _ driver.Authenticator = &driverClient{}
var _ driver.Cluster = &driverClient{}
var _ driver.DBsStatser = &driverClient{}
var _ driver.Pinger = &driverClient{}
var _ driver.Sessioner = &driverClient{}

func (c *driverClient) Authenticate(ctx context.Context, authenticator interface{}) error {
	expected := &ExpectedAuthenticate{
		authType: reflect.TypeOf(authenticator).Name(),
	}
	if err := c.nextExpectation(expected); err != nil {
		return err
	}
	if expected.callback != nil {
		return expected.callback(ctx, authenticator)
	}
	return expected.wait(ctx)
}

func (c *driverClient) CreateDB(ctx context.Context, name string, options map[string]interface{}) error {
	expected := &ExpectedCreateDB{
		arg0: name,
		commonExpectation: commonExpectation{
			options: options,
		},
	}
	if err := c.nextExpectation(expected); err != nil {
		return err
	}
	if expected.callback != nil {
		return expected.callback(ctx, name, options)
	}
	return expected.wait(ctx)
}
