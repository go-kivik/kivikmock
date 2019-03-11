package kivikmock

import (
	"context"
	"time"

	"github.com/go-kivik/kivik/driver"
)

// Changes is a mocked collection of Changes results.
type Changes struct {
	iter
}

type driverChanges struct {
	context.Context
	*Changes
}

var _ driver.Changes = &driverChanges{}

func (r *driverChanges) Next(res *driver.Change) error {
	result, err := r.unshift(r.Context)
	if err != nil {
		return err
	}
	*res = *result.(*driver.Change)
	return nil
}

// CloseError sets an error to be returned when the iterator is closed.
func (r *Changes) CloseError(err error) *Changes {
	r.closeErr = err
	return r
}

// AddChange adds a change result to be returned by the iterator. If
// AddResultError has been set, this method will panic.
func (r *Changes) AddChange(change *driver.Change) *Changes {
	if r.resultErr != nil {
		panic("It is invalid to set more changes after AddChangeError is defined.")
	}
	r.push(&item{item: change})
	return r
}

// AddChangeError adds an error to be returned during iteration.
func (r *Changes) AddChangeError(err error) *Changes {
	r.resultErr = err
	return r
}

// AddDelay adds a delay before the next iteration will complete.
func (r *Changes) AddDelay(delay time.Duration) *Changes {
	r.push(&item{delay: delay})
	return r
}

// Final converts the Changes object to a driver.Changes. This method is
// intended for use within WillExecute() to return results.
func (r *Changes) Final() driver.Changes {
	return &driverChanges{Changes: r}
}
