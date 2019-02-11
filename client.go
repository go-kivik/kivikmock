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

	return expected.wait(ctx)
}

func (c *driverClient) DBsStats(ctx context.Context, names []string) ([]*driver.DBStats, error) {
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

func (c *driverClient) Session(ctx context.Context) (*driver.Session, error) {
	expected := &ExpectedSession{}
	if err := c.nextExpectation(expected); err != nil {
		return nil, err
	}
	var s *driver.Session
	if expected.session != nil {
		s = new(driver.Session)
		*s = driver.Session(*expected.session)
	}
	return s, expected.wait(ctx)
}

func (c *driverClient) Version(ctx context.Context) (*driver.Version, error) {
	expected := &ExpectedVersion{}
	if err := c.nextExpectation(expected); err != nil {
		return nil, err
	}
	var v *driver.Version
	if expected.version != nil {
		v = new(driver.Version)
		*v = driver.Version(*expected.version)
	}
	return v, expected.wait(ctx)
}

func (c *driverClient) DB(ctx context.Context, name string, options map[string]interface{}) (driver.DB, error) {
	expected := &ExpectedDB{
		arg0:    name,
		options: options,
	}
	if err := c.nextExpectation(expected); err != nil {
		return nil, err
	}
	expected.db.name = name
	return &driverDB{MockDB: expected.db}, expected.wait(ctx)
}

func (c *driverClient) CreateDB(ctx context.Context, name string, options map[string]interface{}) error {
	expected := &ExpectedCreateDB{
		arg0:    name,
		options: options,
	}
	if err := c.nextExpectation(expected); err != nil {
		return err
	}
	return expected.wait(ctx)
}
