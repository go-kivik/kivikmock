package kivikmock

import (
	"context"
	"io"
	"time"

	"github.com/go-kivik/kivik/driver"
)

type delayedUpdate struct {
	delay time.Duration
	*driver.DBUpdate
}

type Updates struct {
	closeErr  error
	results   []*delayedUpdate
	resultErr error
}

type driverDBUpdates struct {
	context.Context
	*Updates
}

var _ driver.DBUpdates = &driverDBUpdates{}

func (u *driverDBUpdates) Close() error { return u.closeErr }

func (u *driverDBUpdates) Next(update *driver.DBUpdate) error {
	if len(u.results) == 0 {
		if u.resultErr != nil {
			return u.resultErr
		}
		return io.EOF
	}
	var result *delayedUpdate
	result, u.results = u.results[0], u.results[1:]
	if result.delay > 0 {
		if err := pause(u.Context, result.delay); err != nil {
			return err
		}
		return u.Next(update)
	}
	*update = *result.DBUpdate
	return nil
}

// CloseError sets an error to be returned when the updates iterator is closed.
func (u *Updates) CloseError(err error) *Updates {
	u.closeErr = err
	return u
}

func (u *Updates) count() int {
	if u == nil || u.results == nil {
		return 0
	}
	var count int
	for _, result := range u.results {
		if result != nil && result.DBUpdate != nil {
			count++
		}
	}

	return count
}
