package kivikmock

import (
	"context"
	"io"
	"time"

	"github.com/go-kivik/kivik/driver"
)

type delayedRow struct {
	delay time.Duration
	*driver.Row
}

// Rows is a mocked collection of rows.
type Rows struct {
	closeErr  error
	results   []*delayedRow
	resultErr error
	offset    int64
	updateSeq string
	totalRows int64
	warning   string
}

type driverRows struct {
	context.Context
	*Rows
}

var _ driver.Rows = &driverRows{}
var _ driver.RowsWarner = &driverRows{}

func (r *driverRows) Close() error      { return r.closeErr }
func (r *driverRows) Offset() int64     { return r.offset }
func (r *driverRows) UpdateSeq() string { return r.updateSeq }
func (r *driverRows) TotalRows() int64  { return r.totalRows }
func (r *driverRows) Warning() string   { return r.warning }

func (r *driverRows) Next(row *driver.Row) error {
	if len(r.results) == 0 {
		if r.resultErr != nil {
			return r.resultErr
		}
		return io.EOF
	}
	var result *delayedRow
	result, r.results = r.results[0], r.results[1:]
	if result.delay > 0 {
		if err := pause(r.Context, result.delay); err != nil {
			return err
		}
		return r.Next(row)
	}
	*row = *result.Row
	return nil
}

// CloseError sets an error to be returned when the rows iterator is closed.
func (r *Rows) CloseError(err error) *Rows {
	r.closeErr = err
	return r
}

// Offset sets the offset value to be returned by the rows iterator.
func (r *Rows) Offset(offset int64) *Rows {
	r.offset = offset
	return r
}

// TotalRows sets the total rows value to be returned by the rows iterator.
func (r *Rows) TotalRows(totalRows int64) *Rows {
	r.totalRows = totalRows
	return r
}

// UpdateSeq sets the update sequence value to be returned by the rows iterator.
func (r *Rows) UpdateSeq(seq string) *Rows {
	r.updateSeq = seq
	return r
}

// Warning sets the warning value to be returned by the rows iterator.
func (r *Rows) Warning(warning string) *Rows {
	r.warning = warning
	return r
}

// AddRow adds a row to be returned by the rows iterator. If AddrowError has
// been set, this method will panic.
func (r *Rows) AddRow(row *driver.Row) *Rows {
	if r.resultErr != nil {
		panic("It is invalid to set more rows after AddRowError is defined.")
	}
	r.results = append(r.results, &delayedRow{Row: row})
	return r
}

// AddRowError adds an error to be returned during row iteration.
func (r *Rows) AddRowError(err error) *Rows {
	r.resultErr = err
	return r
}

// AddDelay adds a delay before the next iteration will complete.
func (r *Rows) AddDelay(delay time.Duration) *Rows {
	r.results = append(r.results, &delayedRow{delay: delay})
	return r
}

// rowCount calculates the rows remaining in this iterator
func (r *Rows) rowCount() int {
	if r == nil || r.results == nil {
		return 0
	}
	var count int
	for _, result := range r.results {
		if result != nil && result.Row != nil {
			count++
		}
	}

	return count
}
