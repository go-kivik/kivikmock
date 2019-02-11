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
	db *MockDB
}

func (e *ExpectedDBClose) method(v bool) string {
	if v {
		return "DB.Close(ctx)"
	}
	return "DB.Close()"
}

func (e *ExpectedDBClose) met(ex expectation) bool {
	exp := ex.(*ExpectedDBClose)
	return e.db.name == exp.db.name
}

// WillReturnError allows setting an error for *kivik.Client.Close action.
func (e *ExpectedDBClose) WillReturnError(err error) *ExpectedDBClose {
	e.err = err
	return e
}

func (e *ExpectedDBClose) String() string {
	msg := fmt.Sprintf("call to DB(%s#%d).Close()", e.db.name, e.db.id)
	extra := delayString(e.delay)
	extra += errorString(e.err)
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
	db      *MockDB
	options map[string]interface{}
	rows    *Rows
}

func (e *ExpectedAllDocs) String() string {
	msg := fmt.Sprintf("call to DB(%s#%d).AllDocs() which:", e.db.name, e.db.id)
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
		return fmt.Sprintf("DB(%s).AllDocs(ctx)", e.db.name)
	}
	return fmt.Sprintf("DB(%s).AllDocs(ctx, %v)", e.db.name, e.options)
}

func (e *ExpectedAllDocs) met(ex expectation) bool {
	exp := ex.(*ExpectedAllDocs)
	if e.db.name != exp.db.name {
		return false
	}
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
	db      *MockDB
	docs    []driver.BulkGetReference
	options map[string]interface{}
	rows    *Rows
}

func (e *ExpectedBulkGet) String() string {
	msg := fmt.Sprintf("call to DB(%s#%d).BulkGet() which:", e.db.name, e.db.id)
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
	return fmt.Sprintf("DB(%s).BulkGet(ctx, %s%s)", e.db.name, docs, options)
}

func (e *ExpectedBulkGet) met(ex expectation) bool {
	exp := ex.(*ExpectedBulkGet)
	if e.db.name != exp.db.name {
		return false
	}
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
	db    *MockDB
	query interface{}
	rows  *Rows
}

func (e *ExpectedFind) String() string {
	msg := fmt.Sprintf("call to DB(%s#%d).Find() which:", e.db.name, e.db.id)
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
		return fmt.Sprintf("DB(%s).Find(ctx, ?)", e.db.name)
	}
	return fmt.Sprintf("DB(%s).Find(ctx, %v)", e.db.name, e.query)
}

