/* This file is auto-generated. Do not edit it! */

package kivikmock

import "time"

// ExpectedCompact represents an expectation for a call to DB.Compact().
type ExpectedCompact struct {
	commonExpectation
	db *MockDB
}

// WillReturnError sets the error value that will be returned by the call to DB.Compact().
func (e *ExpectedCompact) WillReturnError(err error) *ExpectedCompact {
	e.err = err
	return e
}

// WillDelay causes the call to DB.Compact() to delay.
func (e *ExpectedCompact) WillDelay(delay time.Duration) *ExpectedCompact {
	e.delay = delay
	return e
}

// ExpectedCompactView represents an expectation for a call to DB.CompactView().
type ExpectedCompactView struct {
	commonExpectation
	db   *MockDB
	arg0 string
}

// WillReturnError sets the error value that will be returned by the call to DB.CompactView().
func (e *ExpectedCompactView) WillReturnError(err error) *ExpectedCompactView {
	e.err = err
	return e
}

// WillDelay causes the call to DB.CompactView() to delay.
func (e *ExpectedCompactView) WillDelay(delay time.Duration) *ExpectedCompactView {
	e.delay = delay
	return e
}

// ExpectedCopy represents an expectation for a call to DB.Copy().
type ExpectedCopy struct {
	commonExpectation
	db      *MockDB
	arg0    string
	arg1    string
	options map[string]interface{}
	ret0    string
}

// WithOptions sets the expected options for the call to DB.Copy().
func (e *ExpectedCopy) WithOptions(options map[string]interface{}) *ExpectedCopy {
	e.options = options
	return e
}

// WillReturn sets the values that will be returned by the call to DB.Copy().
func (e *ExpectedCopy) WillReturn(ret0 string) *ExpectedCopy {
	e.ret0 = ret0
	return e
}

// WillReturnError sets the error value that will be returned by the call to DB.Copy().
func (e *ExpectedCopy) WillReturnError(err error) *ExpectedCopy {
	e.err = err
	return e
}

// WillDelay causes the call to DB.Copy() to delay.
func (e *ExpectedCopy) WillDelay(delay time.Duration) *ExpectedCopy {
	e.delay = delay
	return e
}

// ExpectedCreateDoc represents an expectation for a call to DB.CreateDoc().
type ExpectedCreateDoc struct {
	commonExpectation
	db      *MockDB
	arg0    interface{}
	options map[string]interface{}
	ret0    string
	ret1    string
}

// WithOptions sets the expected options for the call to DB.CreateDoc().
func (e *ExpectedCreateDoc) WithOptions(options map[string]interface{}) *ExpectedCreateDoc {
	e.options = options
	return e
}

// WillReturn sets the values that will be returned by the call to DB.CreateDoc().
func (e *ExpectedCreateDoc) WillReturn(ret0 string, ret1 string) *ExpectedCreateDoc {
	e.ret0 = ret0
	e.ret1 = ret1
	return e
}

// WillReturnError sets the error value that will be returned by the call to DB.CreateDoc().
func (e *ExpectedCreateDoc) WillReturnError(err error) *ExpectedCreateDoc {
	e.err = err
	return e
}

// WillDelay causes the call to DB.CreateDoc() to delay.
func (e *ExpectedCreateDoc) WillDelay(delay time.Duration) *ExpectedCreateDoc {
	e.delay = delay
	return e
}

// ExpectedCreateIndex represents an expectation for a call to DB.CreateIndex().
type ExpectedCreateIndex struct {
	commonExpectation
	db   *MockDB
	arg0 string
	arg1 string
	arg2 interface{}
}

// WillReturnError sets the error value that will be returned by the call to DB.CreateIndex().
func (e *ExpectedCreateIndex) WillReturnError(err error) *ExpectedCreateIndex {
	e.err = err
	return e
}

// WillDelay causes the call to DB.CreateIndex() to delay.
func (e *ExpectedCreateIndex) WillDelay(delay time.Duration) *ExpectedCreateIndex {
	e.delay = delay
	return e
}

// ExpectedDelete represents an expectation for a call to DB.Delete().
type ExpectedDelete struct {
	commonExpectation
	db      *MockDB
	arg0    string
	arg1    string
	options map[string]interface{}
	ret0    string
}

// WithOptions sets the expected options for the call to DB.Delete().
func (e *ExpectedDelete) WithOptions(options map[string]interface{}) *ExpectedDelete {
	e.options = options
	return e
}

// WillReturn sets the values that will be returned by the call to DB.Delete().
func (e *ExpectedDelete) WillReturn(ret0 string) *ExpectedDelete {
	e.ret0 = ret0
	return e
}

// WillReturnError sets the error value that will be returned by the call to DB.Delete().
func (e *ExpectedDelete) WillReturnError(err error) *ExpectedDelete {
	e.err = err
	return e
}

// WillDelay causes the call to DB.Delete() to delay.
func (e *ExpectedDelete) WillDelay(delay time.Duration) *ExpectedDelete {
	e.delay = delay
	return e
}

