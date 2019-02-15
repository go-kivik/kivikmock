package kivikmock

import (
	"context"
	"io"

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
	if u.count() == 0 {
		if u.resultErr != nil {
			return u.resultErr
		}
		return io.EOF
	}
	result := u.unshift()
	if result.delay > 0 {
		if err := pause(u.Context, result.delay); err != nil {
			return err
		}
		return u.Next(update)
	}
	*update = *result.item.(*driver.DBUpdate)
	return nil
}

// CloseError sets an error to be returned when the updates iterator is closed.
func (u *Updates) CloseError(err error) *Updates {
	u.closeErr = err
	return u
}
