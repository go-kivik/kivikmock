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

// ExpectedAllDBs represents an expectation for a call to AllDBs().
type ExpectedAllDBs struct {
	commonExpectation
	ret0 []string
}

// WithOptions sets the expected options for the call to AllDBs().
func (e *ExpectedAllDBs) WithOptions(options map[string]interface{}) *ExpectedAllDBs {
	e.options = options
	return e
}

// WillReturn sets the values that will be returned by the call to AllDBs().
func (e *ExpectedAllDBs) WillReturn(ret0 []string) *ExpectedAllDBs {
	e.ret0 = ret0
	return e
}

// WillReturnError sets the error value that will be returned by the call to AllDBs().
func (e *ExpectedAllDBs) WillReturnError(err error) *ExpectedAllDBs {
	e.err = err
	return e
}

// WillDelay causes the call to AllDBs() to delay.
func (e *ExpectedAllDBs) WillDelay(delay time.Duration) *ExpectedAllDBs {
	e.delay = delay
	return e
}

func (e *ExpectedAllDBs) met(_ expectation) bool {
	return true
}

func (e *ExpectedAllDBs) method(v bool) string {
	if !v {
		return "AllDBs()"
	}
	options := defaultOptionPlaceholder
	if e.options != nil {
		options = fmt.Sprintf("%v", e.options)
	}
	return fmt.Sprintf("AllDBs(ctx, %s)", options)
}

// ExpectedClose represents an expectation for a call to Close().
type ExpectedClose struct {
	commonExpectation
}

// WillReturnError sets the error value that will be returned by the call to Close().
func (e *ExpectedClose) WillReturnError(err error) *ExpectedClose {
	e.err = err
	return e
}

// WillDelay causes the call to Close() to delay.
func (e *ExpectedClose) WillDelay(delay time.Duration) *ExpectedClose {
	e.delay = delay
	return e
}

func (e *ExpectedClose) met(_ expectation) bool {
	return true
}

func (e *ExpectedClose) method(v bool) string {
	if !v {
		return "Close()"
	}
	return fmt.Sprintf("Close(ctx)")
}

// ExpectedClusterSetup represents an expectation for a call to ClusterSetup().
type ExpectedClusterSetup struct {
	commonExpectation
	arg0 interface{}
}

// WillReturnError sets the error value that will be returned by the call to ClusterSetup().
func (e *ExpectedClusterSetup) WillReturnError(err error) *ExpectedClusterSetup {
	e.err = err
	return e
}

// WillDelay causes the call to ClusterSetup() to delay.
func (e *ExpectedClusterSetup) WillDelay(delay time.Duration) *ExpectedClusterSetup {
	e.delay = delay
	return e
}

func (e *ExpectedClusterSetup) met(ex expectation) bool {
	exp := ex.(*ExpectedClusterSetup)
	if exp.arg0 != nil && !jsonMeets(exp.arg0, e.arg0) {
		return false
	}
	return true
}

func (e *ExpectedClusterSetup) method(v bool) string {
	if !v {
		return "ClusterSetup()"
	}
	arg0 := "?"
	if e.arg0 != nil {
		arg0 = fmt.Sprintf("%v", e.arg0)
	}
	return fmt.Sprintf("ClusterSetup(ctx, %s)", arg0)
}

// ExpectedClusterStatus represents an expectation for a call to ClusterStatus().
type ExpectedClusterStatus struct {
	commonExpectation
	ret0 string
}

// WithOptions sets the expected options for the call to ClusterStatus().
func (e *ExpectedClusterStatus) WithOptions(options map[string]interface{}) *ExpectedClusterStatus {
	e.options = options
	return e
}

// WillReturn sets the values that will be returned by the call to ClusterStatus().
func (e *ExpectedClusterStatus) WillReturn(ret0 string) *ExpectedClusterStatus {
	e.ret0 = ret0
	return e
}

// WillReturnError sets the error value that will be returned by the call to ClusterStatus().
func (e *ExpectedClusterStatus) WillReturnError(err error) *ExpectedClusterStatus {
	e.err = err
	return e
}

// WillDelay causes the call to ClusterStatus() to delay.
func (e *ExpectedClusterStatus) WillDelay(delay time.Duration) *ExpectedClusterStatus {
	e.delay = delay
	return e
}

func (e *ExpectedClusterStatus) met(_ expectation) bool {
	return true
}

func (e *ExpectedClusterStatus) method(v bool) string {
	if !v {
		return "ClusterStatus()"
	}
	options := defaultOptionPlaceholder
	if e.options != nil {
		options = fmt.Sprintf("%v", e.options)
	}
	return fmt.Sprintf("ClusterStatus(ctx, %s)", options)
}

// ExpectedDBExists represents an expectation for a call to DBExists().
type ExpectedDBExists struct {
	commonExpectation
	arg0 string
	ret0 bool
}

