package kivikmock

import (
	"context"
	"io"
	"time"

	"github.com/go-kivik/kivik/driver"
)

type delayedBulkResult struct {
	delay time.Duration
	*driver.BulkResult
}

// BulkResults is a mocked collection of BulkDoc results.
type BulkResults struct {
	closeErr  error
	results   []*delayedBulkResult
	resultErr error
}

type driverBulkResults struct {
	context.Context
	*BulkResults
}

var _ driver.BulkResults = &driverBulkResults{}

func (r *driverBulkResults) Close() error {
	return r.closeErr
}

func (r *driverBulkResults) Next(res *driver.BulkResult) error {
	if len(r.results) == 0 {
		if r.resultErr != nil {
			return r.resultErr
		}
		return io.EOF
	}
	var result *delayedBulkResult
	result, r.results = r.results[0], r.results[1:]
	if result.delay > 0 {
		if err := pause(r.Context, result.delay); err != nil {
			return err
		}
		return r.Next(res)
	}
	*res = *result.BulkResult
	return nil
}

// CloseError sets an error to be returned when the iterator is closed.
func (r *BulkResults) CloseError(err error) *BulkResults {
	r.closeErr = err
	return r
}

// AddResult adds a bulk result to be returned by the iterator. If
// AddResultError has been set, this method will panic.
func (r *BulkResults) AddResult(row *driver.BulkResult) *BulkResults {
	if r.resultErr != nil {
		panic("It is invalid to set more rows after AddRowError is defined.")
	}
	r.results = append(r.results, &delayedBulkResult{BulkResult: row})
	return r
}

// AddResultError adds an error to be returned during iteration.
func (r *BulkResults) AddResultError(err error) *BulkResults {
	r.resultErr = err
	return r
}

// AddDelay adds a delay before the next iteration will complete.
func (r *BulkResults) AddDelay(delay time.Duration) *BulkResults {
	r.results = append(r.results, &delayedBulkResult{delay: delay})
	return r
}

// rowCount calculates the rows remaining in this iterator
func (r *BulkResults) rowCount() int {
	if r == nil || r.results == nil {
		return 0
	}
	var count int
	for _, result := range r.results {
		if result != nil && result.BulkResult != nil {
			count++
		}
	}

	return count
}
