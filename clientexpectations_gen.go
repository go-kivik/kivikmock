/* This file is auto-generated. Do not edit it! */

package kivikmock

import "time"

// ExpectedAllDBs represents an expectation for a call to AllDBs().
type ExpectedAllDBs struct {
	commonExpectation
	options map[string]interface{}
	ret0    []string
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

// ExpectedClusterStatus represents an expectation for a call to ClusterStatus().
type ExpectedClusterStatus struct {
	commonExpectation
	options map[string]interface{}
	ret0    string
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

// ExpectedDBExists represents an expectation for a call to DBExists().
type ExpectedDBExists struct {
	commonExpectation
	arg0    string
	options map[string]interface{}
	ret0    bool
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

// ExpectedDestroyDB represents an expectation for a call to DestroyDB().
type ExpectedDestroyDB struct {
	commonExpectation
	arg0    string
	options map[string]interface{}
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