func (e *ExpectedFind) met(ex expectation) bool {
	exp := ex.(*ExpectedFind)
	if e.db.name != exp.db.name {
		return false
	}
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

func (e *ExpectedCreateIndex) String() string {
	msg := fmt.Sprintf("call to DB(%s#%d).CreateIndex() which:", e.db.name, e.db.id)
	if e.arg0 == "" {
		msg += "\n\t- has any ddoc"
	} else {
		msg += "\n\t- has ddoc: " + e.arg0
	}
	if e.arg1 == "" {
		msg += "\n\t- has any name"
	} else {
		msg += "\n\t- has name: " + e.arg1
	}
	if e.arg2 == nil {
		msg += "\n\t- has any index"
	} else {
		msg += fmt.Sprintf("\n\t- has index: %v", e.arg2)
	}
	msg += delayString(e.delay)
	msg += errorString(e.err)
	return msg
}

func (e *ExpectedCreateIndex) method(v bool) string {
	if !v {
		return "DB.CreateIndex()"
	}
	var ddoc, name, index string
	if e.arg0 == "" {
		ddoc = "?"
	} else {
		ddoc = fmt.Sprintf("%q", e.arg0)
	}
	if e.arg1 == "" {
		name = "?"
	} else {
		name = fmt.Sprintf("%q", e.arg1)
	}
	if e.arg2 == nil {
		index = "?"
	} else {
		index = fmt.Sprintf("%v", e.arg2)
	}
	return fmt.Sprintf("DB(%s).CreateIndex(ctx, %s, %s, %s)", e.db.name, ddoc, name, index)
}

func (e *ExpectedCreateIndex) met(ex expectation) bool {
	exp := ex.(*ExpectedCreateIndex)
	if e.db.name != exp.db.name {
		return false
	}
	if exp.arg0 != "" && exp.arg0 != e.arg0 {
		return false
	}
	if exp.arg1 != "" && exp.arg1 != e.arg1 {
		return false
	}
	return exp.arg2 == nil || diff.AsJSON(exp.arg2, e.arg2) == nil
}

// WithDDoc sets the expected ddoc value for the DB.CreateIndex() call.
func (e *ExpectedCreateIndex) WithDDoc(ddoc string) *ExpectedCreateIndex {
	e.arg0 = ddoc
	return e
}

// WithName sets the expected name value for the DB.CreateIndex() call.
func (e *ExpectedCreateIndex) WithName(name string) *ExpectedCreateIndex {
	e.arg1 = name
	return e
}

// WithIndex sets the expected index value for the DB.CreateIndex() call.
func (e *ExpectedCreateIndex) WithIndex(index interface{}) *ExpectedCreateIndex {
	e.arg2 = index
	return e
}

// ExpectedGetIndexes represents an expectation to call GetIndexes().
type ExpectedGetIndexes struct {
	commonExpectation
	db      *MockDB
	indexes []kivik.Index
}

func (e *ExpectedGetIndexes) String() string {
	msg := fmt.Sprintf("call to DB(%s#%d).GetIndexes()", e.db.name, e.db.id)
	var extra string
	if e.indexes != nil {
		extra += fmt.Sprintf("\n\t- should return indexes: %v", e.indexes)
	}
	extra += delayString(e.delay)
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
	return fmt.Sprintf("DB(%s).GetIndexes(ctx)", e.db.name)
}

func (e *ExpectedGetIndexes) met(ex expectation) bool {
	exp := ex.(*ExpectedGetIndexes)
	return e.db.name == exp.db.name
}

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

func (e *ExpectedDeleteIndex) String() string {
	msg := fmt.Sprintf("call to DB(%s#%d).DeleteIndex() which:", e.db.name, e.db.id)
	if e.arg0 == "" {
		msg += "\n\t- has any ddoc"
	} else {
		msg += "\n\t- has ddoc: " + e.arg0
	}
	msg += nameString(e.arg1)
	msg += delayString(e.delay)
	msg += errorString(e.err)
	return msg
}

func (e *ExpectedDeleteIndex) method(v bool) string {
	if !v {
		return "DB.DeleteIndex()"
	}
	var ddoc, name string
	if e.arg0 == "" {
		ddoc = "?"
	} else {
		ddoc = fmt.Sprintf("%q", e.arg0)
	}
	if e.arg1 == "" {
		name = "?"
	} else {
		name = fmt.Sprintf("%q", e.arg1)
	}
	return fmt.Sprintf("DB(%s).DeleteIndex(ctx, %s, %s)", e.db.name, ddoc, name)
}

func (e *ExpectedDeleteIndex) met(ex expectation) bool {
	exp := ex.(*ExpectedDeleteIndex)
	if e.db.name != exp.db.name {
		return false
	}
	if exp.arg0 != "" && exp.arg0 != e.arg0 {
		return false
	}
	if exp.arg1 != "" && exp.arg1 != e.arg1 {
		return false
	}
	return true
}

// WithDDoc sets the expected ddoc to be passed to the DB.DeleteIndex() call.
func (e *ExpectedDeleteIndex) WithDDoc(ddoc string) *ExpectedDeleteIndex {
	e.arg0 = ddoc
	return e
}

// WithName sets the expected name to be passed to the DB.DeleteIndex() call.
func (e *ExpectedDeleteIndex) WithName(name string) *ExpectedDeleteIndex {
	e.arg1 = name
	return e
}

// ExpectedExplain represents an expectation for a DB.Explain() call.
type ExpectedExplain struct {
	commonExpectation
	db    *MockDB
	query interface{}
	plan  *kivik.QueryPlan
}

func (e *ExpectedExplain) String() string {
	msg := fmt.Sprintf("call to DB(%s#%d).Explain() which:", e.db.name, e.db.id)
	if e.query == nil {
		msg += "\n\t- has any query"
	} else {
		msg += fmt.Sprintf("\n\t- has query: %v", e.query)
	}
	if e.plan != nil {
		msg += fmt.Sprintf("\n\t- should return query plan: %v", e.plan)
	}
	msg += delayString(e.delay)
	msg += errorString(e.err)
	return msg
}

func (e *ExpectedExplain) method(v bool) string {
	if !v {
		return "DB.Explain()"
	}
	if e.query != nil {
		return fmt.Sprintf("DB(%s).Explain(ctx, %v)", e.db.name, e.query)
	}
	return fmt.Sprintf("DB(%s).Explain(ctx, ?)", e.db.name)
}

func (e *ExpectedExplain) met(ex expectation) bool {
	exp := ex.(*ExpectedExplain)
	if e.db.name != exp.db.name {
		return false
	}
	return e.plan == nil || diff.AsJSON(e.plan, exp.plan) == nil
}

// WithQuery sets the expected query for the Explain() call.
func (e *ExpectedExplain) WithQuery(query interface{}) *ExpectedExplain {
	e.query = query
	return e
}

// WillReturn sets the query plan to be returned by the Explain() call.
func (e *ExpectedExplain) WillReturn(plan *kivik.QueryPlan) *ExpectedExplain {
	e.plan = plan
	return e
}

// WillReturnError sets the error to be returned by the Explain() call.
func (e *ExpectedExplain) WillReturnError(err error) *ExpectedExplain {
	e.err = err
	return e
}

// WillDelay causes the Explain() call to delay.
func (e *ExpectedExplain) WillDelay(delay time.Duration) *ExpectedExplain {
	e.delay = delay
	return e
}

func (e *ExpectedCreateDoc) String() string {
	msg := fmt.Sprintf("call to DB(%s#%d).CreateDoc() which:", e.db.name, e.db.id)
	if e.arg0 == nil {
		msg += "\n\t- has any doc"
	} else {
		msg += fmt.Sprintf("\n\t- has doc: %v", e.arg0)
	}
	msg += optionsString(e.options)
	if e.ret0 != "" {
		msg += "\n\t- should return docID: " + e.ret0
	}
	if e.ret1 != "" {
		msg += "\n\t- should return rev: " + e.ret1
	}
	msg += delayString(e.delay)
	msg += errorString(e.err)
	return msg
}

func (e *ExpectedCreateDoc) method(v bool) string {
	if !v {
		return "DB.CreateDoc()"
	}
	var doc, options string
	if e.arg0 == nil {
		doc = "?"
	} else {
		doc = fmt.Sprintf("%v", e.arg0)
	}
	if e.options != nil {
		options = fmt.Sprintf(", %v", e.options)
	}
	return fmt.Sprintf("DB(%s).CreateDoc(ctx, %s%s)", e.db.name, doc, options)
}

func (e *ExpectedCreateDoc) met(ex expectation) bool {
	exp := ex.(*ExpectedCreateDoc)
	if e.db.name != exp.db.name {
		return false
	}
	if exp.arg0 != nil && diff.AsJSON(e.arg0, exp.arg0) != nil {
		return false
	}
	return exp.options == nil || reflect.DeepEqual(e.options, exp.options)
}

// WithDoc sets the expected doc for the call to CreateDoc().
func (e *ExpectedCreateDoc) WithDoc(doc interface{}) *ExpectedCreateDoc {
	e.arg0 = doc
	return e
}

func (e *ExpectedCompact) String() string                   { return "" }
func (e *ExpectedCompact) method(v bool) string             { return "" }
func (e *ExpectedCompact) met(ex expectation) bool          { return false }
func (e *ExpectedViewCleanup) String() string               { return "" }
func (e *ExpectedViewCleanup) method(v bool) string         { return "" }
func (e *ExpectedViewCleanup) met(ex expectation) bool      { return false }
func (e *ExpectedPut) String() string                       { return "" }
func (e *ExpectedPut) method(v bool) string                 { return "" }
func (e *ExpectedPut) met(ex expectation) bool              { return false }
func (e *ExpectedGetMeta) String() string                   { return "" }
func (e *ExpectedGetMeta) method(v bool) string             { return "" }
func (e *ExpectedGetMeta) met(ex expectation) bool          { return false }
func (e *ExpectedFlush) String() string                     { return "" }
func (e *ExpectedFlush) method(v bool) string               { return "" }
func (e *ExpectedFlush) met(ex expectation) bool            { return false }
func (e *ExpectedDeleteAttachment) String() string          { return "" }
func (e *ExpectedDeleteAttachment) method(v bool) string    { return "" }
func (e *ExpectedDeleteAttachment) met(ex expectation) bool { return false }
func (e *ExpectedDelete) String() string                    { return "" }
func (e *ExpectedDelete) method(v bool) string              { return "" }
func (e *ExpectedDelete) met(ex expectation) bool           { return false }
func (e *ExpectedCopy) String() string                      { return "" }
func (e *ExpectedCopy) method(v bool) string                { return "" }
func (e *ExpectedCopy) met(ex expectation) bool             { return false }
func (e *ExpectedCompactView) String() string               { return "" }
func (e *ExpectedCompactView) method(v bool) string         { return "" }
func (e *ExpectedCompactView) met(ex expectation) bool      { return false }
