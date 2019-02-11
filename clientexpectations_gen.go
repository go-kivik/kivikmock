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