// ExpectedDeleteAttachment represents an expectation for a call to DB.DeleteAttachment().
type ExpectedDeleteAttachment struct {
	commonExpectation
	db      *MockDB
	arg0    string
	arg1    string
	arg2    string
	options map[string]interface{}
	ret0    string
}

// WithOptions sets the expected options for the call to DB.DeleteAttachment().
func (e *ExpectedDeleteAttachment) WithOptions(options map[string]interface{}) *ExpectedDeleteAttachment {
	e.options = options
	return e
}

// WillReturn sets the values that will be returned by the call to DB.DeleteAttachment().
func (e *ExpectedDeleteAttachment) WillReturn(ret0 string) *ExpectedDeleteAttachment {
	e.ret0 = ret0
	return e
}

// WillReturnError sets the error value that will be returned by the call to DB.DeleteAttachment().
func (e *ExpectedDeleteAttachment) WillReturnError(err error) *ExpectedDeleteAttachment {
	e.err = err
	return e
}

// WillDelay causes the call to DB.DeleteAttachment() to delay.
func (e *ExpectedDeleteAttachment) WillDelay(delay time.Duration) *ExpectedDeleteAttachment {
	e.delay = delay
	return e
}

// ExpectedDeleteIndex represents an expectation for a call to DB.DeleteIndex().
type ExpectedDeleteIndex struct {
	commonExpectation
	db   *MockDB
	arg0 string
	arg1 string
}

// WillReturnError sets the error value that will be returned by the call to DB.DeleteIndex().
func (e *ExpectedDeleteIndex) WillReturnError(err error) *ExpectedDeleteIndex {
	e.err = err
	return e
}

// WillDelay causes the call to DB.DeleteIndex() to delay.
func (e *ExpectedDeleteIndex) WillDelay(delay time.Duration) *ExpectedDeleteIndex {
	e.delay = delay
	return e
}

// ExpectedFlush represents an expectation for a call to DB.Flush().
type ExpectedFlush struct {
	commonExpectation
	db *MockDB
}

// WillReturnError sets the error value that will be returned by the call to DB.Flush().
func (e *ExpectedFlush) WillReturnError(err error) *ExpectedFlush {
	e.err = err
	return e
}

// WillDelay causes the call to DB.Flush() to delay.
func (e *ExpectedFlush) WillDelay(delay time.Duration) *ExpectedFlush {
	e.delay = delay
	return e
}

// ExpectedGetMeta represents an expectation for a call to DB.GetMeta().
type ExpectedGetMeta struct {
	commonExpectation
	db      *MockDB
	arg0    string
	options map[string]interface{}
	ret0    int64
	ret1    string
}

// WithOptions sets the expected options for the call to DB.GetMeta().
func (e *ExpectedGetMeta) WithOptions(options map[string]interface{}) *ExpectedGetMeta {
	e.options = options
	return e
}

// WillReturn sets the values that will be returned by the call to DB.GetMeta().
func (e *ExpectedGetMeta) WillReturn(ret0 int64, ret1 string) *ExpectedGetMeta {
	e.ret0 = ret0
	e.ret1 = ret1
	return e
}

// WillReturnError sets the error value that will be returned by the call to DB.GetMeta().
func (e *ExpectedGetMeta) WillReturnError(err error) *ExpectedGetMeta {
	e.err = err
	return e
}

// WillDelay causes the call to DB.GetMeta() to delay.
func (e *ExpectedGetMeta) WillDelay(delay time.Duration) *ExpectedGetMeta {
	e.delay = delay
	return e
}

// ExpectedPut represents an expectation for a call to DB.Put().
type ExpectedPut struct {
	commonExpectation
	db      *MockDB
	arg0    string
	arg1    interface{}
	options map[string]interface{}
	ret0    string
}

// WithOptions sets the expected options for the call to DB.Put().
func (e *ExpectedPut) WithOptions(options map[string]interface{}) *ExpectedPut {
	e.options = options
	return e
}

// WillReturn sets the values that will be returned by the call to DB.Put().
func (e *ExpectedPut) WillReturn(ret0 string) *ExpectedPut {
	e.ret0 = ret0
	return e
}

// WillReturnError sets the error value that will be returned by the call to DB.Put().
func (e *ExpectedPut) WillReturnError(err error) *ExpectedPut {
	e.err = err
	return e
}

// WillDelay causes the call to DB.Put() to delay.
func (e *ExpectedPut) WillDelay(delay time.Duration) *ExpectedPut {
	e.delay = delay
	return e
}

// ExpectedViewCleanup represents an expectation for a call to DB.ViewCleanup().
type ExpectedViewCleanup struct {
	commonExpectation
	db *MockDB
}

// WillReturnError sets the error value that will be returned by the call to DB.ViewCleanup().
func (e *ExpectedViewCleanup) WillReturnError(err error) *ExpectedViewCleanup {
	e.err = err
	return e
}

// WillDelay causes the call to DB.ViewCleanup() to delay.
func (e *ExpectedViewCleanup) WillDelay(delay time.Duration) *ExpectedViewCleanup {
	e.delay = delay
	return e
}
