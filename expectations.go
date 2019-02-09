package kivikmock

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/go-kivik/kivik"
)

type expectation interface {
	fulfill()
	fulfilled() bool
	Lock()
	Unlock()
	String() string
	// method should return the name of the method that would trigger this
	// condition. If verbose is true, the output should disambiguate between
	// different calls to the same method.
	method(verbose bool) string
}

// commonExpectation satisfies the expectation interface, except the String()
// and method() methods.
type commonExpectation struct {
	sync.Mutex
	triggered bool
	err       error // nolint: structcheck
}

func (e *commonExpectation) fulfill() {
	e.triggered = true
}

func (e *commonExpectation) fulfilled() bool {
	return e.triggered
}

// ExpectedClose is used to manage *kivik.Client.Close expectation returned
// by Mock.ExpectClose.
type ExpectedClose struct {
	commonExpectation
}

func (e *ExpectedClose) method(v bool) string {
	if v {
		return "Close(ctx)"
	}
	return "Close()"
}

func (e *ExpectedClose) equal(_ expectation) bool { return true }

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

func (e *ExpectedAllDBs) method(v bool) string {
	if v {
		if e.options == nil {
			return "AllDBs(ctx, nil)"
		}
		return fmt.Sprintf("AllDBs(ctx, %#v)", e.options)
	}
	return "AllDBs()"
}

func (e *ExpectedAllDBs) equal(other expectation) bool {
	if e.options == nil {
		return true
	}
	return reflect.DeepEqual(e.options, other.(*ExpectedAllDBs).options)
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

func (e *ExpectedAuthenticate) method(v bool) string {
	if v {
		return fmt.Sprintf("Authenticate(ctx, <%s>)", e.authType)
	}
	return "Authenticate()"
}

func (e *ExpectedAuthenticate) equal(other expectation) bool {
	if e.authType == "" {
		return true
	}
	o, _ := other.(*ExpectedAuthenticate)
	return e.authType == o.authType
}

func (e *ExpectedAuthenticate) String() string {
	msg := "ExpectedAuthenticate => expecting Authenticate"
	modifiers := []string{}
	if e.authType != "" {
		modifiers = append(modifiers, fmt.Sprintf("expects an authenticator of type '%s'", e.authType))
	}
	if e.err != nil {
		modifiers = append(modifiers, "should return an error")
	}
	if len(modifiers) > 0 {
		return msg + " which " + strings.Join(modifiers, " and ")
	}
	return msg
}

func (e *ExpectedAuthenticate) Format(f fmt.State, verb rune) {
	switch verb {
	case 's':
		fmt.Fprint(f, e.String())
	case 'v':
		if !f.Flag('+') {
			fmt.Fprintf(f, "%s", e)
			return
		}
		msg := "ExpectedAuthenticate => expecting Authenticate which:"
		if e.authType == "" {
			msg += "\n\t- expects any authenticator"
		} else {
			msg += "\n\t- expects an authenticator of type: " + e.authType
		}
		if e.err != nil {
			msg += fmt.Sprintf("\n\t- should return error: %s", e.err)
		}
		fmt.Fprint(f, msg)
	}
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

func (e *ExpectedClusterSetup) method(v bool) string {
	if v {
		return fmt.Sprintf("ClusterSetup(ctx, %#v)", e.action)
	}
	return "ClusterSetup()"
}

// Format satisfies the fmt.Formatter interface
func (e *ExpectedClusterSetup) Format(f fmt.State, verb rune) {
	switch verb {
	case 's':
		fmt.Fprint(f, e.String())
	case 'v':
		if !f.Flag('+') {
			fmt.Fprintf(f, "%s", e)
			return
		}
		msg := "ExpectedClusterSetup => expecting ClusterSetup which:"
		if e.action == nil {
			msg += "\n\t- expects any action"
		} else {
			msg += "\n\t- expects the following action:"
			b, err := json.MarshalIndent(e.action, "\t\t", "  ")
			if err != nil {
				msg += fmt.Sprintf("\n\t\t<<unmarshalable: %s>>", err)
			} else {
				msg += "\n\t\t" + string(b)
			}
		}
		if e.err != nil {
			msg += fmt.Sprintf("\n\t- should return error: %s", e.err)
		}
		fmt.Fprint(f, msg)
	}
}

func (e *ExpectedClusterSetup) String() string {
	msg := "ExpectedClusterSetup => expecting ClusterSetup"
	modifiers := []string{}
	if e.action != nil {
		modifiers = append(modifiers, "has the desired action")
	}
	if e.err != nil {
		modifiers = append(modifiers, "should return an error")
	}
	if len(modifiers) > 0 {
		msg = msg + " which " + strings.Join(modifiers, " and ")
	}
	return msg
}
