package kivikmock

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sync"

	"github.com/go-kivik/kivik"
)

type expectation interface {
	fulfilled() bool
	Lock()
	Unlock()
	String() string
}

// commonExpectation satisfies the expectation interface, except the String()
// method.
type commonExpectation struct {
	sync.Mutex
	triggered bool
	err       error // nolint: structcheck
}

func (e *commonExpectation) fulfilled() bool {
	return e.triggered
}

// ExpectedClose is used to manage *kivik.Client.Close expectation returned
// by Mock.ExpectClose.
type ExpectedClose struct {
	commonExpectation
}

// WillReturnError allows setting an error for *kivik.Client.Close action.
func (e *ExpectedClose) WillReturnError(err error) *ExpectedClose {
	e.err = err
	return e
}

func (e *ExpectedClose) String() string {
	msg := "ExpectedClose => expecting client Close"
	if e.err != nil {
		msg += fmt.Sprintf(", which should return error: %s", e.err)
	}
	return msg
}

// ExpectedAllDBs is used to manage *kivik.Client.AllDBs expectation returned
// by Mock.ExpectAllDBs.
type ExpectedAllDBs struct {
	commonExpectation
	options map[string]interface{}
	results []string
}

func (e *ExpectedAllDBs) String() string {
	msg := "ExpectedAllDBs => expecting AllDBs which:"
	if e.options == nil {
		msg += "\n\t- is without options"
	} else {
		msg += fmt.Sprintf("\n\t- is with options %+v", e.options)
	}
	if len(e.results) > 0 {
		msg += fmt.Sprintf("\n\t- should return: %v", e.results)
	}
	if e.err != nil {
		msg += fmt.Sprintf("\n\t- should return error: %s", e.err)
	}
	return msg
}

// WillReturnError allows setting an error for *kivik.Client.Close action.
func (e *ExpectedAllDBs) WillReturnError(err error) *ExpectedAllDBs {
	e.err = err
	return e
}

// WithOptions will match the provided options against actual options passed
// during execution.
func (e *ExpectedAllDBs) WithOptions(options kivik.Options) *ExpectedAllDBs {
	e.options = options
	return e
}

// WillReturn sets the expected results.
func (e *ExpectedAllDBs) WillReturn(results []string) *ExpectedAllDBs {
	e.results = results
	return e
}

// ExpectedAuthenticate is used to manage *kivik.Client.Authenticate
// expectation returned by Mock.ExpectAuthenticate.
type ExpectedAuthenticate struct {
	commonExpectation
	authType string
}

func (e *ExpectedAuthenticate) String() string {
	msg := "ExpectedAuthenticate => expecting Authenticate which:"
	if e.authType == "" {
		msg += "\n\t- has any authenticator"
	} else {
		msg += fmt.Sprintf("\n\t- has authenticator of type %s", e.authType)
	}
	if e.err != nil {
		msg += fmt.Sprintf("\n\t- should return error: %s", e.err)
	}
	return msg
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

// ExpectedClusterSetup is used to manage *kivik.Client.ClusterSetup
// expectation returned by Mock.ExpectClusterSetup.
type ExpectedClusterSetup struct {
	commonExpectation
	action interface{}
}

func (e *ExpectedClusterSetup) String() string {
	msg := "ExpectedClusterSetup => expecting ClusterSetup which:"
	if e.action == nil {
		msg += "\n\t- has any action"
	} else {
		msg += "\n\t- has the action:\n\t\t"
		b, err := json.MarshalIndent(e.action, "\t\t", "  ")
		if err != nil {
			msg += fmt.Sprintf("<<unmarshalable object: %s>>", err)
		} else {
			msg += fmt.Sprintf("%s", string(b))
		}
	}
	if e.err != nil {
		msg += fmt.Sprintf("\n\t- should return error: %s", e.err)
	}
	return msg
}
