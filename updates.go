package kivikmock

import (
	"context"

	"github.com/go-kivik/kivik/driver"
)

// Updates is a mocked collection of database updates.
type Updates struct {
	iter
}

type driverDBUpdates struct {
	context.Context
	*Updates
}

var _ driver.DBUpdates = &driverDBUpdates{}

func (u *driverDBUpdates) Next(update *driver.DBUpdate) error {
	result, err := u.unshift(u.Context)
	if err != nil {
		return err
	}
	*update = *result.item.(*driver.DBUpdate)
	return nil
}

// CloseError sets an error to be returned when the updates iterator is closed.
func (u *Updates) CloseError(err error) *Updates {
	u.closeErr = err
	return u
}
