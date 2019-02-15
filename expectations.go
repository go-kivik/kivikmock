package kivikmock

import (
	"fmt"
	"reflect"
	"time"
)

func (e *ExpectedClose) String() string {
	extra := delayString(e.delay) + errorString(e.err)
	msg := "call to Close()"
	if extra != "" {
		msg += " which:" + extra
	}
	return msg
}

// String satisfies the fmt.Stringer interface.
func (e *ExpectedAllDBs) String() string {
	return "call to AllDBs() which:" +
		optionsString(e.options) +
		delayString(e.delay) +
		errorString(e.err)
}

// ExpectedAuthenticate is used to manage *kivik.Client.Authenticate
// expectation returned by Mock.ExpectAuthenticate.
type ExpectedAuthenticate struct {
	commonExpectation
	authType string
}

// String satisfies the fmt.Stringer interface.
func (e *ExpectedAuthenticate) String() string {
	msg := fmt.Sprintf("call to %s which:", e.method(false))
	if e.authType == "" {
		msg += "\n\t- has any authenticator"
	} else {
		msg += fmt.Sprint("\n\t- has an authenticator of type: " + e.authType)
	}
	msg += delayString(e.delay)
	msg += errorString(e.err)
	return msg
}

func (e *ExpectedAuthenticate) method(v bool) string {
	if !v {
		return "Authenticate()"
	}
	if e.authType == "" {
		return "Authenticate(ctx, ?)"
	}
	return fmt.Sprintf("Authenticate(ctx, <%s>)", e.authType)
}

func (e *ExpectedAuthenticate) met(ex expectation) bool {
	exp := ex.(*ExpectedAuthenticate)
	if exp.authType == "" {
		return true
	}
	return e.authType == exp.authType
}

// WillReturnError allows setting an error for *kivik.Client.Authenticate action.
func (e *ExpectedAuthenticate) WillReturnError(err error) *ExpectedAuthenticate {
	e.err = err
	return e
}

// WithAuthenticator will match the the provide authenticator _type_ against
// that provided. There is no way to validate the authenticated credentials
// with this method.
func (e *ExpectedAuthenticate) WithAuthenticator(authenticator interface{}) *ExpectedAuthenticate {
	e.authType = reflect.TypeOf(authenticator).Name()
	return e
}

// WillDelay will cause execution of Authenticate() to delay by duration d.
func (e *ExpectedAuthenticate) WillDelay(d time.Duration) *ExpectedAuthenticate {
	e.delay = d
	return e
}

func (e *ExpectedClusterSetup) String() string {
	msg := "call to ClusterSetup() which:"
	if e.arg0 == nil {
		msg += "\n\t- has any action"
	} else {
		msg += fmt.Sprintf("\n\t- has the action: %v", e.arg0)
	}
	msg += delayString(e.delay)
	msg += errorString(e.err)
	return msg
}

// String satisfies the fmt.Stringer interface
func (e *ExpectedClusterStatus) String() string {
	return "call to ClusterStatus() which:" +
		optionsString(e.options) +
		delayString(e.delay) +
		errorString(e.err)
}

// WithAction specifies the action to be matched. Note that this expectation
// is compared with the actual action's marshaled JSON output, so it is not
// essential that the data types match exactly, in a Go sense.
func (e *ExpectedClusterSetup) WithAction(action interface{}) *ExpectedClusterSetup {
	e.arg0 = action
	return e
}

func (e *ExpectedDBExists) String() string {
	msg := "call to DBExists() which:" +
		nameString(e.arg0) +
		optionsString(e.options) +
		delayString(e.delay)
	if e.err == nil {
		msg += fmt.Sprintf("\n\t- should return: %t", e.ret0)
	} else {
		msg += fmt.Sprintf("\n\t- should return error: %s", e.err)
	}
	return msg
}

// WithName sets the expectation that DBExists will be called with the provided
// name.
func (e *ExpectedDBExists) WithName(name string) *ExpectedDBExists {
	e.arg0 = name
	return e
}

func (e *ExpectedDestroyDB) String() string {
	return "call to DestroyDB() which:" +
		nameString(e.arg0) +
		optionsString(e.options) +
		delayString(e.delay) +
		errorString(e.err)
}

// WithName sets the expectation that DestroyDB will be called with this name.
func (e *ExpectedDestroyDB) WithName(name string) *ExpectedDestroyDB {
	e.arg0 = name
	return e
}

func (e *ExpectedDBsStats) String() string {
	msg := "call to DBsStats() which:"
	if e.arg0 == nil {
		msg += "\n\t- has any names"
	} else {
		msg += fmt.Sprintf("\n\t- has names: %s", e.arg0)
	}
	return msg + delayString(e.delay) + errorString(e.err)
}