// WithOptions sets the expected options for the call to DBExists().
func (e *ExpectedDBExists) WithOptions(options map[string]interface{}) *ExpectedDBExists {
	e.options = options
	return e
}

// WillReturn sets the values that will be returned by the call to DBExists().
func (e *ExpectedDBExists) WillReturn(ret0 bool) *ExpectedDBExists {
	e.ret0 = ret0
	return e
}

// WillReturnError sets the error value that will be returned by the call to DBExists().
func (e *ExpectedDBExists) WillReturnError(err error) *ExpectedDBExists {
	e.err = err
	return e
}

// WillDelay causes the call to DBExists() to delay.
func (e *ExpectedDBExists) WillDelay(delay time.Duration) *ExpectedDBExists {
	e.delay = delay
	return e
}

func (e *ExpectedDBExists) met(ex expectation) bool {
	exp := ex.(*ExpectedDBExists)
	if exp.arg0 != "" && exp.arg0 != e.arg0 {
		return false
	}
	return true
}

func (e *ExpectedDBExists) method(v bool) string {
	if !v {
		return "DBExists()"
	}
	arg0, options := "?", defaultOptionPlaceholder
	if e.arg0 != "" {
		arg0 = fmt.Sprintf("%q", e.arg0)
	}
	if e.options != nil {
		options = fmt.Sprintf("%v", e.options)
	}
	return fmt.Sprintf("DBExists(ctx, %s, %s)", arg0, options)
}

// ExpectedDestroyDB represents an expectation for a call to DestroyDB().
type ExpectedDestroyDB struct {
	commonExpectation
	arg0 string
}

// WithOptions sets the expected options for the call to DestroyDB().
func (e *ExpectedDestroyDB) WithOptions(options map[string]interface{}) *ExpectedDestroyDB {
	e.options = options
	return e
}

// WillReturnError sets the error value that will be returned by the call to DestroyDB().
func (e *ExpectedDestroyDB) WillReturnError(err error) *ExpectedDestroyDB {
	e.err = err
	return e
}

// WillDelay causes the call to DestroyDB() to delay.
func (e *ExpectedDestroyDB) WillDelay(delay time.Duration) *ExpectedDestroyDB {
	e.delay = delay
	return e
}

func (e *ExpectedDestroyDB) met(ex expectation) bool {
	exp := ex.(*ExpectedDestroyDB)
	if exp.arg0 != "" && exp.arg0 != e.arg0 {
		return false
	}
	return true
}

func (e *ExpectedDestroyDB) method(v bool) string {
	if !v {
		return "DestroyDB()"
	}
	arg0, options := "?", defaultOptionPlaceholder
	if e.arg0 != "" {
		arg0 = fmt.Sprintf("%q", e.arg0)
	}
	if e.options != nil {
		options = fmt.Sprintf("%v", e.options)
	}
	return fmt.Sprintf("DestroyDB(ctx, %s, %s)", arg0, options)
}

// ExpectedPing represents an expectation for a call to Ping().
type ExpectedPing struct {
	commonExpectation
	ret0 bool
}

// WillReturn sets the values that will be returned by the call to Ping().
func (e *ExpectedPing) WillReturn(ret0 bool) *ExpectedPing {
	e.ret0 = ret0
	return e
}

// WillReturnError sets the error value that will be returned by the call to Ping().
func (e *ExpectedPing) WillReturnError(err error) *ExpectedPing {
	e.err = err
	return e
}

// WillDelay causes the call to Ping() to delay.
func (e *ExpectedPing) WillDelay(delay time.Duration) *ExpectedPing {
	e.delay = delay
	return e
}

func (e *ExpectedPing) met(_ expectation) bool {
	return true
}

func (e *ExpectedPing) method(v bool) string {
	if !v {
		return "Ping()"
	}
	return fmt.Sprintf("Ping(ctx)")
}

// ExpectedDB represents an expectation for a call to DB().
type ExpectedDB struct {
	commonExpectation
	arg0 string
	ret0 *MockDB
}

// WithOptions sets the expected options for the call to DB().
func (e *ExpectedDB) WithOptions(options map[string]interface{}) *ExpectedDB {
	e.options = options
	return e
}

// WillReturn sets the values that will be returned by the call to DB().
func (e *ExpectedDB) WillReturn(ret0 *MockDB) *ExpectedDB {
	e.ret0 = ret0
	return e
}

// WillReturnError sets the error value that will be returned by the call to DB().
func (e *ExpectedDB) WillReturnError(err error) *ExpectedDB {
	e.err = err
	return e
}

// WillDelay causes the call to DB() to delay.
func (e *ExpectedDB) WillDelay(delay time.Duration) *ExpectedDB {
	e.delay = delay
	return e
}

func (e *ExpectedDB) met(ex expectation) bool {
	exp := ex.(*ExpectedDB)
	if exp.arg0 != "" && exp.arg0 != e.arg0 {
		return false
	}
	return true
}

