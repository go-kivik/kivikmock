package kivikmock

import (
	"fmt"
	"reflect"
	"time"
)

// ExpectedDBClose is used to manage *kivik.Client.Close expectation returned
// by Mock.ExpectClose.
type ExpectedDBClose struct {
	commonExpectation
}

func (e *ExpectedDBClose) method(v bool) string {
	if v {
		return "DB.Close(ctx)"
	}
	return "DB.Close()"
}

func (e *ExpectedDBClose) met(_ expectation) bool { return true }

// WillReturnError allows setting an error for *kivik.Client.Close action.
func (e *ExpectedDBClose) WillReturnError(err error) *ExpectedDBClose {
	e.err = err
	return e
}

func (e *ExpectedDBClose) String() string {
	extra := delayString(e.delay) + errorString(e.err)
	msg := "call to DB.Close()"
	if extra != "" {
		msg += " which:" + extra
	}
	return msg
}

// WillDelay will cause execution of Close() to delay by duration d.
func (e *ExpectedDBClose) WillDelay(d time.Duration) *ExpectedDBClose {
	e.delay = d
	return e
}

// ExpectedAllDocs represents an expectation to call DB.AllDocs().
type ExpectedAllDocs struct {
	commonExpectation
	options map[string]interface{}
	rows    *Rows
}

func (e *ExpectedAllDocs) String() string {
	msg := "call to DB.AllDocs() which:"
	msg += optionsString(e.options)
	var count int
	if e.rows != nil {
		for _, r := range e.rows.results {
			if r != nil && r.Row != nil {
				count++
			}
		}
	}
	msg += fmt.Sprintf("\n\t- should return: %d results", count)
	msg += delayString(e.delay)
	msg += errorString(e.err)
	return msg
}

func (e *ExpectedAllDocs) method(v bool) string {
	if !v {
		return "DB.AllDocs()"
	}
	if e.options == nil {
		return "DB.AllDocs(ctx)"
	}
	return fmt.Sprintf("DB.AllDocs(ctx, %v)", e.options)
}

func (e *ExpectedAllDocs) met(ex expectation) bool {
	exp := ex.(*ExpectedAllDocs)
	return reflect.DeepEqual(e.options, exp.options)
}

// WithOptions sets the expected options for the AllDocs() call.
func (e *ExpectedAllDocs) WithOptions(options map[string]interface{}) *ExpectedAllDocs {
	e.options = options
	return e
}

// WillReturnRows sets rows to be returned by AllDocs().
func (e *ExpectedAllDocs) WillReturnRows(rows *Rows) *ExpectedAllDocs {
	e.rows = rows
	return e
}

// WillReturnError sets the error that will be returned by AllDocs().
func (e *ExpectedAllDocs) WillReturnError(err error) *ExpectedAllDocs {
	e.err = err
	return e
}

// WillDelay causes AllDocs() to delay execution by the specified duration.
func (e *ExpectedAllDocs) WillDelay(delay time.Duration) *ExpectedAllDocs {
	e.delay = delay
	return e
}