// WithNames sets the expectation that DBsStats will be called with these names.
func (e *ExpectedDBsStats) WithNames(names []string) *ExpectedDBsStats {
	e.arg0 = names
	return e
}

func (e *ExpectedPing) String() string {
	msg := "call to Ping()"
	extra := delayString(e.delay) + errorString(e.err)
	if extra != "" {
		msg += " which:" + extra
	}
	return msg
}

func (e *ExpectedSession) String() string {
	msg := "call to Session()"
	extra := ""
	if e.ret0 != nil {
		extra += fmt.Sprintf("\n\t- should return: %v", e.ret0)
	}
	extra += delayString(e.delay) + errorString(e.err)
	if extra != "" {
		msg += " which:" + extra
	}
	return msg
}

func (e *ExpectedVersion) String() string {
	msg := "call to Version()"
	extra := ""
	if e.ret0 != nil {
		extra += fmt.Sprintf("\n\t- should return: %v", e.ret0)
	}
	extra += delayString(e.delay) + errorString(e.err)
	if extra != "" {
		msg += " which:" + extra
	}
	return msg
}

func (e *ExpectedDB) String() string {
	msg := "call to DB() which:" +
		nameString(e.arg0) +
		optionsString(e.options)
	if e.db != nil {
		msg += fmt.Sprintf("\n\t- should return database with %d expectations", e.db.expectations())
	}
	msg += delayString(e.delay)
	return msg
}

// WithName sets the expectation that DB() will be called with this name.
func (e *ExpectedDB) WithName(name string) *ExpectedDB {
	e.arg0 = name
	return e
}

// ExpectedCreateDB represents an expectation to call the CreateDB() method.
//
// Implementation note: Because kivik always calls DB() after a
// successful CreateDB() is executed, ExpectCreateDB() creates two
// expectations under the covers, one for the backend CreateDB() call,
// and one for the DB() call. If WillReturnError() is called, the DB() call
// expectation is removed.
type ExpectedCreateDB struct {
	commonExpectation
	arg0       string
	expectedDB *ExpectedDB
}

func (e *ExpectedCreateDB) String() string {
	msg := "call to CreateDB() which:" +
		nameString(e.arg0) +
		optionsString(e.options)
	if e.db != nil {
		msg += fmt.Sprintf("\n\t- should return database with %d expectations", e.db.expectations())
	}
	msg += delayString(e.delay)
	msg += errorString(e.err)
	return msg
}

func (e *ExpectedCreateDB) method(v bool) string {
	if !v {
		return "CreateDB()"
	}
	var name, options string
	if e.arg0 == "" {
		name = "?"
	} else {
		name = fmt.Sprintf("%q", e.arg0)
	}
	if e.options != nil {
		options = fmt.Sprintf(", %v", e.options)
	}
	return fmt.Sprintf("CreateDB(ctx, %s%s)", name, options)
}

func (e *ExpectedCreateDB) met(ex expectation) bool {
	exp := ex.(*ExpectedCreateDB)
	nameOK := exp.arg0 == "" || exp.arg0 == e.arg0
	optionsOK := exp.options == nil || reflect.DeepEqual(exp.options, e.options)
	return nameOK && optionsOK
}

// WithOptions set the expectation that DB() will be called with these options.
func (e *ExpectedCreateDB) WithOptions(options map[string]interface{}) *ExpectedCreateDB {
	e.options = options
	return e
}

// WithName sets the expectation that DB() will be called with this name.
func (e *ExpectedCreateDB) WithName(name string) *ExpectedCreateDB {
	e.expectedDB.arg0 = name
	e.arg0 = name
	return e
}

// WillReturn sets the return value for the DB() call.
func (e *ExpectedCreateDB) WillReturn(db *MockDB) *ExpectedCreateDB {
	e.expectedDB.ret0 = db
	return e
}

// WillReturnError sets the return value for the DB() call.
func (e *ExpectedCreateDB) WillReturnError(err error) *ExpectedCreateDB {
	e.err = err
	return e
}

// WillDelay will cause execution of DB() to delay by duration d.
func (e *ExpectedCreateDB) WillDelay(delay time.Duration) *ExpectedCreateDB {
	e.delay = delay
	return e
}

func (e *ExpectedDBUpdates) String() string {
	msg := "call to DBUpdates()"
	var extra string
	if e.ret0 != nil {
		extra += fmt.Sprintf("\n\t- should return: %d results", e.ret0.count())
	}
	extra += delayString(e.delay)
	extra += errorString(e.err)
	if extra != "" {
		msg += " which:" + extra
	}
	return msg
}
