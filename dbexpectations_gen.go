/* This file is auto-generated. Do not edit it! */

package kivikmock

import (
	"fmt"
	"reflect"
	"time"

	"github.com/go-kivik/kivik/driver"
)

var _ = &driver.Attachment{}
var _ = reflect.Int

// ExpectedCompact represents an expectation for a call to DB.Compact().
type ExpectedCompact struct {
	commonExpectation
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

func (e *ExpectedCompact) met(_ expectation) bool {
	return true
}

func (e *ExpectedCompact) method(v bool) string {
	if !v {
		return "DB.Compact()"
	}
	return fmt.Sprintf("DB(%s).Compact(ctx)", e.DB().name)
}

// ExpectedCompactView represents an expectation for a call to DB.CompactView().
type ExpectedCompactView struct {
	commonExpectation
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

func (e *ExpectedCompactView) met(ex expectation) bool {
	exp := ex.(*ExpectedCompactView)
	if exp.arg0 != "" && exp.arg0 != e.arg0 {
		return false
	}
	return true
}

func (e *ExpectedCompactView) method(v bool) string {
	if !v {
		return "DB.CompactView()"
	}
	arg0 := "?"
	if e.arg0 != "" {
		arg0 = fmt.Sprintf("%q", e.arg0)
	}
	return fmt.Sprintf("DB(%s).CompactView(ctx, %s)", e.DB().name, arg0)
}

// ExpectedCopy represents an expectation for a call to DB.Copy().
type ExpectedCopy struct {
	commonExpectation
	arg0 string
	arg1 string
	ret0 string
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

func (e *ExpectedCopy) met(ex expectation) bool {
	exp := ex.(*ExpectedCopy)
	if exp.arg0 != "" && exp.arg0 != e.arg0 {
		return false
	}
	if exp.arg1 != "" && exp.arg1 != e.arg1 {
		return false
	}
	return true
}

func (e *ExpectedCopy) method(v bool) string {
	if !v {
		return "DB.Copy()"
	}
	arg0, arg1, options := "?", "?", "[?]"
	if e.arg0 != "" {
		arg0 = fmt.Sprintf("%q", e.arg0)
	}
	if e.arg1 != "" {
		arg1 = fmt.Sprintf("%q", e.arg1)
	}
	if e.options != nil {
		options = fmt.Sprintf("%v", e.options)
	}
	return fmt.Sprintf("DB(%s).Copy(ctx, %s, %s, %s)", e.DB().name, arg0, arg1, options)
}

// ExpectedCreateDoc represents an expectation for a call to DB.CreateDoc().
type ExpectedCreateDoc struct {
	commonExpectation
	arg0 interface{}
	ret0 string
	ret1 string
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

func (e *ExpectedCreateDoc) met(ex expectation) bool {
	exp := ex.(*ExpectedCreateDoc)
	if exp.arg0 != nil && !jsonMeets(exp.arg0, e.arg0) {
		return false
	}
	return true
}

func (e *ExpectedCreateDoc) method(v bool) string {
	if !v {
		return "DB.CreateDoc()"
	}
	arg0, options := "?", "[?]"
	if e.arg0 != nil {
		arg0 = fmt.Sprintf("%v", e.arg0)
	}
	if e.options != nil {
		options = fmt.Sprintf("%v", e.options)
	}
	return fmt.Sprintf("DB(%s).CreateDoc(ctx, %s, %s)", e.DB().name, arg0, options)
}

// ExpectedCreateIndex represents an expectation for a call to DB.CreateIndex().
type ExpectedCreateIndex struct {
	commonExpectation
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

func (e *ExpectedCreateIndex) met(ex expectation) bool {
	exp := ex.(*ExpectedCreateIndex)
	if exp.arg0 != "" && exp.arg0 != e.arg0 {
		return false
	}
	if exp.arg1 != "" && exp.arg1 != e.arg1 {
		return false
	}
	if exp.arg2 != nil && !jsonMeets(exp.arg2, e.arg2) {
		return false
	}
	return true
}

func (e *ExpectedCreateIndex) method(v bool) string {
	if !v {
		return "DB.CreateIndex()"
	}
	arg0, arg1, arg2 := "?", "?", "?"
	if e.arg0 != "" {
		arg0 = fmt.Sprintf("%q", e.arg0)
	}
	if e.arg1 != "" {
		arg1 = fmt.Sprintf("%q", e.arg1)
	}
	if e.arg2 != nil {
		arg2 = fmt.Sprintf("%v", e.arg2)
	}
	return fmt.Sprintf("DB(%s).CreateIndex(ctx, %s, %s, %s)", e.DB().name, arg0, arg1, arg2)
}

// ExpectedDelete represents an expectation for a call to DB.Delete().
type ExpectedDelete struct {
	commonExpectation
	arg0 string
	arg1 string
	ret0 string
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

func (e *ExpectedDelete) met(ex expectation) bool {
	exp := ex.(*ExpectedDelete)
	if exp.arg0 != "" && exp.arg0 != e.arg0 {
		return false
	}
	if exp.arg1 != "" && exp.arg1 != e.arg1 {
		return false
	}
	return true
}

func (e *ExpectedDelete) method(v bool) string {
	if !v {
		return "DB.Delete()"
	}
	arg0, arg1, options := "?", "?", "[?]"
	if e.arg0 != "" {
		arg0 = fmt.Sprintf("%q", e.arg0)
	}
	if e.arg1 != "" {
		arg1 = fmt.Sprintf("%q", e.arg1)
	}
	if e.options != nil {
		options = fmt.Sprintf("%v", e.options)
	}
	return fmt.Sprintf("DB(%s).Delete(ctx, %s, %s, %s)", e.DB().name, arg0, arg1, options)
}

// ExpectedDeleteAttachment represents an expectation for a call to DB.DeleteAttachment().
type ExpectedDeleteAttachment struct {
	commonExpectation
	arg0 string
	arg1 string
	arg2 string
	ret0 string
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

func (e *ExpectedDeleteAttachment) met(ex expectation) bool {
	exp := ex.(*ExpectedDeleteAttachment)
	if exp.arg0 != "" && exp.arg0 != e.arg0 {
		return false
	}
	if exp.arg1 != "" && exp.arg1 != e.arg1 {
		return false
	}
	if exp.arg2 != "" && exp.arg2 != e.arg2 {
		return false
	}
	return true
}

func (e *ExpectedDeleteAttachment) method(v bool) string {
	if !v {
		return "DB.DeleteAttachment()"
	}
	arg0, arg1, arg2, options := "?", "?", "?", "[?]"
	if e.arg0 != "" {
		arg0 = fmt.Sprintf("%q", e.arg0)
	}
	if e.arg1 != "" {
		arg1 = fmt.Sprintf("%q", e.arg1)
	}
	if e.arg2 != "" {
		arg2 = fmt.Sprintf("%q", e.arg2)
	}
	if e.options != nil {
		options = fmt.Sprintf("%v", e.options)
	}
	return fmt.Sprintf("DB(%s).DeleteAttachment(ctx, %s, %s, %s, %s)", e.DB().name, arg0, arg1, arg2, options)
}

// ExpectedDeleteIndex represents an expectation for a call to DB.DeleteIndex().
type ExpectedDeleteIndex struct {
	commonExpectation
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

func (e *ExpectedDeleteIndex) met(ex expectation) bool {
	exp := ex.(*ExpectedDeleteIndex)
	if exp.arg0 != "" && exp.arg0 != e.arg0 {
		return false
	}
	if exp.arg1 != "" && exp.arg1 != e.arg1 {
		return false
	}
	return true
}

func (e *ExpectedDeleteIndex) method(v bool) string {
	if !v {
		return "DB.DeleteIndex()"
	}
	arg0, arg1 := "?", "?"
	if e.arg0 != "" {
		arg0 = fmt.Sprintf("%q", e.arg0)
	}
	if e.arg1 != "" {
		arg1 = fmt.Sprintf("%q", e.arg1)
	}
	return fmt.Sprintf("DB(%s).DeleteIndex(ctx, %s, %s)", e.DB().name, arg0, arg1)
}

// ExpectedFlush represents an expectation for a call to DB.Flush().
type ExpectedFlush struct {
	commonExpectation
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

func (e *ExpectedFlush) met(_ expectation) bool {
	return true
}

func (e *ExpectedFlush) method(v bool) string {
	if !v {
		return "DB.Flush()"
	}
	return fmt.Sprintf("DB(%s).Flush(ctx)", e.DB().name)
}

// ExpectedGetMeta represents an expectation for a call to DB.GetMeta().
type ExpectedGetMeta struct {
	commonExpectation
	arg0 string
	ret0 int64
	ret1 string
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

func (e *ExpectedGetMeta) met(ex expectation) bool {
	exp := ex.(*ExpectedGetMeta)
	if exp.arg0 != "" && exp.arg0 != e.arg0 {
		return false
	}
	return true
}

func (e *ExpectedGetMeta) method(v bool) string {
	if !v {
		return "DB.GetMeta()"
	}
	arg0, options := "?", "[?]"
	if e.arg0 != "" {
		arg0 = fmt.Sprintf("%q", e.arg0)
	}
	if e.options != nil {
		options = fmt.Sprintf("%v", e.options)
	}
	return fmt.Sprintf("DB(%s).GetMeta(ctx, %s, %s)", e.DB().name, arg0, options)
}

// ExpectedPut represents an expectation for a call to DB.Put().
type ExpectedPut struct {
	commonExpectation
	arg0 string
	arg1 interface{}
	ret0 string
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

func (e *ExpectedPut) met(ex expectation) bool {
	exp := ex.(*ExpectedPut)
	if exp.arg0 != "" && exp.arg0 != e.arg0 {
		return false
	}
	if exp.arg1 != nil && !jsonMeets(exp.arg1, e.arg1) {
		return false
	}
	return true
}

func (e *ExpectedPut) method(v bool) string {
	if !v {
		return "DB.Put()"
	}
	arg0, arg1, options := "?", "?", "[?]"
	if e.arg0 != "" {
		arg0 = fmt.Sprintf("%q", e.arg0)
	}
	if e.arg1 != nil {
		arg1 = fmt.Sprintf("%v", e.arg1)
	}
	if e.options != nil {
		options = fmt.Sprintf("%v", e.options)
	}
	return fmt.Sprintf("DB(%s).Put(ctx, %s, %s, %s)", e.DB().name, arg0, arg1, options)
}

// ExpectedViewCleanup represents an expectation for a call to DB.ViewCleanup().
type ExpectedViewCleanup struct {
	commonExpectation
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

func (e *ExpectedViewCleanup) met(_ expectation) bool {
	return true
}

func (e *ExpectedViewCleanup) method(v bool) string {
	if !v {
		return "DB.ViewCleanup()"
	}
	return fmt.Sprintf("DB(%s).ViewCleanup(ctx)", e.DB().name)
}

// ExpectedAllDocs represents an expectation for a call to DB.AllDocs().
type ExpectedAllDocs struct {
	commonExpectation
	ret0 *Rows
}

// WithOptions sets the expected options for the call to DB.AllDocs().
func (e *ExpectedAllDocs) WithOptions(options map[string]interface{}) *ExpectedAllDocs {
	e.options = options
	return e
}

// WillReturn sets the values that will be returned by the call to DB.AllDocs().
func (e *ExpectedAllDocs) WillReturn(ret0 *Rows) *ExpectedAllDocs {
	e.ret0 = ret0
	return e
}

// WillReturnError sets the error value that will be returned by the call to DB.AllDocs().
func (e *ExpectedAllDocs) WillReturnError(err error) *ExpectedAllDocs {
	e.err = err
	return e
}

// WillDelay causes the call to DB.AllDocs() to delay.
func (e *ExpectedAllDocs) WillDelay(delay time.Duration) *ExpectedAllDocs {
	e.delay = delay
	return e
}

func (e *ExpectedAllDocs) met(_ expectation) bool {
	return true
}

func (e *ExpectedAllDocs) method(v bool) string {
	if !v {
		return "DB.AllDocs()"
	}
	options := "[?]"
	if e.options != nil {
		options = fmt.Sprintf("%v", e.options)
	}
	return fmt.Sprintf("DB(%s).AllDocs(ctx, %s)", e.DB().name, options)
}

// ExpectedBulkDocs represents an expectation for a call to DB.BulkDocs().
type ExpectedBulkDocs struct {
	commonExpectation
	arg0 []interface{}
	ret0 driver.BulkResults
}

// WithOptions sets the expected options for the call to DB.BulkDocs().
func (e *ExpectedBulkDocs) WithOptions(options map[string]interface{}) *ExpectedBulkDocs {
	e.options = options
	return e
}

// WillReturn sets the values that will be returned by the call to DB.BulkDocs().
func (e *ExpectedBulkDocs) WillReturn(ret0 driver.BulkResults) *ExpectedBulkDocs {
	e.ret0 = ret0
	return e
}

// WillReturnError sets the error value that will be returned by the call to DB.BulkDocs().
func (e *ExpectedBulkDocs) WillReturnError(err error) *ExpectedBulkDocs {
	e.err = err
	return e
}

// WillDelay causes the call to DB.BulkDocs() to delay.
func (e *ExpectedBulkDocs) WillDelay(delay time.Duration) *ExpectedBulkDocs {
	e.delay = delay
	return e
}

func (e *ExpectedBulkDocs) met(ex expectation) bool {
	exp := ex.(*ExpectedBulkDocs)
	if exp.arg0 != nil && !reflect.DeepEqual(exp.arg0, e.arg0) {
		return false
	}
	return true
}

func (e *ExpectedBulkDocs) method(v bool) string {
	if !v {
		return "DB.BulkDocs()"
	}
	arg0, options := "?", "[?]"
	if e.arg0 != nil {
		arg0 = fmt.Sprintf("%v", e.arg0)
	}
	if e.options != nil {
		options = fmt.Sprintf("%v", e.options)
	}
	return fmt.Sprintf("DB(%s).BulkDocs(ctx, %s, %s)", e.DB().name, arg0, options)
}

// ExpectedBulkGet represents an expectation for a call to DB.BulkGet().
type ExpectedBulkGet struct {
	commonExpectation
	arg0 []driver.BulkGetReference
	ret0 *Rows
}

// WithOptions sets the expected options for the call to DB.BulkGet().
func (e *ExpectedBulkGet) WithOptions(options map[string]interface{}) *ExpectedBulkGet {
	e.options = options
	return e
}

// WillReturn sets the values that will be returned by the call to DB.BulkGet().
func (e *ExpectedBulkGet) WillReturn(ret0 *Rows) *ExpectedBulkGet {
	e.ret0 = ret0
	return e
}

// WillReturnError sets the error value that will be returned by the call to DB.BulkGet().
func (e *ExpectedBulkGet) WillReturnError(err error) *ExpectedBulkGet {
	e.err = err
	return e
}

// WillDelay causes the call to DB.BulkGet() to delay.
func (e *ExpectedBulkGet) WillDelay(delay time.Duration) *ExpectedBulkGet {
	e.delay = delay
	return e
}

func (e *ExpectedBulkGet) met(ex expectation) bool {
	exp := ex.(*ExpectedBulkGet)
	if exp.arg0 != nil && !reflect.DeepEqual(exp.arg0, e.arg0) {
		return false
	}
	return true
}

func (e *ExpectedBulkGet) method(v bool) string {
	if !v {
		return "DB.BulkGet()"
	}
	arg0, options := "?", "[?]"
	if e.arg0 != nil {
		arg0 = fmt.Sprintf("%v", e.arg0)
	}
	if e.options != nil {
		options = fmt.Sprintf("%v", e.options)
	}
	return fmt.Sprintf("DB(%s).BulkGet(ctx, %s, %s)", e.DB().name, arg0, options)
}

// ExpectedChanges represents an expectation for a call to DB.Changes().
type ExpectedChanges struct {
	commonExpectation
	ret0 driver.Changes
}

// WithOptions sets the expected options for the call to DB.Changes().
func (e *ExpectedChanges) WithOptions(options map[string]interface{}) *ExpectedChanges {
	e.options = options
	return e
}

// WillReturn sets the values that will be returned by the call to DB.Changes().
func (e *ExpectedChanges) WillReturn(ret0 driver.Changes) *ExpectedChanges {
	e.ret0 = ret0
	return e
}

// WillReturnError sets the error value that will be returned by the call to DB.Changes().
func (e *ExpectedChanges) WillReturnError(err error) *ExpectedChanges {
	e.err = err
	return e
}

// WillDelay causes the call to DB.Changes() to delay.
func (e *ExpectedChanges) WillDelay(delay time.Duration) *ExpectedChanges {
	e.delay = delay
	return e
}

func (e *ExpectedChanges) met(_ expectation) bool {
	return true
}

func (e *ExpectedChanges) method(v bool) string {
	if !v {
		return "DB.Changes()"
	}
	options := "[?]"
	if e.options != nil {
		options = fmt.Sprintf("%v", e.options)
	}
	return fmt.Sprintf("DB(%s).Changes(ctx, %s)", e.DB().name, options)
}

// ExpectedDesignDocs represents an expectation for a call to DB.DesignDocs().
type ExpectedDesignDocs struct {
	commonExpectation
	ret0 *Rows
}

// WithOptions sets the expected options for the call to DB.DesignDocs().
func (e *ExpectedDesignDocs) WithOptions(options map[string]interface{}) *ExpectedDesignDocs {
	e.options = options
	return e
}

// WillReturn sets the values that will be returned by the call to DB.DesignDocs().
func (e *ExpectedDesignDocs) WillReturn(ret0 *Rows) *ExpectedDesignDocs {
	e.ret0 = ret0
	return e
}

// WillReturnError sets the error value that will be returned by the call to DB.DesignDocs().
func (e *ExpectedDesignDocs) WillReturnError(err error) *ExpectedDesignDocs {
	e.err = err
	return e
}

// WillDelay causes the call to DB.DesignDocs() to delay.
func (e *ExpectedDesignDocs) WillDelay(delay time.Duration) *ExpectedDesignDocs {
	e.delay = delay
	return e
}

func (e *ExpectedDesignDocs) met(_ expectation) bool {
	return true
}

func (e *ExpectedDesignDocs) method(v bool) string {
	if !v {
		return "DB.DesignDocs()"
	}
	options := "[?]"
	if e.options != nil {
		options = fmt.Sprintf("%v", e.options)
	}
	return fmt.Sprintf("DB(%s).DesignDocs(ctx, %s)", e.DB().name, options)
}

// ExpectedExplain represents an expectation for a call to DB.Explain().
type ExpectedExplain struct {
	commonExpectation
	arg0 interface{}
	ret0 *driver.QueryPlan
}

// WillReturn sets the values that will be returned by the call to DB.Explain().
func (e *ExpectedExplain) WillReturn(ret0 *driver.QueryPlan) *ExpectedExplain {
	e.ret0 = ret0
	return e
}

// WillReturnError sets the error value that will be returned by the call to DB.Explain().
func (e *ExpectedExplain) WillReturnError(err error) *ExpectedExplain {
	e.err = err
	return e
}

// WillDelay causes the call to DB.Explain() to delay.
func (e *ExpectedExplain) WillDelay(delay time.Duration) *ExpectedExplain {
	e.delay = delay
	return e
}

func (e *ExpectedExplain) met(ex expectation) bool {
	exp := ex.(*ExpectedExplain)
	if exp.arg0 != nil && !jsonMeets(exp.arg0, e.arg0) {
		return false
	}
	return true
}

func (e *ExpectedExplain) method(v bool) string {
	if !v {
		return "DB.Explain()"
	}
	arg0 := "?"
	if e.arg0 != nil {
		arg0 = fmt.Sprintf("%v", e.arg0)
	}
	return fmt.Sprintf("DB(%s).Explain(ctx, %s)", e.DB().name, arg0)
}

// ExpectedFind represents an expectation for a call to DB.Find().
type ExpectedFind struct {
	commonExpectation
	arg0 interface{}
	ret0 *Rows
}

// WillReturn sets the values that will be returned by the call to DB.Find().
func (e *ExpectedFind) WillReturn(ret0 *Rows) *ExpectedFind {
	e.ret0 = ret0
	return e
}

// WillReturnError sets the error value that will be returned by the call to DB.Find().
func (e *ExpectedFind) WillReturnError(err error) *ExpectedFind {
	e.err = err
	return e
}

// WillDelay causes the call to DB.Find() to delay.
func (e *ExpectedFind) WillDelay(delay time.Duration) *ExpectedFind {
	e.delay = delay
	return e
}

func (e *ExpectedFind) met(ex expectation) bool {
	exp := ex.(*ExpectedFind)
	if exp.arg0 != nil && !jsonMeets(exp.arg0, e.arg0) {
		return false
	}
	return true
}

func (e *ExpectedFind) method(v bool) string {
	if !v {
		return "DB.Find()"
	}
	arg0 := "?"
	if e.arg0 != nil {
		arg0 = fmt.Sprintf("%v", e.arg0)
	}
	return fmt.Sprintf("DB(%s).Find(ctx, %s)", e.DB().name, arg0)
}

// ExpectedGet represents an expectation for a call to DB.Get().
type ExpectedGet struct {
	commonExpectation
	arg0 string
	ret0 *driver.Document
}

// WithOptions sets the expected options for the call to DB.Get().
func (e *ExpectedGet) WithOptions(options map[string]interface{}) *ExpectedGet {
	e.options = options
	return e
}

// WillReturn sets the values that will be returned by the call to DB.Get().
func (e *ExpectedGet) WillReturn(ret0 *driver.Document) *ExpectedGet {
	e.ret0 = ret0
	return e
}

// WillReturnError sets the error value that will be returned by the call to DB.Get().
func (e *ExpectedGet) WillReturnError(err error) *ExpectedGet {
	e.err = err
	return e
}

// WillDelay causes the call to DB.Get() to delay.
func (e *ExpectedGet) WillDelay(delay time.Duration) *ExpectedGet {
	e.delay = delay
	return e
}

func (e *ExpectedGet) met(ex expectation) bool {
	exp := ex.(*ExpectedGet)
	if exp.arg0 != "" && exp.arg0 != e.arg0 {
		return false
	}
	return true
}

func (e *ExpectedGet) method(v bool) string {
	if !v {
		return "DB.Get()"
	}
	arg0, options := "?", "[?]"
	if e.arg0 != "" {
		arg0 = fmt.Sprintf("%q", e.arg0)
	}
	if e.options != nil {
		options = fmt.Sprintf("%v", e.options)
	}
	return fmt.Sprintf("DB(%s).Get(ctx, %s, %s)", e.DB().name, arg0, options)
}

// ExpectedGetAttachment represents an expectation for a call to DB.GetAttachment().
type ExpectedGetAttachment struct {
	commonExpectation
	arg0 string
	arg1 string
	ret0 *driver.Attachment
}

// WithOptions sets the expected options for the call to DB.GetAttachment().
func (e *ExpectedGetAttachment) WithOptions(options map[string]interface{}) *ExpectedGetAttachment {
	e.options = options
	return e
}

// WillReturn sets the values that will be returned by the call to DB.GetAttachment().
func (e *ExpectedGetAttachment) WillReturn(ret0 *driver.Attachment) *ExpectedGetAttachment {
	e.ret0 = ret0
	return e
}

// WillReturnError sets the error value that will be returned by the call to DB.GetAttachment().
func (e *ExpectedGetAttachment) WillReturnError(err error) *ExpectedGetAttachment {
	e.err = err
	return e
}

// WillDelay causes the call to DB.GetAttachment() to delay.
func (e *ExpectedGetAttachment) WillDelay(delay time.Duration) *ExpectedGetAttachment {
	e.delay = delay
	return e
}

func (e *ExpectedGetAttachment) met(ex expectation) bool {
	exp := ex.(*ExpectedGetAttachment)
	if exp.arg0 != "" && exp.arg0 != e.arg0 {
		return false
	}
	if exp.arg1 != "" && exp.arg1 != e.arg1 {
		return false
	}
	return true
}

func (e *ExpectedGetAttachment) method(v bool) string {
	if !v {
		return "DB.GetAttachment()"
	}
	arg0, arg1, options := "?", "?", "[?]"
	if e.arg0 != "" {
		arg0 = fmt.Sprintf("%q", e.arg0)
	}
	if e.arg1 != "" {
		arg1 = fmt.Sprintf("%q", e.arg1)
	}
	if e.options != nil {
		options = fmt.Sprintf("%v", e.options)
	}
	return fmt.Sprintf("DB(%s).GetAttachment(ctx, %s, %s, %s)", e.DB().name, arg0, arg1, options)
}

// ExpectedGetAttachmentMeta represents an expectation for a call to DB.GetAttachmentMeta().
type ExpectedGetAttachmentMeta struct {
	commonExpectation
	arg0 string
	arg1 string
	ret0 *driver.Attachment
}

// WithOptions sets the expected options for the call to DB.GetAttachmentMeta().
func (e *ExpectedGetAttachmentMeta) WithOptions(options map[string]interface{}) *ExpectedGetAttachmentMeta {
	e.options = options
	return e
}

// WillReturn sets the values that will be returned by the call to DB.GetAttachmentMeta().
func (e *ExpectedGetAttachmentMeta) WillReturn(ret0 *driver.Attachment) *ExpectedGetAttachmentMeta {
	e.ret0 = ret0
	return e
}

// WillReturnError sets the error value that will be returned by the call to DB.GetAttachmentMeta().
func (e *ExpectedGetAttachmentMeta) WillReturnError(err error) *ExpectedGetAttachmentMeta {
	e.err = err
	return e
}

// WillDelay causes the call to DB.GetAttachmentMeta() to delay.
func (e *ExpectedGetAttachmentMeta) WillDelay(delay time.Duration) *ExpectedGetAttachmentMeta {
	e.delay = delay
	return e
}

func (e *ExpectedGetAttachmentMeta) met(ex expectation) bool {
	exp := ex.(*ExpectedGetAttachmentMeta)
	if exp.arg0 != "" && exp.arg0 != e.arg0 {
		return false
	}
	if exp.arg1 != "" && exp.arg1 != e.arg1 {
		return false
	}
	return true
}

func (e *ExpectedGetAttachmentMeta) method(v bool) string {
	if !v {
		return "DB.GetAttachmentMeta()"
	}
	arg0, arg1, options := "?", "?", "[?]"
	if e.arg0 != "" {
		arg0 = fmt.Sprintf("%q", e.arg0)
	}
	if e.arg1 != "" {
		arg1 = fmt.Sprintf("%q", e.arg1)
	}
	if e.options != nil {
		options = fmt.Sprintf("%v", e.options)
	}
	return fmt.Sprintf("DB(%s).GetAttachmentMeta(ctx, %s, %s, %s)", e.DB().name, arg0, arg1, options)
}

// ExpectedGetIndexes represents an expectation for a call to DB.GetIndexes().
type ExpectedGetIndexes struct {
	commonExpectation
	ret0 []driver.Index
}

// WillReturn sets the values that will be returned by the call to DB.GetIndexes().
func (e *ExpectedGetIndexes) WillReturn(ret0 []driver.Index) *ExpectedGetIndexes {
	e.ret0 = ret0
	return e
}

// WillReturnError sets the error value that will be returned by the call to DB.GetIndexes().
func (e *ExpectedGetIndexes) WillReturnError(err error) *ExpectedGetIndexes {
	e.err = err
	return e
}

// WillDelay causes the call to DB.GetIndexes() to delay.
func (e *ExpectedGetIndexes) WillDelay(delay time.Duration) *ExpectedGetIndexes {
	e.delay = delay
	return e
}

func (e *ExpectedGetIndexes) met(_ expectation) bool {
	return true
}

func (e *ExpectedGetIndexes) method(v bool) string {
	if !v {
		return "DB.GetIndexes()"
	}
	return fmt.Sprintf("DB(%s).GetIndexes(ctx)", e.DB().name)
}

// ExpectedLocalDocs represents an expectation for a call to DB.LocalDocs().
type ExpectedLocalDocs struct {
	commonExpectation
	ret0 *Rows
}

// WithOptions sets the expected options for the call to DB.LocalDocs().
func (e *ExpectedLocalDocs) WithOptions(options map[string]interface{}) *ExpectedLocalDocs {
	e.options = options
	return e
}

// WillReturn sets the values that will be returned by the call to DB.LocalDocs().
func (e *ExpectedLocalDocs) WillReturn(ret0 *Rows) *ExpectedLocalDocs {
	e.ret0 = ret0
	return e
}

// WillReturnError sets the error value that will be returned by the call to DB.LocalDocs().
func (e *ExpectedLocalDocs) WillReturnError(err error) *ExpectedLocalDocs {
	e.err = err
	return e
}

// WillDelay causes the call to DB.LocalDocs() to delay.
func (e *ExpectedLocalDocs) WillDelay(delay time.Duration) *ExpectedLocalDocs {
	e.delay = delay
	return e
}

func (e *ExpectedLocalDocs) met(_ expectation) bool {
	return true
}

func (e *ExpectedLocalDocs) method(v bool) string {
	if !v {
		return "DB.LocalDocs()"
	}
	options := "[?]"
	if e.options != nil {
		options = fmt.Sprintf("%v", e.options)
	}
	return fmt.Sprintf("DB(%s).LocalDocs(ctx, %s)", e.DB().name, options)
}

// ExpectedPurge represents an expectation for a call to DB.Purge().
type ExpectedPurge struct {
	commonExpectation
	arg0 map[string][]string
	ret0 *driver.PurgeResult
}

// WillReturn sets the values that will be returned by the call to DB.Purge().
func (e *ExpectedPurge) WillReturn(ret0 *driver.PurgeResult) *ExpectedPurge {
	e.ret0 = ret0
	return e
}

// WillReturnError sets the error value that will be returned by the call to DB.Purge().
func (e *ExpectedPurge) WillReturnError(err error) *ExpectedPurge {
	e.err = err
	return e
}

// WillDelay causes the call to DB.Purge() to delay.
func (e *ExpectedPurge) WillDelay(delay time.Duration) *ExpectedPurge {
	e.delay = delay
	return e
}

func (e *ExpectedPurge) met(ex expectation) bool {
	exp := ex.(*ExpectedPurge)
	if exp.arg0 != nil && !reflect.DeepEqual(exp.arg0, e.arg0) {
		return false
	}
	return true
}

func (e *ExpectedPurge) method(v bool) string {
	if !v {
		return "DB.Purge()"
	}
	arg0 := "?"
	if e.arg0 != nil {
		arg0 = fmt.Sprintf("%v", e.arg0)
	}
	return fmt.Sprintf("DB(%s).Purge(ctx, %s)", e.DB().name, arg0)
}

// ExpectedPutAttachment represents an expectation for a call to DB.PutAttachment().
type ExpectedPutAttachment struct {
	commonExpectation
	arg0 string
	arg1 string
	arg2 *driver.Attachment
	ret0 string
}

// WithOptions sets the expected options for the call to DB.PutAttachment().
func (e *ExpectedPutAttachment) WithOptions(options map[string]interface{}) *ExpectedPutAttachment {
	e.options = options
	return e
}

// WillReturn sets the values that will be returned by the call to DB.PutAttachment().
func (e *ExpectedPutAttachment) WillReturn(ret0 string) *ExpectedPutAttachment {
	e.ret0 = ret0
	return e
}

// WillReturnError sets the error value that will be returned by the call to DB.PutAttachment().
func (e *ExpectedPutAttachment) WillReturnError(err error) *ExpectedPutAttachment {
	e.err = err
	return e
}

// WillDelay causes the call to DB.PutAttachment() to delay.
func (e *ExpectedPutAttachment) WillDelay(delay time.Duration) *ExpectedPutAttachment {
	e.delay = delay
	return e
}

func (e *ExpectedPutAttachment) met(ex expectation) bool {
	exp := ex.(*ExpectedPutAttachment)
	if exp.arg0 != "" && exp.arg0 != e.arg0 {
		return false
	}
	if exp.arg1 != "" && exp.arg1 != e.arg1 {
		return false
	}
	if exp.arg2 != nil && !reflect.DeepEqual(exp.arg2, e.arg2) {
		return false
	}
	return true
}

func (e *ExpectedPutAttachment) method(v bool) string {
	if !v {
		return "DB.PutAttachment()"
	}
	arg0, arg1, arg2, options := "?", "?", "?", "[?]"
	if e.arg0 != "" {
		arg0 = fmt.Sprintf("%q", e.arg0)
	}
	if e.arg1 != "" {
		arg1 = fmt.Sprintf("%q", e.arg1)
	}
	if e.arg2 != nil {
		arg2 = fmt.Sprintf("%v", e.arg2)
	}
	if e.options != nil {
		options = fmt.Sprintf("%v", e.options)
	}
	return fmt.Sprintf("DB(%s).PutAttachment(ctx, %s, %s, %s, %s)", e.DB().name, arg0, arg1, arg2, options)
}

// ExpectedQuery represents an expectation for a call to DB.Query().
type ExpectedQuery struct {
	commonExpectation
	arg0 string
	arg1 string
	ret0 *Rows
}

// WithOptions sets the expected options for the call to DB.Query().
func (e *ExpectedQuery) WithOptions(options map[string]interface{}) *ExpectedQuery {
	e.options = options
	return e
}

// WillReturn sets the values that will be returned by the call to DB.Query().
func (e *ExpectedQuery) WillReturn(ret0 *Rows) *ExpectedQuery {
	e.ret0 = ret0
	return e
}

// WillReturnError sets the error value that will be returned by the call to DB.Query().
func (e *ExpectedQuery) WillReturnError(err error) *ExpectedQuery {
	e.err = err
	return e
}

// WillDelay causes the call to DB.Query() to delay.
func (e *ExpectedQuery) WillDelay(delay time.Duration) *ExpectedQuery {
	e.delay = delay
	return e
}

func (e *ExpectedQuery) met(ex expectation) bool {
	exp := ex.(*ExpectedQuery)
	if exp.arg0 != "" && exp.arg0 != e.arg0 {
		return false
	}
	if exp.arg1 != "" && exp.arg1 != e.arg1 {
		return false
	}
	return true
}

func (e *ExpectedQuery) method(v bool) string {
	if !v {
		return "DB.Query()"
	}
	arg0, arg1, options := "?", "?", "[?]"
	if e.arg0 != "" {
		arg0 = fmt.Sprintf("%q", e.arg0)
	}
	if e.arg1 != "" {
		arg1 = fmt.Sprintf("%q", e.arg1)
	}
	if e.options != nil {
		options = fmt.Sprintf("%v", e.options)
	}
	return fmt.Sprintf("DB(%s).Query(ctx, %s, %s, %s)", e.DB().name, arg0, arg1, options)
}

// ExpectedSecurity represents an expectation for a call to DB.Security().
type ExpectedSecurity struct {
	commonExpectation
	ret0 *driver.Security
}

// WillReturn sets the values that will be returned by the call to DB.Security().
func (e *ExpectedSecurity) WillReturn(ret0 *driver.Security) *ExpectedSecurity {
	e.ret0 = ret0
	return e
}

// WillReturnError sets the error value that will be returned by the call to DB.Security().
func (e *ExpectedSecurity) WillReturnError(err error) *ExpectedSecurity {
	e.err = err
	return e
}

// WillDelay causes the call to DB.Security() to delay.
func (e *ExpectedSecurity) WillDelay(delay time.Duration) *ExpectedSecurity {
	e.delay = delay
	return e
}

func (e *ExpectedSecurity) met(_ expectation) bool {
	return true
}

func (e *ExpectedSecurity) method(v bool) string {
	if !v {
		return "DB.Security()"
	}
	return fmt.Sprintf("DB(%s).Security(ctx)", e.DB().name)
}

// ExpectedSetSecurity represents an expectation for a call to DB.SetSecurity().
type ExpectedSetSecurity struct {
	commonExpectation
	arg0 *driver.Security
}

// WillReturnError sets the error value that will be returned by the call to DB.SetSecurity().
func (e *ExpectedSetSecurity) WillReturnError(err error) *ExpectedSetSecurity {
	e.err = err
	return e
}

// WillDelay causes the call to DB.SetSecurity() to delay.
func (e *ExpectedSetSecurity) WillDelay(delay time.Duration) *ExpectedSetSecurity {
	e.delay = delay
	return e
}

func (e *ExpectedSetSecurity) met(ex expectation) bool {
	exp := ex.(*ExpectedSetSecurity)
	if exp.arg0 != nil && !reflect.DeepEqual(exp.arg0, e.arg0) {
		return false
	}
	return true
}

func (e *ExpectedSetSecurity) method(v bool) string {
	if !v {
		return "DB.SetSecurity()"
	}
	arg0 := "?"
	if e.arg0 != nil {
		arg0 = fmt.Sprintf("%v", e.arg0)
	}
	return fmt.Sprintf("DB(%s).SetSecurity(ctx, %s)", e.DB().name, arg0)
}

// ExpectedStats represents an expectation for a call to DB.Stats().
type ExpectedStats struct {
	commonExpectation
	ret0 *driver.DBStats
}

// WillReturn sets the values that will be returned by the call to DB.Stats().
func (e *ExpectedStats) WillReturn(ret0 *driver.DBStats) *ExpectedStats {
	e.ret0 = ret0
	return e
}

// WillReturnError sets the error value that will be returned by the call to DB.Stats().
func (e *ExpectedStats) WillReturnError(err error) *ExpectedStats {
	e.err = err
	return e
}

// WillDelay causes the call to DB.Stats() to delay.
func (e *ExpectedStats) WillDelay(delay time.Duration) *ExpectedStats {
	e.delay = delay
	return e
}

func (e *ExpectedStats) met(_ expectation) bool {
	return true
}

func (e *ExpectedStats) method(v bool) string {
	if !v {
		return "DB.Stats()"
	}
	return fmt.Sprintf("DB(%s).Stats(ctx)", e.DB().name)
}
