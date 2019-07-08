package kivikmock

import (
	"context"
	"errors"
	"reflect"
	"time"

	"github.com/go-kivik/kivik/driver"
)

type driverClient struct {
	*Client
}

var _ driver.Client = &driverClient{}
var _ driver.ClientCloser = &driverClient{}
var _ driver.Authenticator = &driverClient{}
var _ driver.Cluster = &driverClient{}
var _ driver.DBsStatser = &driverClient{}
var _ driver.Pinger = &driverClient{}
var _ driver.Sessioner = &driverClient{}
var _ driver.Configer = &driverClient{}

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

type driverReplication struct {
	*Replication
}

var _ driver.Replication = &driverReplication{}

func (r *driverReplication) ReplicationID() string {
	return r.Replication.ID
}

func (r *driverReplication) Source() string {
	return r.Replication.Source
}

func (r *driverReplication) Target() string {
	return r.Replication.Target
}

func (r *driverReplication) StartTime() time.Time {
	return r.Replication.StartTime
}

func (r *driverReplication) EndTime() time.Time {
	return r.Replication.EndTime
}

func (r *driverReplication) State() string {
	return r.Replication.State
}

func (r *driverReplication) Err() error {
	return r.Replication.Err
}

func (r *driverReplication) Delete(_ context.Context) error {
	return errors.New("not implemented")
}

func (r *driverReplication) Update(_ context.Context, _ *driver.ReplicationInfo) error {
	return errors.New("not implemented")
}
