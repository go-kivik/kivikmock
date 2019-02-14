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
