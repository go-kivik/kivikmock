package kivikmock

import (
	"context"
	"io"
	"time"

	"github.com/go-kivik/kivik/driver"
)

type delayedChange struct {
	delay time.Duration
	*driver.Change
}

// Changes is a mocked collection of Changes results.
type Changes struct {
	closeErr  error
	results   []*delayedChange
	resultErr error
}

type driverChanges struct {
	context.Context
	*Changes
}

var _ driver.Changes = &driverChanges{}

func (r *driverChanges) Close() error {
	return r.closeErr
}

func (r *driverChanges) Next(res *driver.Change) error {
	if len(r.results) == 0 {
		if r.resultErr != nil {
			return r.resultErr
		}
		return io.EOF
	}
	var result *delayedChange
	result, r.results = r.results[0], r.results[1:]
	if result.delay > 0 {
		if err := pause(r.Context, result.delay); err != nil {
			return err
		}
		return r.Next(res)
	}
	*res = *result.Change
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
	r.results = append(r.results, &delayedChange{Change: change})
	return r
}

// AddChangeError adds an error to be returned during iteration.
func (r *Changes) AddChangeError(err error) *Changes {
	r.resultErr = err
	return r
}

// AddDelay adds a delay before the next iteration will complete.
func (r *Changes) AddDelay(delay time.Duration) *Changes {
	r.results = append(r.results, &delayedChange{delay: delay})
	return r
}

// rowCount calculates the rows remaining in this iterator
func (r *Changes) rowCount() int {
	if r == nil || r.results == nil {
		return 0
	}
	var count int
	for _, result := range r.results {
		if result != nil && result.Change != nil {
			count++
		}
	}

	return count
}