func (e *ExpectedDB) method(v bool) string {
	if !v {
		return "DB()"
	}
	arg0, options := "?", defaultOptionPlaceholder
	if e.arg0 != "" {
		arg0 = fmt.Sprintf("%q", e.arg0)
	}
	if e.options != nil {
		options = fmt.Sprintf("%v", e.options)
	}
	return fmt.Sprintf("DB(ctx, %s, %s)", arg0, options)
}

// ExpectedDBUpdates represents an expectation for a call to DBUpdates().
type ExpectedDBUpdates struct {
	commonExpectation
	ret0 *Updates
}

// WillReturn sets the values that will be returned by the call to DBUpdates().
func (e *ExpectedDBUpdates) WillReturn(ret0 *Updates) *ExpectedDBUpdates {
	e.ret0 = ret0
	return e
}

// WillReturnError sets the error value that will be returned by the call to DBUpdates().
func (e *ExpectedDBUpdates) WillReturnError(err error) *ExpectedDBUpdates {
	e.err = err
	return e
}

// WillDelay causes the call to DBUpdates() to delay.
func (e *ExpectedDBUpdates) WillDelay(delay time.Duration) *ExpectedDBUpdates {
	e.delay = delay
	return e
}

func (e *ExpectedDBUpdates) met(_ expectation) bool {
	return true
}

func (e *ExpectedDBUpdates) method(v bool) string {
	if !v {
		return "DBUpdates()"
	}
	return fmt.Sprintf("DBUpdates(ctx)")
}

// ExpectedDBsStats represents an expectation for a call to DBsStats().
type ExpectedDBsStats struct {
	commonExpectation
	arg0 []string
	ret0 []*driver.DBStats
}

// WillReturn sets the values that will be returned by the call to DBsStats().
func (e *ExpectedDBsStats) WillReturn(ret0 []*driver.DBStats) *ExpectedDBsStats {
	e.ret0 = ret0
	return e
}

// WillReturnError sets the error value that will be returned by the call to DBsStats().
func (e *ExpectedDBsStats) WillReturnError(err error) *ExpectedDBsStats {
	e.err = err
	return e
}

// WillDelay causes the call to DBsStats() to delay.
func (e *ExpectedDBsStats) WillDelay(delay time.Duration) *ExpectedDBsStats {
	e.delay = delay
	return e
}

func (e *ExpectedDBsStats) met(ex expectation) bool {
	exp := ex.(*ExpectedDBsStats)
	if exp.arg0 != nil && !reflect.DeepEqual(exp.arg0, e.arg0) {
		return false
	}
	return true
}

func (e *ExpectedDBsStats) method(v bool) string {
	if !v {
		return "DBsStats()"
	}
	arg0 := "?"
	if e.arg0 != nil {
		arg0 = fmt.Sprintf("%v", e.arg0)
	}
	return fmt.Sprintf("DBsStats(ctx, %s)", arg0)
}

// ExpectedSession represents an expectation for a call to Session().
type ExpectedSession struct {
	commonExpectation
	ret0 *driver.Session
}

// WillReturn sets the values that will be returned by the call to Session().
func (e *ExpectedSession) WillReturn(ret0 *driver.Session) *ExpectedSession {
	e.ret0 = ret0
	return e
}

// WillReturnError sets the error value that will be returned by the call to Session().
func (e *ExpectedSession) WillReturnError(err error) *ExpectedSession {
	e.err = err
	return e
}

// WillDelay causes the call to Session() to delay.
func (e *ExpectedSession) WillDelay(delay time.Duration) *ExpectedSession {
	e.delay = delay
	return e
}

func (e *ExpectedSession) met(_ expectation) bool {
	return true
}

func (e *ExpectedSession) method(v bool) string {
	if !v {
		return "Session()"
	}
	return fmt.Sprintf("Session(ctx)")
}

// ExpectedVersion represents an expectation for a call to Version().
type ExpectedVersion struct {
	commonExpectation
	ret0 *driver.Version
}

// WillReturn sets the values that will be returned by the call to Version().
func (e *ExpectedVersion) WillReturn(ret0 *driver.Version) *ExpectedVersion {
	e.ret0 = ret0
	return e
}

// WillReturnError sets the error value that will be returned by the call to Version().
func (e *ExpectedVersion) WillReturnError(err error) *ExpectedVersion {
	e.err = err
	return e
}

// WillDelay causes the call to Version() to delay.
func (e *ExpectedVersion) WillDelay(delay time.Duration) *ExpectedVersion {
	e.delay = delay
	return e
}

func (e *ExpectedVersion) met(_ expectation) bool {
	return true
}

func (e *ExpectedVersion) method(v bool) string {
	if !v {
		return "Version()"
	}
	return fmt.Sprintf("Version(ctx)")
}
