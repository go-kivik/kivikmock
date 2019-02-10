package kivikmock

import (
	"fmt"
	"reflect"
	"time"

	"github.com/flimzy/diff"
	"github.com/go-kivik/kivik"
	"github.com/go-kivik/kivik/driver"
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
	msg += fmt.Sprintf("\n\t- should return: %d results", e.rows.rowCount())
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

// ExpectedBulkGet represents an expectation to call DB.BulkGet().
type ExpectedBulkGet struct {
	commonExpectation
	docs    []driver.BulkGetReference
	options map[string]interface{}
	rows    *Rows
}

func (e *ExpectedBulkGet) String() string {
	msg := "call to DB.BulkGet() which:"
	if e.docs == nil {
		msg += "\n\t- has any doc references"
	} else {
		msg += fmt.Sprintf("\n\t- has doc references: %v", e.docs)
	}
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

func (e *ExpectedBulkGet) method(v bool) string {
	if !v {
		return "DB.BulkGet()"
	}
	var docs, options string
	if e.docs == nil {
		docs = "?"
	} else {
		docs = fmt.Sprintf("%v", e.docs)
	}
	if e.options == nil {
		options = ""
	} else {
		options = fmt.Sprintf(", %v", e.options)
	}
	return fmt.Sprintf("DB.BulkGet(ctx, %s%s)", docs, options)
}

func (e *ExpectedBulkGet) met(ex expectation) bool {
	exp := ex.(*ExpectedBulkGet)
	return reflect.DeepEqual(e.options, exp.options)
}

// WithOptions sets the expected options for the BulkGet() call.
func (e *ExpectedBulkGet) WithOptions(options map[string]interface{}) *ExpectedBulkGet {
	e.options = options
	return e
}

// WillReturnRows sets rows to be returned by BulkGet().
func (e *ExpectedBulkGet) WillReturnRows(rows *Rows) *ExpectedBulkGet {
	e.rows = rows
	return e
}

// WillReturnError sets the error that will be returned by BulkGet().
func (e *ExpectedBulkGet) WillReturnError(err error) *ExpectedBulkGet {
	e.err = err
	return e
}

// WillDelay causes BulkGet() to delay execution by the specified duration.
func (e *ExpectedBulkGet) WillDelay(delay time.Duration) *ExpectedBulkGet {
	e.delay = delay
	return e
}

// ExpectedFind represents an expectation to call DB.Find().
type ExpectedFind struct {
	commonExpectation
	query interface{}
	rows  *Rows
}

func (e *ExpectedFind) String() string {
	msg := "call to DB.Find() which:"
	if e.query == nil {
		msg += "\n\t- has any query"
	} else {
		msg += fmt.Sprintf("\n\t- has query: %v", e.query)
	}
	msg += fmt.Sprintf("\n\t- should return: %d results", e.rows.rowCount())
	msg += delayString(e.delay)
	msg += errorString(e.err)
	return msg
}

func (e *ExpectedFind) method(v bool) string {
	if !v {
		return "DB.Find()"
	}
	if e.query == nil {
		return "DB.Find(ctx, ?)"
	}
	return fmt.Sprintf("DB.Find(ctx, %v)", e.query)
}

func (e *ExpectedFind) met(ex expectation) bool {
	exp := ex.(*ExpectedFind)
	return exp.query == nil || diff.AsJSON(e.query, exp.query) == nil
}

// WithQuery sets the expected query for the Find() call.
func (e *ExpectedFind) WithQuery(query interface{}) *ExpectedFind {
	e.query = query
	return e
}

// WillReturnRows sets rows to be returned by Find().
func (e *ExpectedFind) WillReturnRows(rows *Rows) *ExpectedFind {
	e.rows = rows
	return e
}

// WillReturnError sets the error that will be returned by Find().
func (e *ExpectedFind) WillReturnError(err error) *ExpectedFind {
	e.err = err
	return e
}

// WillDelay causes Find() to delay execution by the specified duration.
func (e *ExpectedFind) WillDelay(delay time.Duration) *ExpectedFind {
	e.delay = delay
	return e
}

// ExpectedCreateIndex represents an expectation to call DB.CreateIndex().
type ExpectedCreateIndex struct {
	commonExpectation
	ddoc, name string
	index      interface{}
}

func (e *ExpectedCreateIndex) String() string {
	msg := "call to DB.CreateIndex() which:"
	if e.ddoc == "" {
		msg += "\n\t- has any ddoc"
	} else {
		msg += "\n\t- has ddoc: " + e.ddoc
	}
	if e.name == "" {
		msg += "\n\t- has any name"
	} else {
		msg += "\n\t- has name: " + e.name
	}
	if e.index == nil {
		msg += "\n\t- has any index"
	} else {
		msg += fmt.Sprintf("\n\t- has index: %v", e.index)
	}
	msg += errorString(e.err)
	return msg
}

func (e *ExpectedCreateIndex) method(v bool) string {
	if !v {
		return "DB.CreateIndex()"
	}
	var ddoc, name, index string
	if e.ddoc == "" {
		ddoc = "?"
	} else {
		ddoc = fmt.Sprintf("%q", e.ddoc)
	}
	if e.name == "" {
		name = "?"
	} else {
		name = fmt.Sprintf("%q", e.name)
	}
	if e.index == nil {
		index = "?"
	} else {
		index = fmt.Sprintf("%v", e.index)
	}
	return fmt.Sprintf("DB.CreateIndex(ctx, %s, %s, %s)", ddoc, name, index)
}

func (e *ExpectedCreateIndex) met(ex expectation) bool {
	exp := ex.(*ExpectedCreateIndex)
	if exp.ddoc != "" && exp.ddoc != e.ddoc {
		return false
	}
	if exp.name != "" && exp.name != e.name {
		return false
	}
	return exp.index == nil || diff.AsJSON(exp.index, e.index) == nil
}

// WithDDoc sets the expected ddoc value for the DB.CreateIndex() call.
func (e *ExpectedCreateIndex) WithDDoc(ddoc string) *ExpectedCreateIndex {
	e.ddoc = ddoc
	return e
}

// WithName sets the expected name value for the DB.CreateIndex() call.
func (e *ExpectedCreateIndex) WithName(name string) *ExpectedCreateIndex {
	e.name = name
	return e
}

// WithIndex sets the expected index value for the DB.CreateIndex() call.
func (e *ExpectedCreateIndex) WithIndex(index interface{}) *ExpectedCreateIndex {
	e.index = index
	return e
}

// WillReturnError sets the error to be returned by the DB.CreateIndex() call.
func (e *ExpectedCreateIndex) WillReturnError(err error) *ExpectedCreateIndex {
	e.err = err
	return e
}

// WillDelay causes the DB.CreateIndex() call to delay.
func (e *ExpectedCreateIndex) WillDelay(delay time.Duration) *ExpectedCreateIndex {
	e.delay = delay
	return e
}

// ExpectedGetIndexes represents an expectation to call GetIndexes().
type ExpectedGetIndexes struct {
	commonExpectation
	indexes []kivik.Index
}

func (e *ExpectedGetIndexes) String() string {
	msg := "call to DB.GetIndexes()"
	var extra string
	if e.indexes != nil {
		extra += fmt.Sprintf("\n\t- should return indexes: %v", e.indexes)
	}
	extra += errorString(e.err)
	if extra != "" {
		msg += " which:" + extra
	}
	return msg
}

func (e *ExpectedGetIndexes) method(v bool) string {
	if !v {
		return "DB.GetIndexes()"
	}
	return "DB.GetIndexes(ctx)"
}

func (e *ExpectedGetIndexes) met(_ expectation) bool { return true }

// WillReturn sets the indexes that will be returned by the call to
// DB.GetIndexes().
func (e *ExpectedGetIndexes) WillReturn(indexes []kivik.Index) *ExpectedGetIndexes {
	e.indexes = indexes
	return e
}

// WillReturnError sets the error that will be returned by the call to
// DB.GetIndexes().
func (e *ExpectedGetIndexes) WillReturnError(err error) *ExpectedGetIndexes {
	e.err = err
	return e
}

// WillDelay causes the call to DB.GetIndexes() to delay.
func (e *ExpectedGetIndexes) WillDelay(delay time.Duration) *ExpectedGetIndexes {
	e.delay = delay
	return e
}
